package decoder

import (
	"bytes"
	//"compress/zlib"
	//"crypto"
	//"crypto/sha256"
	//"crypto/x509"
	//"errors"
	"fmt"
	//"io"
	"strings"
	//"time"

	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"

	"github.com/liyue201/goqr"
	"github.com/minvws/base45-go/eubase45"
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
		uc, err := Decode(string(qrCode.Payload))
		if err != nil {
			fmt.Errorf("error while decoding: %v", err)
			return
		}

		fmt.Println("Name of the person: ", uc.claims.HCert.DCC.Name)
		fmt.Println("Date of birth of the person: ", uc.claims.HCert.DCC.DateOfBirth)
		fmt.Println("Vaccines:")
		for _, vaccine := range uc.claims.HCert.DCC.Vaccines {
			fmt.Println(vaccine)
		}
	}
}

func base45decode(s string) ([]byte, error) {
	return eubase45.EUBase45Decode([]byte(s))
}

func removePrefix(s string) (string, error) {
	if !strings.HasPrefix(s, "HC1:") && !strings.HasPrefix(s, "LT1:") {
		return "", fmt.Errorf("should start with HC1 or LT1")
	}

	return strings.TrimPrefix(strings.TrimPrefix(s, "HC1:"), "LT1:"), nil
}

// Decode returns the decoded EUDCC.
func Decode(s string) (*unverifiedCOSE, error) {
	// Remove the prefix (`HC1:` or `LT1:`)
	unprefixed, err := removePrefix(s)
	if err != nil {
		return nil, err
	}

	// Base45 decode
	compressed, err := base45decode(unprefixed)
	if err != nil {
		return nil, err
	}

	coseData, err := decompress(compressed)
	if err != nil {
		return nil, err
	}

	return decodeCOSE(coseData)
}
