package utils

import (
	"fmt"
	"log"
)

func IsLetter(n byte) bool {
	return (n >= 'a' && n <= 'z') || (n >= 'A' && n <= 'Z')
}

func IsDigit(n byte) bool {
	return n >= '0' && n <= '9'
}

func IsRegularSymbol(n byte) bool {
	return n == '_' || n == '$'
}

func IsRegular(n byte) bool {
	return IsLetter(n) || IsRegularSymbol(n) || IsDigit(n)
}

func MustGet[T comparable](data T, err error) T {
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error: %s", err))
	}
	return data
}
