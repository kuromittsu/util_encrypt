# Util Encrypt

## Instal

```
go get github.com/kuromittsu/util_encrypt
```

## Generate Key & IV

you can generate key & iv using [rgen](https://github.com/kuromittsu/rgen)

## Available Methods

```go
// Text encryption & decryption
AesEncrypt(text, key, iv string) (string, error)
AesDecrypt(encryptedText, key, iv string) (string, error)

// File encryption & decryption
NewFile(inputPath, extensionFile, outputFilename, outputPath string) *file
```

## Usage

> You can use [rgen](https://github.com/kuromittsu/rgen) to generate key & iv

### Text

#### Encrypt Text

Convert raw text to encrypted version by key & iv

```go
// import ued "github.com/kuromittsu/util_encrypt"

key := "custom32byteskey0000000000000000"
iv := "custom16bytesiv0"
text := "Alfiras"

encryptedText, err := ued.AesEncrypt(text, key, iv)
if err != nil {
  fmt.Printf("error while encrypt text | %v \n", err)
  os.Exit(0)
}
fmt.Printf("encryptedText: %s", encryptedText)
```

result

```txt
encryptedText: de6eac0942b8fb
```

#### Decrypt Text

Convert encrypted text to raw text by key & iv

```go
// import ued "github.com/kuromittsu/util_encrypt"

// key & iv same as encrypt on top

decryptedText, err := util_encrypt.AesDecrypt(encryptedText, key, iv)
if err != nil {
  fmt.Printf("error while decrypt text | %v \n", err)
  os.Exit(0)
}
fmt.Printf("decryptedText: %s", decryptedText)
```

result

```txt
decryptedText: Alfiras
```

### File

#### Struct

```go
// *file

func(f *file) Encrypt(key []byte, deleteOld bool) error
func(f *file) Decrypt(key []byte, deleteOld bool) error
```

#### Encrypt File

```go
fe := util_encrypt.NewFile(
  "C:Temp/files/file_jpg.jpg", // path to file
  "enc",                       // output file extension
  "encrypted",                 // output file name
  "",                          // output folder
)
if err := fe.Encrypt(
  []byte(key), // encrypt key
  false,       // delete input file after success encrypt
); err != nil {
  fmt.Printf("error while encrypt file | %v \n", err)
  os.Exit(0)
}
```

#### Decrypt File

```go
fd := util_encrypt.NewFile(
  "C:Temp/files/encrypted.enc", // path to file
  "jpg",                        // output file extension
  "decrypted",                  // output file name
  "",                           // output folder
)

if err := fd.Decrypt(
  []byte(key), // decrypt key
  false,       // delete input file after success decrypt
); err != nil {
  fmt.Printf("error while decrypt file | %v \n", err)
  os.Exit(0)
}
```