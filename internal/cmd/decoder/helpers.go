package decoder

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"strings"

	"github.com/fallais/govid/internal/cmd/decoder/models"

	"github.com/fxamacker/cbor"
	"github.com/minvws/base45-go/eubase45"
)

func decode(cose string) (*models.UnverifiedCOSE, error) {
	// Remove the prefix
	coseWithoutPrefix, err := removePrefix(cose)
	if err != nil {
		return nil, err
	}

	// Base45 decoding
	coseCompressed, err := base45decode(coseWithoutPrefix)
	if err != nil {
		return nil, err
	}

	// Zlib decompressing
	coseDecompressed, err := decompress(coseCompressed)
	if err != nil {
		return nil, err
	}

	return decodeCOSE(coseDecompressed)
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

func decodeCOSE(coseData []byte) (*models.UnverifiedCOSE, error) {
	var v models.SignedCWT
	if err := cbor.Unmarshal(coseData, &v); err != nil {
		return nil, fmt.Errorf("cbor.Unmarshal: %v", err)
	}

	var p models.COSEHeader
	if len(v.Protected) > 0 {
		if err := cbor.Unmarshal(v.Protected, &p); err != nil {
			return nil, fmt.Errorf("cbor.Unmarshal(v.Protected): %v", err)
		}
	}

	var c models.Claims
	if err := cbor.Unmarshal(v.Payload, &c); err != nil {
		return nil, fmt.Errorf("cbor.Unmarshal(v.Payload): %v", err)
	}

	return &models.UnverifiedCOSE{
		V:      v,
		P:      p,
		Claims: c,
	}, nil
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
