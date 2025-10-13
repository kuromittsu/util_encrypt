package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kuromittsu/util_encrypt"
)

func file_encrypt() {

	fe := util_encrypt.NewFile(
		"./files/file_jpg.jpg", // path to file
		"enc",                  // output file extension
		"encrypted",            // output file name
		"",                     // output folder
	)

	if err := fe.Encrypt(
		[]byte(key), // encrypt key
		false,       // delete input file after success encrypt
	); err != nil {
		fmt.Printf("error while encrypt file | %v \n", err)
		os.Exit(0)
	}

	time.Sleep(3000)

	fd := util_encrypt.NewFile(
		"./files/encrypted.enc", // path to file
		"jpg",                   // output file extension
		"decrypted",             // output file name
		"",                      //output folder
	)

	if err := fd.Decrypt(
		[]byte(key), // decrypt key
		false,       // delete input file after success decrypt
	); err != nil {
		fmt.Printf("error while decrypt file | %v \n", err)
		os.Exit(0)
	}
}
