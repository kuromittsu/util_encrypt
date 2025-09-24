package util_encrypt

func AesEncrypt(text, key, iv string) (string, error) {

	return aesEncrypt(text, key, iv)
}

func AesDecrypt(encryptedText, key, iv string) (string, error) {

	return aesDecrypt(encryptedText, key, iv)
}
