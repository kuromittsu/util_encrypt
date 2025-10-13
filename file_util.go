package util_encrypt

import (
	"path/filepath"
	"strings"
)

func getNormalizeDir(inputPath, dir string) string {

	normalizedDir := filepath.ToSlash(dir)
	if strings.HasPrefix(inputPath, "./") && !strings.HasPrefix(normalizedDir, "./") {
		normalizedDir = "./" + normalizedDir
	}
	if !strings.HasSuffix(normalizedDir, "/") {
		normalizedDir += "/"
	}
	return normalizedDir
}

func filterExt(ext string) string {

	if strings.HasPrefix(".", ext) {
		return ext[1:]
	}
	return "." + ext
}
