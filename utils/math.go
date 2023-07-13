package utils

import "math"

func Pow(x float64, y float64) float64 {
	return math.Pow(x, y)
}

func IsTypeArray(obj interface{}) bool {
	switch obj.(type) {
	case []interface{}:
		return true
	default:
		return false
	}
}
