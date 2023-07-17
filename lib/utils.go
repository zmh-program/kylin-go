package lib

func InArray[T comparable](value T, array []T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
