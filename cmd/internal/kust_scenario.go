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
