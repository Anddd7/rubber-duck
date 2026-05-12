package main

import (
	"fmt"
	"os"

	"github.com/Anddd7/rubber-duck/cmd/internal"
)

func main() {
	if err := internal.NewRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
