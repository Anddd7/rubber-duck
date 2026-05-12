package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "rubber-duck",
		Short: "A rubber duck in your cluster",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("rubber-duck: ready")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
