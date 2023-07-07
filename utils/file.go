package utils

import (
	"os"
	"strings"
)

func ReadFile(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func ReadKylinFile(module string) string {
	if !strings.HasSuffix(module, ".ky") {
		module += ".ky"
	}
	data, err := ReadFile(module)
	if err != nil {
		panic(err)
	}
	return data
}
