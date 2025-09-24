package util_encrypt

import (
	"errors"
	"fmt"
)

func validateTextLength(value, text string) error {
	if len(value) == 0 {
		return fmt.Errorf("%s is required", text)
	}
	return nil
}

// Validate length of key (must either 16 or 24 or 32 byte)
func validateKeyLength(key string) error {

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return errors.New("key length must be 16, 24, atau 32 bytes")
	}
	return nil
}

// Validate length of iv (must fixed 16 byte)
func validateIvLength(iv string) error {

	if len(iv) != 16 {
		return errors.New("iv length must fixed 16 bytes")
	}
	return nil
}
