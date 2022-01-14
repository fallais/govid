# GoVID

**Golang** tool that decodes and validates the Covid QRCode, which is formerly called an **EU Digital COVID Certificate (EUDCC)** (previously a Digital Green Certificate...).

## Usage

```go
go run main.go decode --qrcode qrcode.png
```

## How the COVID QRCode works ?

### Overview of the mecansim

The whole Covid QRCode system can be described as a PKI. The main component, called the **Digital COVID Certificate Gateway (DCCG)**, holds all the public keys and the **validation rules**.

Each state, which is called a **Member State**, maintains one or many **Country Signing Certificate Authority (CSCA)**, and has to publish the private keys used in the signing process to the DCCG.

This CSCA issue **Document Signer Certificates (DSCs)**, these are public keys used by the **Document Signers**. My guess is that what they call a Document Signer is a center of vaccination, or a pharmacy, etc.. To be verified.

![Coop](https://github.com/fallais/govid/blob/master/assets/mecanism_overview.png)

### The digital certificate

As far as I understood, the schema of the DCC is defined by the EHN. You can find [more information here](https://github.com/ehn-dcc-development/hcert-spec). The schema that seems to be used by all the *Member States* is located in [this repository](https://github.com/ehn-dcc-development/ehn-dcc-schema)

Basically, the DCC is a QRCode, which is a **CBOR Web Token (CWT)** (more information [here](https://datatracker.ietf.org/doc/html/rfc8949)), encoded in **base45**.  

So when reading the QRCode, we obtain a string starting with `HC1:`, for exemple:

```
HC1:6BF$RDA-SMAHN-H6SKJPT.-7G2TZ97+S8Y4CXEJW.TFJTXG4DCESYCTSJ2TP$2G*P5Y
IJ6S4HZ6SH9+2QH/56SP.E5BQ95ZM376ZIE7PM1VE.Q6T:H:1NPK9/1AAJ1KK9%OC+G9QJP
NF67J6QW6D9RY466PPXY0E7J7UJQWT.+S1QDC8CI6C6XIO$9KZ56DE/.QC$Q3J62:6LZ6O5
9++9-G9+E93ZM$96PZ6+Q6X46+E54A9NF625F646L+9AKPCPP7ZMN27QW6ZOQ.NE/E2$4JY
/K9:K4F7D*G2SV /KF-KXIN2SV6AL3*I**GYZQVG9YJC/HLIJLKNF8JF1727WLTPL 6KNYM
B26REDFAQOQUX9RGQR.RE1G6*%3PNNPXSAXB*LM.NEMUN50I++VKZ39GTLV75+UJ+0W8J:A
T0F6 WAKWVV7U*V7J:VM30AGRK2
```

We need to remove the prefix `HC1:` and what we now have is a string encoded with base45 (base45 is an [on-going project](https://datatracker.ietf.org/doc/draft-faltstrom-base45/)).

If we decode this string, we obtain bytes that are compressed with **zlib**.

We need to decompress it and we obtain the CBOR.

### Business rules

The DCC must then be validated against a list of **Business Rules**. These are based on **JsonLogic** (https://jsonlogic.com) which is called **CertLogic**. Guidelines can be [found here](ttps://github.com/eu-digital-green-certificates/dgc-business-rules-testdata).

From what I understood by *"reversing"* the code of **TAC Verif (TousAntiCovid Verif)**, the application used by the people who verify your QRCode, use [hard-coded rules](https://gitlab.inria.fr/tousanticovid-verif/tousanticovid-verif-android/-/blob/master/app/src/main/assets/sync/sync_rules.json).

I expected that the rules should have been loaded when the mobile-app starts, but as of now, I can't confirm. That means that every changes of the governement rules implies a app update ?
