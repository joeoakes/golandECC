package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Generate an ECC private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// Generate a corresponding public key
	publicKey := &privateKey.PublicKey

	// Data to be signed
	data := []byte("Hello, ECC!")

	// Sign the data using the private key
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashData(data))
	if err != nil {
		fmt.Println("Error signing data:", err)
		return
	}

	// Verify the signature using the public key
	if ecdsa.Verify(publicKey, hashData(data), r, s) {
		fmt.Println("Signature is valid.")
	} else {
		fmt.Println("Signature is invalid.")
	}
}

// Helper function to hash data
func hashData(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
