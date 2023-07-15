package lib

import (
	"fmt"
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
		Fatal(fmt.Sprintf("File is not exist: %s", module))
	}
	return data
}

func Fatal(err ...interface{}) {
	message := fmt.Sprint(err...)
	_, _ = fmt.Fprintf(os.Stderr, message+"\n\n")
	os.Exit(1)
}

func Debug[T comparable](data T) T {
	fmt.Println(data)
	return data
}
