# Util Encrypt

## Instal

```
go get github.com/kuromittsu/util_encrypt
```

## Generate Key & IV

you can generate key & iv using [rgen](https://github.com/kuromittsu/rgen)

## Available Methods

```
AesEncrypt(text, key, iv string) (string, error)
```

```
AesDecrypt(encryptedText, key, iv string) (string, error)
```

## Usage

see example