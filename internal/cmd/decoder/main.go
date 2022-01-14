package decoder

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"

	"github.com/liyue201/goqr"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Run is a convenient function for Cobra.
func Run(cmd *cobra.Command, args []string) {
	// Flags
	dcc, err := cmd.Flags().GetString("dcc")
	if err != nil {
		logrus.WithError(err).Fatalln("Error while getting flag")
	}

	// Read the DCC
	// TODO: local or URL
	data, err := ioutil.ReadFile(dcc)
	if err != nil {
		logrus.WithError(err).Fatalln("Error while reading configuration file")
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		fmt.Printf("image.Decode error: %v\n", err)
		return
	}

	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		fmt.Printf("Recognize failed: %v\n", err)
		return
	}

	for _, qrCode := range qrCodes {
		// Decode the payload
		uc, err := decode(string(qrCode.Payload))
		if err != nil {
			fmt.Errorf("error while decoding: %v", err)
			return
		}

		fmt.Println("Name of the person: ", uc.Claims.HCert.DCC.Name)
		fmt.Println("Date of birth of the person: ", uc.Claims.HCert.DCC.DateOfBirth)
		fmt.Println("Vaccines:")
		for _, vaccine := range uc.Claims.HCert.DCC.Vaccines {
			fmt.Println(vaccine)
		}
	}
}
