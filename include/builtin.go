package include

import (
	"fmt"
	"kylin/utils"
	"os"
	"strings"
	"time"
)

func NewGlobalScope() *Scope {
	scope := NewScope(nil)

	{
		scope.Set("print", Print)
		scope.Set("input", Input)
		scope.Set("sum", Sum)
		scope.Set("max", Max)
		scope.Set("min", Min)
		scope.Set("len", Len)
		scope.Set("type", Type)
		scope.Set("abs", Abs)
		scope.Set("all", All)
		scope.Set("any", Any)
		scope.Set("split", Split)
		scope.Set("join", Join)
		scope.Set("time", Time)
		scope.Set("timenano", TimeNano)
		scope.Set("sleep", Sleep)
		scope.Set("range", Range)
		scope.Set("exit", Exit)
	}

	return scope
}

func Print(obj ...interface{}) interface{} {
	var str []string
	for _, v := range obj {
		if v == nil {
			str = append(str, "null")
			continue
		}
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
		return "object"
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
		if v == nil {
			return false
		}
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

func Join(arr []interface{}, sep string) interface{} {
	var str []string
	for _, v := range arr {
		str = append(str, utils.ToString(v))
	}
	return strings.Join(str, sep)
}

func Time() interface{} {
	nano := time.Now().UnixNano()
	return float64(nano / int64(time.Millisecond))
}

func TimeNano() interface{} {
	return float64(time.Now().UnixNano())
}

func Sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func Range(args ...interface{}) interface{} {
	start := 0.
	end := 0.
	step := 1.

	switch len(args) {
	case 1:
		end = args[0].(float64)
	case 2:
		start, end = args[0].(float64), args[1].(float64)
	case 3:
		start, end, step = args[0].(float64), args[1].(float64), args[2].(float64)
	}

	var arr []interface{}
	for i := start; i < end; i += step {
		arr = append(arr, i)
		if step < 0 && i <= end {
			break
		} else if i >= end {
			break
		}

		if i+step >= end {
			arr = append(arr, i+step)
			break
		}
	}
	return arr
}

func Input(message ...interface{}) interface{} {
	fmt.Print(message...)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return nil
	}
	return input
}

func Exit(code int) {
	os.Exit(code)
}
