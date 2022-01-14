package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "govid",
	Short: "Golang tool for Digital Covid Certificate (DCC)",
}

func init() {
	rootCmd.AddCommand(decoderCmd)
	//rootCmd.AddCommand(encoderCmd)
}

// Execute the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
