package utils

import "fmt"

func Output[T comparable](data T) T {
	fmt.Println(data)
	return data
}
