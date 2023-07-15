package utils

import (
	"fmt"
	"os"
)

func Fatal(err ...interface{}) {
	message := fmt.Sprint(err...)
	_, _ = fmt.Fprintf(os.Stderr, message+"\n\n")
	os.Exit(1)
}

func Debug[T comparable](data T) T {
	fmt.Println(data)
	return data
}
