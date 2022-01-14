package cmd

import (
	"github.com/fallais/govid/internal/cmd/decoder"

	"github.com/spf13/cobra"
)

var decoderCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode a QRCode to get the Digital Covid Certificate (DCC)",
	Run:   decoder.Run,
}

func init() {
	decoderCmd.Flags().StringP("qrcode", "q", "assets/qrcodes/qrcode2.png", "QRCode")
}
