package internal

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

type componentType string

const (
	componentTypeNormal   componentType = "component"
	componentTypePatch    componentType = "patch"
	componentTypeScenario componentType = "scenario"
)

type componentInfo struct {
	Name     string
	Type     componentType
	Path     string
	Readme   string
	Desc     string
	HasBase  bool
	Overlays []string
}

func newKustCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kust",
		Short: "Manage standard kustomize components",
	}

	cmd.AddCommand(newKustListCmd())
	cmd.AddCommand(newKustGetCmd())
	cmd.AddCommand(newKustInstallCmd())
	cmd.AddCommand(newKustUninstallCmd())

	return cmd
}

func newKustListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List standard components",
		RunE: func(cmd *cobra.Command, args []string) error {
			components, err := discoverComponentsByType(componentTypeNormal)
			if err != nil {
				return err
			}

			w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "NAME\tTYPE\tBASE\tOVERLAYS\tDESCRIPTION")
			for _, c := range components {
				base := "no"
				if c.HasBase {
					base = "yes"
				}
				overlays := "-"
				if len(c.Overlays) > 0 {
					overlays = strings.Join(c.Overlays, ",")
				}
				desc := c.Desc
				if desc == "" {
					desc = "-"
				}
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", c.Name, c.Type, base, overlays, desc)
			}
			return w.Flush()
		},
	}

	return cmd
}

func newKustGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <component>",
		Short: "Get standard component details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			component, err := findComponentByType(args[0], componentTypeNormal)
			if err != nil {
				return err
			}

			printComponentDetail(cmd, component)
			return nil
		},
	}

	return cmd
}

func newKustInstallCmd() *cobra.Command {
	var overlay string
	var namespace string
	var dryRun bool

	cmd := &cobra.Command{
		Use:   "install <component>",
		Short: "Install a standard component",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if overlay == "" {
				return errors.New("--overlay is required")
			}

			component, err := findComponentByType(args[0], componentTypeNormal)
			if err != nil {
				return err
			}

			targetPath, err := resolveTargetKustomization(component, overlay)
			if err != nil {
				return err
			}

			kArgs := []string{"apply", "-k", targetPath}
			if dryRun {
				kArgs = append(kArgs, "--dry-run=client")
			}
			if namespace != "" {
				kArgs = append(kArgs, "-n", namespace)
			}

			return runCommand(cmd, "kubectl", kArgs...)
		},
	}

	cmd.Flags().StringVar(&overlay, "overlay", "", "overlay name or 'base'")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "print actions without applying")
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "", "target namespace")

	return cmd
}

func newKustUninstallCmd() *cobra.Command {
	var overlay string
	var namespace string
	var dryRun bool

	cmd := &cobra.Command{
		Use:   "uninstall <component>",
		Short: "Uninstall a standard component",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if overlay == "" {
				return errors.New("--overlay is required")
			}

			component, err := findComponentByType(args[0], componentTypeNormal)
			if err != nil {
				return err
			}

			targetPath, err := resolveTargetKustomization(component, overlay)
			if err != nil {
				return err
			}

			kArgs := []string{"delete", "-k", targetPath}
			if dryRun {
				kArgs = append(kArgs, "--dry-run=client")
			}
			if namespace != "" {
				kArgs = append(kArgs, "-n", namespace)
			}

			return runCommand(cmd, "kubectl", kArgs...)
		},
	}

	cmd.Flags().StringVar(&overlay, "overlay", "", "overlay name or 'base'")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "print actions without deleting")
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "", "target namespace")

	return cmd
}

func discoverComponentsByType(t componentType) ([]componentInfo, error) {
	root, err := repoRoot()
	if err != nil {
		return nil, err
	}
	baseDir := filepath.Join(root, "kustomize")

	switch t {
	case componentTypeNormal:
		return discoverNormalComponents(baseDir)
	case componentTypePatch:
		return discoverPatchComponents(baseDir)
	case componentTypeScenario:
		return discoverScenarioComponents(baseDir)
	default:
		return nil, fmt.Errorf("unsupported component type %q", t)
	}
}

func discoverNormalComponents(baseDir string) ([]componentInfo, error) {
	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return nil, err
	}

	out := make([]componentInfo, 0)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasPrefix(name, "_") {
			continue
		}

		path := filepath.Join(baseDir, name)
		out = append(out, componentInfo{
			Name:     name,
			Type:     componentTypeNormal,
			Path:     path,
			Readme:   findReadme(path),
			Desc:     readFirstNonHeaderLine(findReadme(path)),
			HasBase:  hasFile(filepath.Join(path, "kustomization", "base", "kustomization.yaml")),
			Overlays: listOverlayNames(filepath.Join(path, "kustomization", "overlays")),
		})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}

func discoverPatchComponents(baseDir string) ([]componentInfo, error) {
	parent := filepath.Join(baseDir, "_patches")
	entries, err := os.ReadDir(parent)
	if err != nil {
		return nil, err
	}

	out := make([]componentInfo, 0)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		path := filepath.Join(parent, name)
		readme := findReadme(path)
		out = append(out, componentInfo{
			Name:    name,
			Type:    componentTypePatch,
			Path:    path,
			Readme:  readme,
			Desc:    readFirstNonHeaderLine(readme),
			HasBase: hasFile(filepath.Join(path, "kustomization.yaml")),
		})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}

