# GoVID

**Golang** tool that decodes and validates the Covid QRCode, which is formerly called an **EU Digital COVID Certificate (EUDCC)** (previously a Digital Green Certificate...).

## Usage

```go
go run main.go decode -dcc myImage.png
```

## How the COVID QRCode works ?

### Overview of the mecansim

The whole Covid QRCode system can be described as a PKI. The main component, called the **Digital COVID Certificate Gateway (DCCG)**, holds all the public keys and the **validation rules**.

Each state, which is called a **Member State**, maintains one or many **Country Signing Certificate Authority (CSCA)**, and has to publish the private keys used in the signing process to the DCCG.

This CSCA issue **Document Signer Certificates (DSCs)**, these are public keys used by the **Document Signers**.

![Coop](https://github.com/fallais/govid/blob/master/assets/mecanism_overview.png)

### The digital certificate

As far as I understood, the schema of the DCC is defined by the EHN. You can find [more information here](https://github.com/ehn-dcc-development/hcert-spec). The schema that seems to be used by all the *Member States* is located in [this repository](https://github.com/ehn-dcc-development/ehn-dcc-schema)

Basically, the DCC is a QRCode, which wraps a CBOR Web Token (signed, so it is called a COSE), encoded in base45. Note that this encoded content starts with `HC1:` or `LT1:`.

### Business rules

The DCC must then be validated against a list of **Business Rules**. These are based on **JsonLogic** (https://jsonlogic.com) which is called **CertLogic**. Guidelines can be [found here](ttps://github.com/eu-digital-green-certificates/dgc-business-rules-testdata).

From what I understood by *"reversing"* the code of **TAC Verif (TousAntiCovid Verif)**, the application used by the people who verify your QRCode, use [hard-coded rules](https://gitlab.inria.fr/tousanticovid-verif/tousanticovid-verif-android/-/blob/master/app/src/main/assets/sync/sync_rules.json).

I expected that the rules should have been loaded when the mobile-app starts, but as of now, I can't confirm. That means that every changes of the governement rules implies a app update ?
