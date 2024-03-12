package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
)

// Reads plaintext input from stdin
// Symmetrically encrypts the plaintext using a 32-byte [cryptographically]
// random generated key (AES-256 GCM)

func main() {
	fmt.Println("Enter plaintext to encrypt: ")

	userInput := make([]byte, 1024)
	_, err := os.Stdin.Read(userInput)
	if err != nil {
		fmt.Println("Error reading from stdin")
	}

	//generate random AES-256 key (32 bytes)
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		fmt.Println("Error generating random key")
	}

	// create new cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating new cipher")
	}

	// wrap 128-bit block in GCM (Galous Counter Mode)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Error creating new GCM")
	}

	// generate random nonce, has to be 12 bytes
	nonce := make([]byte, 12)
	if _, err := rand.Read(nonce); err != nil {
		fmt.Println("Error generating random nonce")
	}

	encrypted := aesGCM.Seal(
		nil,
		nonce,
		userInput,
		nil,
	)

	decrypted, err := aesGCM.Open(
		nil,
		nonce,
		encrypted,
		nil,
	)

	if err != nil {
		fmt.Println("Error decrypting ciphertext")
	}

	fmt.Printf("decrypted: %s\n", decrypted)
}
