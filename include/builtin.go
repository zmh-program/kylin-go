package include

import (
	"kylin/utils"
	"strings"
)

type KyFunc func(...interface{})

func NewGlobalScope() *Scope {
	scope := NewScope(nil)

	{
		scope.Set("print", Print)
		scope.Set("sum", Sum)
		scope.Set("max", Max)
		scope.Set("min", Min)
		scope.Set("len", Len)
		scope.Set("type", Type)
		scope.Set("abs", Abs)
		scope.Set("all", All)
		scope.Set("any", Any)
		scope.Set("split", Split)
	}

	return scope
}

func Print(obj ...interface{}) interface{} {
	var str []string
	for _, v := range obj {
		str = append(str, utils.ToString(v))
	}
	println(strings.Join(str, " "))

	return nil
}

func Sum(args ...interface{}) interface{} {
	var sum float64
	for _, v := range args {
		switch v.(type) {
		case int:
			sum += float64(v.(int))
		case float64:
			sum += v.(float64)
		}
	}
	return sum
}

func Max(args ...interface{}) interface{} {
	if len(args) == 1 && utils.IsTypeArray(args[0]) {
		return Max(args[0].([]interface{})...)
	}

	var max float64
	for _, v := range args {
		switch v.(type) {
		case int:
			if float64(v.(int)) > max {
				max = float64(v.(int))
			}
		case float64:
			if v.(float64) > max {
				max = v.(float64)
			}
		}
	}
	return max
}

func Min(args ...interface{}) interface{} {
	if len(args) == 1 && utils.IsTypeArray(args[0]) {
		return Min(args[0].([]interface{})...)
	}

	min := args[0].(float64)
	for _, v := range args {
		switch v.(type) {
		case int:
			if float64(v.(int)) < min {
				min = float64(v.(int))
			}
		case float64:
			if v.(float64) < min {
				min = v.(float64)
			}
		}
	}
	return min
}

func Len(obj interface{}) interface{} {
	switch obj.(type) {
	case string:
		return len(obj.(string))
	case []interface{}:
		return len(obj.([]interface{}))
	}
	return 0
}

func Type(obj interface{}) interface{} {
	switch obj.(type) {
	case int:
		return "int"
	case float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	case []interface{}:
		return "array"
	case map[string]interface{}:
		return "map"
	default:
		return "object"
	}
}

func Abs(num float64) interface{} {
	if num < 0 {
		return -num
	}
	return num
}

func All(args ...interface{}) interface{} {
	for _, v := range args {
		if !v.(bool) {
			return false
		}
	}
	return true
}

func Any(args ...interface{}) interface{} {
	for _, v := range args {
		if v.(bool) {
			return true
		}
	}
	return false
}

func Split(str string, sep string) interface{} {
	return strings.Split(str, sep)
}
