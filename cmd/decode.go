package cmd

import (
	"github.com/govid/internal/cmd/decoder"

	"github.com/spf13/cobra"
)

var decoderCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode a Digital Covid Certificate (DCC)",
	Run:   decoder.Run,
}

func init() {
	decoderCmd.Flags().StringP("dcc", "d", "dcc.png", "Digital Covid Certificate")
}
