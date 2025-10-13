package util_encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

// AES encrypt
// encrypt raw text to encrypted version using key and iv (initialization vector)
func aesEncrypt(text, key, iv string) (string, error) {

	// Validate text
	if err := validateTextLength(text, "text"); err != nil {
		return "", err
	}

	// Validate key
	if err := validateKey(key); err != nil {
		return "", err
	}

	// Validate iv
	if err := validateIv(iv); err != nil {
		return "", err
	}

	// Creating an AES cipher block
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Create ciphertext with the same length as the text
	ciphertext := make([]byte, len(text))

	// Create an encrypted stream with a fixed IV
	stream := cipher.NewCFBEncrypter(block, []byte(iv))

	// Encrypting text
	stream.XORKeyStream(ciphertext, []byte(text))

	// Returns the result in hexadecimal format
	return hex.EncodeToString(ciphertext), nil
}

// AES decrypt
// decrypt encrypted text to raw text using key and iv (initialization vector)
func aesDecrypt(encryptedText, key, iv string) (string, error) {

	// Validate text
	if err := validateTextLength(encryptedText, "encrypted text"); err != nil {
		return "", err
	}

	// Validate key
	if err := validateKey(key); err != nil {
		return "", err
	}

	// Validate iv
	if err := validateIv(iv); err != nil {
		return "", err
	}

	// Decode ciphertext from hexadecimal
	ciphertext, err := hex.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	// Creating an AES cipher block
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Create a decryption stream with a fixed IV
	stream := cipher.NewCFBDecrypter(block, []byte(iv))

	// Decrypting ciphertext
	plainText := make([]byte, len(ciphertext))
	stream.XORKeyStream(plainText, ciphertext)

	return string(plainText), nil
}
