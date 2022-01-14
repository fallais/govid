package decoder

import (
	"crypto/x509"
	"fmt"

	"github.com/fallais/govid/internal/cmd/decoder/models"

	"github.com/fxamacker/cbor"
)

type coseHeader struct {
	Alg int    `cbor:"1,keyasint,omitempty"`
	Kid []byte `cbor:"4,keyasint,omitempty"`
	IV  []byte `cbor:"5,keyasint,omitempty"`
}

type signedCWT struct {
	_           struct{} `cbor:",toarray"`
	Protected   []byte
	Unprotected map[interface{}]interface{}
	Payload     []byte
	Signature   []byte
}

type hcert struct {
	DCC models.DigitalCovidCertificate `cbor:"1,keyasint"`
}

type claims struct {
	Iss       string `cbor:"1,keyasint"`
	Sub       string `cbor:"2,keyasint"`
	Aud       string `cbor:"3,keyasint"`
	Exp       int64  `cbor:"4,keyasint"`
	Nbf       int    `cbor:"5,keyasint"`
	Iat       int64  `cbor:"6,keyasint"`
	Cti       []byte `cbor:"7,keyasint"`
	HCert     hcert  `cbor:"-260,keyasint"`
	LightCert hcert  `cbor:"-250,keyasint"`
}

type unverifiedCOSE struct {
	v      signedCWT
	p      coseHeader
	claims claims
	cert   *x509.Certificate // set after verification
}

func decodeCOSE(coseData []byte) (*unverifiedCOSE, error) {
	var v signedCWT
	if err := cbor.Unmarshal(coseData, &v); err != nil {
		return nil, fmt.Errorf("cbor.Unmarshal: %v", err)
	}

	var p coseHeader
	if len(v.Protected) > 0 {
		if err := cbor.Unmarshal(v.Protected, &p); err != nil {
			return nil, fmt.Errorf("cbor.Unmarshal(v.Protected): %v", err)
		}
	}

	var c claims
	if err := cbor.Unmarshal(v.Payload, &c); err != nil {
		return nil, fmt.Errorf("cbor.Unmarshal(v.Payload): %v", err)
	}

	return &unverifiedCOSE{
		v:      v,
		p:      p,
		claims: c,
	}, nil
}
