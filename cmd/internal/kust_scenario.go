package internal

import (
	"fmt"
	"path/filepath"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func newKustScenarioCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kust-scenario",
		Short: "Manage scenario components",
		Long:  "Manage scenarios under kustomize/_scenarios/<scenario>.",
		Example: "  rubber-duck kust-scenario list\n" +
			"  rubber-duck kust-scenario get oauth-github\n" +
			"  rubber-duck kust-scenario install oauth-github -n default\n" +
			"  rubber-duck kust-scenario uninstall oauth-github -n default",
	}

	cmd.AddCommand(newKustScenarioListCmd())
	cmd.AddCommand(newKustScenarioGetCmd())
	cmd.AddCommand(newKustScenarioInstallCmd())
	cmd.AddCommand(newKustScenarioUninstallCmd())

	return cmd
}

func newKustScenarioListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List scenario components",
		Long:  "List reusable scenario compositions.",
		Example: "  rubber-duck kust-scenario list",
		RunE: func(cmd *cobra.Command, args []string) error {
			components, err := discoverComponentsByType(componentTypeScenario)
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

func newKustScenarioGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <scenario>",
		Short: "Get scenario details",
		Long:  "Show scenario metadata, path, and description.",
		Example: "  rubber-duck kust-scenario get oauth-github",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			component, err := findComponentByType(args[0], componentTypeScenario)
			if err != nil {
				return err
			}

			printComponentDetail(cmd, component)
			return nil
		},
	}
}

func newKustScenarioInstallCmd() *cobra.Command {
	var namespace string
	var dryRun bool

	cmd := &cobra.Command{
		Use:   "install <scenario>",
		Short: "Install a scenario",
		Long:  "Install scenario manifests with kubectl apply -k.",
		Example: "  rubber-duck kust-scenario install oauth-github -n default\n" +
			"  rubber-duck kust-scenario install grpc-timeout --dry-run",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			component, err := findComponentByType(args[0], componentTypeScenario)
			if err != nil {
				return err
			}

			targetPath := filepath.Join(component.Path)
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

	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "print actions without applying")
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "", "target namespace")

	return cmd
}

func newKustScenarioUninstallCmd() *cobra.Command {
	var namespace string
	var dryRun bool

	cmd := &cobra.Command{
		Use:   "uninstall <scenario>",
		Short: "Uninstall a scenario",
		Long:  "Uninstall scenario manifests with kubectl delete -k.",
		Example: "  rubber-duck kust-scenario uninstall oauth-github -n default\n" +
			"  rubber-duck kust-scenario uninstall grpc-timeout --dry-run",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			component, err := findComponentByType(args[0], componentTypeScenario)
			if err != nil {
				return err
			}

			targetPath := filepath.Join(component.Path)
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

	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "print actions without deleting")
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "", "target namespace")

	return cmd
}
