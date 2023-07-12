package utils

func Pow(x float64, y float64) float64 {
	return float64(int64(x) ^ int64(y))
}
