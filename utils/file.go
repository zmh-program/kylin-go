package utils

import (
	"os"
	"strings"
)

func ReadFile(path string) (string, error) {
	if !strings.HasSuffix(path, ".ky") {
		path += ".ky"
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
