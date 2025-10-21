package util_encrypt

type fileEncryptResult struct {
	path string
	ext  string

	chipherText []byte
}

func (f *fileEncryptResult) GetByte() []byte {

	return f.chipherText
}

func (f *fileEncryptResult) GetPath() string {

	return f.path
}

func (f *fileEncryptResult) GetExt() string {

	return f.ext
}

// ===

type fileDecryptResult struct {
	path string

	plainText []byte
}

func (f *fileDecryptResult) GetByte() []byte {

	return f.plainText
}

func (f *fileDecryptResult) GetPath() string {

	return f.path
}
