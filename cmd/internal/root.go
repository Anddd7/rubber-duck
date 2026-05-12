package internal

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "rubber-duck",
		Short: "A rubber duck in your cluster",
	}

	rootCmd.AddCommand(newKustCmd())
	rootCmd.AddCommand(newKustPatchCmd())
	rootCmd.AddCommand(newKustScenarioCmd())

	return rootCmd
}
