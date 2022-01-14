package decoder

import (
	"bytes"
	"compress/zlib"
	"crypto"
	"crypto/sha256"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"

	"github.com/fxamacker/cbor"
	"github.com/liyue201/goqr"
	"github.com/minvws/base45-go/eubase45"
	"github.com/veraison/go-cose"
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
		fmt.Printf("qrCode text: %s\n", qrCode.Payload)
	}
}

func base45decode(s string) ([]byte, error) {
	return eubase45.EUBase45Decode([]byte(s))
}
