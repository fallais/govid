# GoVID

Golang tool that reads, decodes and validates the QRCode of a **Digital COVID Certificate (DCC)**.

## How it works ?

The whole Covid QRCode can be seen as a PKI. The main component, called the **Digital COVID Certificate Gateway (DCCG)**, hold all the public keys and the **validation rules**.

Each state, which is called a **Member State**, maintains one or many **Country Signing Certificate Authority (CSCA)**, and has to publish the private keys used in the signing process to the DCCG.

This CSCA issue **Document Signer Certificates (DSCs)**, these are public keys used by the **Document Signers**.


![Coop](https://github.com/fallais/govid/blob/master/assets/mecanism_overview.png)
