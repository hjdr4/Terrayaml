package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Long: "Convert YAML files into JSON or HCL",
}

func init() {
	rootCmd.AddCommand(convert)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