func discoverScenarioComponents(baseDir string) ([]componentInfo, error) {
	parent := filepath.Join(baseDir, "_scenarios")
	entries, err := os.ReadDir(parent)
	if err != nil {
		return nil, err
	}

	out := make([]componentInfo, 0)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		path := filepath.Join(parent, name)
		readme := findReadme(path)
		out = append(out, componentInfo{
			Name:    name,
			Type:    componentTypeScenario,
			Path:    path,
			Readme:  readme,
			Desc:    readFirstNonHeaderLine(readme),
			HasBase: hasFile(filepath.Join(path, "kustomization.yaml")),
		})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}

func findComponentByType(name string, t componentType) (componentInfo, error) {
	components, err := discoverComponentsByType(t)
	if err != nil {
		return componentInfo{}, err
	}

	for _, c := range components {
		if c.Name == name {
			return c, nil
		}
	}

	return componentInfo{}, fmt.Errorf("%s %q not found", t, name)
}

func printComponentDetail(cmd *cobra.Command, component componentInfo) {
	out := cmd.OutOrStdout()
	fmt.Fprintf(out, "name: %s\n", component.Name)
	fmt.Fprintf(out, "type: %s\n", component.Type)
	fmt.Fprintf(out, "path: %s\n", component.Path)
	if component.Desc != "" {
		fmt.Fprintf(out, "description: %s\n", component.Desc)
	}

	if component.Type == componentTypeNormal {
		fmt.Fprintf(out, "base: %t\n", component.HasBase)
		if len(component.Overlays) == 0 {
			fmt.Fprintln(out, "overlays: -")
		} else {
			fmt.Fprintf(out, "overlays: %s\n", strings.Join(component.Overlays, ","))
		}
	}

	if component.Desc == "" {
		if readmeSummary := readFirstNonHeaderLine(component.Readme); readmeSummary != "" {
			fmt.Fprintf(out, "readme: %s\n", readmeSummary)
		}
	}
}

func resolveTargetKustomization(component componentInfo, overlay string) (string, error) {
	if overlay == "base" {
		path := filepath.Join(component.Path, "kustomization", "base")
		if !hasFile(filepath.Join(path, "kustomization.yaml")) {
			return "", fmt.Errorf("base not found for component %q", component.Name)
		}
		return path, nil
	}

	path := filepath.Join(component.Path, "kustomization", "overlays", overlay)
	if !hasFile(filepath.Join(path, "kustomization.yaml")) {
		return "", fmt.Errorf("overlay %q not found for component %q", overlay, component.Name)
	}
	return path, nil
}

func resolveDeploymentFromPod(cmd *cobra.Command, pod, namespace string) (string, error) {
	args := []string{"get", "pod", pod, "-o", "jsonpath={.metadata.ownerReferences[0].name}"}
	if namespace != "" {
		args = append(args, "-n", namespace)
	}

	rsName, err := runCommandOutput("kubectl", args...)
	if err != nil {
		return "", err
	}
	rsName = strings.TrimSpace(rsName)
	if rsName == "" {
		return "", fmt.Errorf("pod %q has no owner", pod)
	}

	args = []string{"get", "rs", rsName, "-o", "jsonpath={.metadata.ownerReferences[0].name}"}
	if namespace != "" {
		args = append(args, "-n", namespace)
	}

	deployName, err := runCommandOutput("kubectl", args...)
	if err != nil {
		return "", err
	}
	deployName = strings.TrimSpace(deployName)
	if deployName == "" {
		return "", fmt.Errorf("replicaset %q has no deployment owner", rsName)
	}

	_ = cmd
	return deployName, nil
}

func extractPatchBlock(content string) string {
	idx := strings.Index(content, "patch: |-")
	if idx < 0 {
		return ""
	}

	lines := strings.Split(content[idx:], "\n")
	if len(lines) < 2 {
		return ""
	}

	var out []string
	for _, line := range lines[1:] {
		if strings.HasPrefix(line, "    ") {
			out = append(out, strings.TrimPrefix(line, "    "))
			continue
		}
		if strings.TrimSpace(line) == "" {
			out = append(out, "")
			continue
		}
		break
	}

	return strings.Join(out, "\n")
}

func runCommand(cmd *cobra.Command, name string, args ...string) error {
	execCmd := exec.Command(name, args...)
	execCmd.Stdout = cmd.OutOrStdout()
	execCmd.Stderr = cmd.ErrOrStderr()
	return execCmd.Run()
}

func runCommandOutput(name string, args ...string) (string, error) {
	execCmd := exec.Command(name, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	execCmd.Stdout = &stdout
	execCmd.Stderr = &stderr
	if err := execCmd.Run(); err != nil {
		if stderr.Len() > 0 {
			return "", errors.New(strings.TrimSpace(stderr.String()))
		}
		return "", err
	}
	return stdout.String(), nil
}

func repoRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return wd, nil
}

func listOverlayNames(overlaysDir string) []string {
	entries, err := os.ReadDir(overlaysDir)
	if err != nil {
		return nil
	}

	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if hasFile(filepath.Join(overlaysDir, entry.Name(), "kustomization.yaml")) {
			names = append(names, entry.Name())
		}
	}
	sort.Strings(names)
	return names
}

func hasFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func findReadme(dir string) string {
	for _, name := range []string{"README.md", "readme.md"} {
		path := filepath.Join(dir, name)
		if hasFile(path) {
			return path
		}
	}
	return ""
}

func readFirstNonHeaderLine(readmePath string) string {
	if readmePath == "" {
		return ""
	}

	f, err := os.Open(readmePath)
	if err != nil {
		return ""
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		return line
	}

	return ""
}
