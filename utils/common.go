package utils

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
