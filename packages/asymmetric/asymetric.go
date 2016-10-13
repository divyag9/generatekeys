package asymmetric

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

// GenerateKeys generates the private and public key for encryption/decryption
func GenerateKeys() (err error) {
	//Generate Private Key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	// Precompute Calculations that speed up private key operations in the future
	privateKey.Precompute()
	//Validate Private Key
	err = privateKey.Validate()
	//Public key address (of an RSA key)
	publicKey := &privateKey.PublicKey
	//Write keys to files
	err = SavePrivateKey(privateKey)
	err = SavePublicKey(publicKey)

	return
}

//SavePrivateKey writes the private key to a file after encoding to pem
func SavePrivateKey(privateKey *rsa.PrivateKey) (err error) {
	pemPriv := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		})
	err = ioutil.WriteFile("private.pem", pemPriv, 0644)

	return
}

//SavePublicKey writes the public key to a file after encoding to pem
func SavePublicKey(publicKey *rsa.PublicKey) (err error) {
	pubBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	pemPub := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubBytes,
		})
	err = ioutil.WriteFile("public.pem", pemPub, 0644)

	return
}
