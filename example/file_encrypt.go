package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kuromittsu/util_encrypt"
)

func file_encrypt() {

	fe := util_encrypt.NewFile(
		"./example/files/file_jpg.jpg", // path to file
		"enc",                          // output file extension
		"encrypted",                    // output file name
		"",                             // output folder
	)

	if encResult, err := fe.Encrypt(
		[]byte(key), // encrypt key
		false,       // delete input file after success encrypt
		true,        // auto save / write to file
	); err != nil {
		fmt.Printf("error while encrypt file | %v \n", err)
		os.Exit(0)
	} else {
		// fmt.Printf("encResult.GetByte(): %v\n", encResult.GetByte())
		fmt.Printf("encResult.GetPath(): %v\n", encResult.GetPath())
		fmt.Printf("encResult.GetExt(): %v\n", encResult.GetExt())
	}

	time.Sleep(3000)

	fd := util_encrypt.NewFile(
		"./example/files/encrypted.enc", // path to file
		"jpg",                           // output file extension
		"decrypted",                     // output file name
		"",                              //output folder
	)

	if decResult, err := fd.Decrypt(
		[]byte(key), // decrypt key
		false,       // delete input file after success decrypt
		true,        // auto save / write to file
	); err != nil {
		fmt.Printf("error while decrypt file | %v \n", err)
		os.Exit(0)
	} else {
		// fmt.Printf("decResult.GetByte(): %v\n", decResult.GetByte())
		fmt.Printf("decResult.GetPath(): %v\n", decResult.GetPath())
	}
}
