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
