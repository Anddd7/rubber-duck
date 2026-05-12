package internal

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func newKustPatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kust-patch",
		Short: "Manage patch components",
		Long:  "Manage patches under kustomize/_patches/<patch> and apply/rollback them to workloads.",
		Example: "  rubber-duck kust-patch list\n" +
			"  rubber-duck kust-patch get curl\n" +
			"  rubber-duck kust-patch install curl --pod <pod-name> -n default\n" +
			"  rubber-duck kust-patch uninstall curl --deploy httpbin -n default",
	}

	cmd.AddCommand(newKustPatchListCmd())
	cmd.AddCommand(newKustPatchGetCmd())
	cmd.AddCommand(newKustPatchInstallCmd())
	cmd.AddCommand(newKustPatchUninstallCmd())

	return cmd
}

func newKustPatchListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List patch components",
		Long:  "List reusable patch components.",
		Example: "  rubber-duck kust-patch list",
		RunE: func(cmd *cobra.Command, args []string) error {
			components, err := discoverComponentsByType(componentTypePatch)
			if err != nil {
				return err
			}

			w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "NAME\tTYPE\tDESCRIPTION")
			for _, c := range components {
				desc := c.Desc
				if desc == "" {
					desc = "-"
				}
				fmt.Fprintf(w, "%s\t%s\t%s\n", c.Name, c.Type, desc)
			}
			return w.Flush()
		},
	}
}

func newKustPatchGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <patch>",
		Short: "Get patch details",
		Long:  "Show patch metadata, path, and description.",
		Example: "  rubber-duck kust-patch get curl",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			component, err := findComponentByType(args[0], componentTypePatch)
			if err != nil {
				return err
			}

			printComponentDetail(cmd, component)
			return nil
		},
	}
}

func newKustPatchInstallCmd() *cobra.Command {
	var pod string
	var namespace string

	cmd := &cobra.Command{
		Use:   "install <patch>",
		Short: "Apply patch to workload owner of a pod",
		Long:  "Resolve deployment owner from the target pod, then patch the deployment template.",
		Example: "  rubber-duck kust-patch install curl --pod httpbin-xxxx -n default",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if pod == "" {
				return errors.New("--pod is required")
			}

			patchComponent, err := findComponentByType(args[0], componentTypePatch)
			if err != nil {
				return err
			}

			deployment, err := resolveDeploymentFromPod(cmd, pod, namespace)
			if err != nil {
				return err
			}

			patchContent, err := os.ReadFile(filepath.Join(patchComponent.Path, "kustomization.yaml"))
			if err != nil {
				return err
			}

			jsonPatch := extractPatchBlock(string(patchContent))
			if jsonPatch == "" {
				return fmt.Errorf("no patch found in %s", patchComponent.Path)
			}

			kArgs := []string{"patch", "deployment", deployment, "--type=json", "-p", jsonPatch}
			if namespace != "" {
				kArgs = append(kArgs, "-n", namespace)
			}

			return runCommand(cmd, "kubectl", kArgs...)
		},
	}

	cmd.Flags().StringVar(&pod, "pod", "", "pod name to resolve workload")
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "", "target namespace")

	return cmd
}

func newKustPatchUninstallCmd() *cobra.Command {
	var pod string
	var deploy string
	var namespace string

	cmd := &cobra.Command{
		Use:   "uninstall <patch>",
		Short: "Rollback workload to previous revision",
		Long:  "Rollback deployment to previous revision using kubectl rollout undo.",
		Example: "  rubber-duck kust-patch uninstall curl --deploy httpbin -n default\n" +
			"  rubber-duck kust-patch uninstall curl --pod httpbin-xxxx -n default",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := findComponentByType(args[0], componentTypePatch); err != nil {
				return err
			}

			targetDeploy := deploy
			if targetDeploy == "" {
				if pod == "" {
					return errors.New("one of --pod or --deploy is required")
				}
				resolved, err := resolveDeploymentFromPod(cmd, pod, namespace)
				if err != nil {
					return err
				}
				targetDeploy = resolved
			}

			kArgs := []string{"rollout", "undo", "deployment", targetDeploy}
			if namespace != "" {
				kArgs = append(kArgs, "-n", namespace)
			}

			return runCommand(cmd, "kubectl", kArgs...)
		},
	}

	cmd.Flags().StringVar(&pod, "pod", "", "pod name to resolve workload")
	cmd.Flags().StringVar(&deploy, "deploy", "", "deployment name")
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "", "target namespace")

	return cmd
}
