package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// ommited function for simplicity

// Export public key as a string in PEM format
func exportPubKeyAsPEMStr(pubkey *rsa.PublicKey) string {
	pubKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pubkey),
		},
	))
	return pubKeyPem
}

// Export private key as a string in PEM format
func exportPrivKeyAsPEMStr(privkey *rsa.PrivateKey) string {
	privKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privkey),
		},
	))
	return privKeyPem

}

// Save string to a file
func saveKeyToFile(keyPem, filename string) {
	pemBytes := []byte(keyPem)
	ioutil.WriteFile(filename, pemBytes, 0400)
}

func generateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	// This method requires a random number of bits.
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// The public key is part of the PrivateKey struct
	return privateKey, &privateKey.PublicKey
}

func main() {
	// Generate a 2048-bits key
	privateKey, publicKey := generateKeyPair(2048)

	fmt.Printf("Private key: %v\n", privateKey)
	fmt.Printf("Public Key: %v", publicKey)

	// Create PEM string
	privKeyStr := exportPrivKeyAsPEMStr(privateKey)
	pubKeyStr := exportPubKeyAsPEMStr(publicKey)

	fmt.Println(privKeyStr)
	fmt.Println(pubKeyStr)

	saveKeyToFile(privKeyStr, "privkey.pem")
	saveKeyToFile(pubKeyStr, "pubkey.pem")
}
