package main

import (
	"fmt"
	"os"

	"github.com/kuromittsu/util_encrypt"
)

var key string = "BOAlNWXHfwKdiQsZIVmEBqGTNrfmNMpe"
var iv string = "EhJkLKhozRriKdso"

func encrypt() {

	text := "Alfiras"

	fmt.Printf("text: %v \n", text)

	encryptedText, err := util_encrypt.AesEncrypt(text, key, iv)
	if err != nil {
		fmt.Printf("error while encrypt text | %v \n", err)
		os.Exit(0)
	}

	fmt.Printf("encrypted text (result): %v\n", encryptedText)

	decryptedText, err := util_encrypt.AesDecrypt(encryptedText, key, iv)
	if err != nil {
		fmt.Printf("error while encrypt text | %v \n", err)
		os.Exit(0)
	}

	fmt.Printf("decrypted text (result): %v\n", decryptedText)
	fmt.Printf("decrypted text == text: %v\n", decryptedText == text)
}
