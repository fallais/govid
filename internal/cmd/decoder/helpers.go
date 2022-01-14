package decoder

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"strings"
	//"encoding/base64"

	"github.com/fallais/govid/internal/cmd/decoder/models"

	"github.com/fxamacker/cbor"
	"github.com/minvws/base45-go/eubase45"
)

const Prefix = "HC1:"

func NewDCCFromQRCode(data string) (*models.DigitalCovidCertificate, error) {
	if !strings.HasPrefix(data, Prefix) {
		return nil, fmt.Errorf("data does not start with `HC1:`")
	}

	// Remove the prefix
	dataWithoutPrefix := strings.TrimPrefix(data, Prefix)

	// Base45 decoding
	decoded, err := eubase45.EUBase45Decode([]byte(dataWithoutPrefix))
	if err != nil {
		return nil, err
	}

	// Zlib decompressing
	decompressed, err := decompress(decoded)
	if err != nil {
		return nil, err
	}

	// Unmarshal into a SignedCWT
	var v SignedCWT
	err = cbor.Unmarshal(decompressed, &v)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling the SignedCWT: %v", err)
	}

	// Unmarshal the claims

	// Unmarshal the claims
	var c Claims
	err = cbor.Unmarshal(v.Payload, &c)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling the claims: %v", err)
	}

	return &c.HCert.DCC, nil
}

func decompress(compressed []byte) ([]byte, error) {
	zr, err := zlib.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, zr)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
