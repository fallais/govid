package cmd

import (
	"fmt"
	"os"

	"github.com/fallais/govid/internal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "govid",
	Short:             "Golang tool for Digital Covid Certificate (DCC)",
	Run:               internal.Run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
