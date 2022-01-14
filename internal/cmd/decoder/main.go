package decoder

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"

	"github.com/fallais/govid/internal/cmd/decoder/models"

	"github.com/liyue201/goqr"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Run is a convenient function for Cobra.
func Run(cmd *cobra.Command, args []string) {
	// Flags
	qrcode, err := cmd.Flags().GetString("qrcode")
	if err != nil {
		logrus.WithError(err).Fatalln("Error while getting flag")
	}

	// Read the DCC
	// TODO: local or URL
	data, err := ioutil.ReadFile(qrcode)
	if err != nil {
		logrus.WithError(err).Fatalln("Error while reading configuration file")
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		logrus.WithError(err).Fatalln("error while decoding image")
		return
	}

	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		logrus.WithError(err).Fatalln("error while recognizing qrcodes into image")
		return
	}

	for _, qrCode := range qrCodes {
		// Get DCC from the QRCode
		dcc, err := NewDCCFromQRCode(string(qrCode.Payload))
		if err != nil {
			logrus.WithError(err).Fatalln("error while get DCC from the QRCode")
			return
		}

		fmt.Println("Name of the person:", dcc.Name.GivenName, dcc.Name.FamilyName)
		fmt.Println("Date of birth of the person:", dcc.DateOfBirth)
		fmt.Println("Vaccines:")
		for _, vaccine := range dcc.Vaccines {
			fmt.Println("Vaccine name: ", models.VaccineMarketingAuthorizationHolder[vaccine.Manufacturer])
			fmt.Println("Dose number: ", vaccine.Doses)
		}
	}
}
