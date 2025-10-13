package util_encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type file struct {
	inputPath  string
	outputPath string
	filename   string
	extension  string
}

func newFile(inputPath, extensionFile, outputFilename, outputPath string) *file {

	// base name
	basename := filepath.Base(inputPath)

	// dir
	dir := filepath.Dir(inputPath)
	if len(outputPath) != 0 {
		dir = outputPath
	}

	// extension
	ext := filepath.Ext(basename)
	if len(extensionFile) != 0 {
		ext = filterExt(extensionFile)
	}

	// filename
	filename := basename[:len(basename)-len(ext)]
	if len(outputFilename) != 0 {
		filename = outputFilename
	}

	normalizedDir := getNormalizeDir(inputPath, dir)

	return &file{
		inputPath:  inputPath,
		outputPath: normalizedDir,
		extension:  ext,
		filename:   filename,
	}
}

// functions

func (f *file) getOutput() string {

	return f.outputPath + f.filename + f.extension
}

func (f *file) Encrypt(key []byte, deleteOld bool) error {

	if err := validateKey(string(key)); err != nil {
		return err
	}

	if len(f.inputPath) == 0 {
		return errors.New("input path is required")
	}

	// read file
	plaintext, err := os.ReadFile(f.inputPath)
	if err != nil {
		return fmt.Errorf("error while read file: %w", err)
	}

	// create blok cipher AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("error while creating cipher AES: %w", err)
	}

	// create GCM (Galois/Counter Mode)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("error while creating GCM: %w", err)
	}

	// create random nonce/IV (Initialization Vector)
	// nonce is random value must be unique for every encrypt operation.
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("error while creating nonce: %w", err)
	}

	// Enkripsi data
	// encrypt data
	// output = Nonce + Ciphertext (Ciphertext already included tag otentikasi GCM)
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	// write encrypted data to output file
	if err := os.WriteFile(f.getOutput(), ciphertext, 0644); err != nil {
		return fmt.Errorf("failed write encrypted file: %w", err)
	}

	if deleteOld {
		return os.Remove(f.inputPath)
	}

	return nil
}

func (f *file) Decrypt(key []byte, deleteOld bool) error {

	if len(f.inputPath) == 0 {
		return errors.New("input path is required")
	}

	if err := validateKey(string(key)); err != nil {
		return err
	}

	// read encrypted data (including nonce) from input file
	ciphertextWithNonce, err := os.ReadFile(f.inputPath)
	if err != nil {
		return fmt.Errorf("error while read encrypted file: %w", err)
	}

	// create block cipher AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("error while creating cipher AES: %w", err)
	}

	// create GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("error while creating GCM: %w", err)
	}

	// separate nonce and ciphertext
	nonceSize := aesGCM.NonceSize()
	if len(ciphertextWithNonce) < nonceSize {
		return fmt.Errorf("file too short to decrypt")
	}

	nonce := ciphertextWithNonce[:nonceSize]
	ciphertext := ciphertextWithNonce[nonceSize:]

	// decrypting data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		// it will fail if wrong key or encrypted file corrupt (failed authentication)
		return fmt.Errorf("decrypt error (wrong key or corrupt data): %w", err)
	}

	// write decrypted data to output file
	if err := os.WriteFile(f.getOutput(), plaintext, 0644); err != nil {
		return fmt.Errorf("failed write decrypted file: %w", err)
	}

	if deleteOld {
		return os.Remove(f.inputPath)
	}

	return nil
}
