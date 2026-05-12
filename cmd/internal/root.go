package internal

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "rubber-duck",
		Short: "A rubber duck in your cluster",
		Long:  "rubber-duck manages reusable DevOps assets (kustomize components, patches, and scenarios).",
		Example: "  rubber-duck kust list\n" +
			"  rubber-duck kust get httpbin\n" +
			"  rubber-duck kust install httpbin --overlay base -n default\n" +
			"  rubber-duck kust-patch install curl --pod <pod-name> -n default\n" +
			"  rubber-duck kust-scenario list",
	}

	rootCmd.AddCommand(newKustCmd())
	rootCmd.AddCommand(newKustPatchCmd())
	rootCmd.AddCommand(newKustScenarioCmd())

	return rootCmd
}
