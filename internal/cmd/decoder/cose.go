package decoder

import (
	"github.com/fallais/govid/internal/cmd/decoder/models"
)

type COSEHeader struct {
	Alg int    `cbor:"1,keyasint,omitempty"`
	Kid []byte `cbor:"4,keyasint,omitempty"`
}

type SignedCWT struct {
	_           struct{} `cbor:",toarray"`
	Protected   []byte
	Unprotected map[interface{}]interface{}
	Payload     []byte
	Signature   []byte
}

type HCert struct {
	DCC models.DigitalCovidCertificate `cbor:"1,keyasint"`
}

type Claims struct {
	Iss       string `cbor:"1,keyasint"`
	Sub       string `cbor:"2,keyasint"`
	Aud       string `cbor:"3,keyasint"`
	Exp       int64  `cbor:"4,keyasint"`
	Nbf       int    `cbor:"5,keyasint"`
	Iat       int64  `cbor:"6,keyasint"`
	Cti       []byte `cbor:"7,keyasint"`
	HCert     HCert  `cbor:"-260,keyasint"`
	LightCert HCert  `cbor:"-250,keyasint"`
}
