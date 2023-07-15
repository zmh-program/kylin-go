package utils

import "unicode"

func IsLetter(n byte) bool {
	return unicode.IsLetter(rune(n))
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

func IsString(n byte) bool {
	return n == '"' || n == '\''
}

func ToBool(data interface{}) bool {
	if data == nil {
		return false
	}
	switch data.(type) {
	case bool:
		return data.(bool)
	case int:
		return data.(int) != 0
	case float64:
		return data.(float64) != 0
	case string:
		return data.(string) != ""
	case []interface{}:
		return len(data.([]interface{})) != 0
	case map[string]interface{}:
		return len(data.(map[string]interface{})) != 0
	default:
		return true
	}
}
