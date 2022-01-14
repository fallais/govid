package models

// DigitalCovidCertificate (DCC) is the certificate for Covid.
// see https://github.com/ehn-dcc-development/ehn-dcc-schema/blob/release/1.3.0/DCC.schema.json
type DigitalCovidCertificate struct {
	Version     string          `cbor:"ver" json:"ver"`
	Name        Name            `cbor:"nam" json:"nam"`
	DateOfBirth string          `cbor:"dob" json:"dob"`
	Vaccines    []VaccineEntry  `cbor:"v" json:"v"`
	Tests       []TestEntry     `cbor:"t" json:"t"`
	Recoveries  []RecoveryEntry `cbor:"r" json:"r"`
}

// Name holds the name of the person.
type Name struct {
	FamilyName    string `cbor:"fn" json:"fn"`
	FamilyNameStd string `cbor:"fnt" json:"fnt"`
	GivenName     string `cbor:"gn" json:"gn"`
	GivenNameStd  string `cbor:"gnt" json:"gnt"`
}

type VaccineEntry struct {
	Target        string  `cbor:"tg" json:"tg"`
	Vaccine       string  `cbor:"vp" json:"vp"`
	Product       string  `cbor:"mp" json:"mp"`
	Manufacturer  string  `cbor:"ma" json:"ma"`
	Doses         float64 `cbor:"dn" json:"dn"` // int per the spec, but float64 e.g. in IE
	DoseSeries    float64 `cbor:"sd" json:"sd"` // int per the spec, but float64 e.g. in IE
	Date          string  `cbor:"dt" json:"dt"`
	Country       string  `cbor:"co" json:"co"`
	Issuer        string  `cbor:"is" json:"is"`
	CertificateID string  `cbor:"ci" json:"ci"`
}

type TestEntry struct {
	Target         string `cbor:"tg" json:"tg"`
	TestType       string `cbor:"tt" json:"tt"`
	Name           string `cbor:"nm" json:"nm"`
	Manufacturer   string `cbor:"ma" json:"ma"`
	SampleDatetime string `cbor:"sc" json:"sc"`
	TestResult     string `cbor:"tr" json:"tr"`
	TestingCentre  string `cbor:"tc" json:"tc"`
	Country        string `cbor:"co" json:"co"`
	Issuer         string `cbor:"is" json:"is"`
	CertificateID  string `cbor:"ci" json:"ci"`
}

type RecoveryEntry struct {
	Target            string `cbor:"tg" json:"tg"`
	Country           string `cbor:"co" json:"co"`
	Issuer            string `cbor:"is" json:"is"`
	FirstPositiveTest string `cbor:"fr" json:"fr"`
	ValidFrom         string `cbor:"df" json:"df"`
	ValidUntil        string `cbor:"du" json:"du"`
	CertificateID     string `cbor:"ci" json:"ci"`
}
