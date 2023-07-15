package include

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func NewGlobalScope() *Scope {
	scope := NewScope(nil)

	{
		scope.Set("print", Print)
		scope.Set("input", Input)
		scope.Set("str", String)
		scope.Set("int", Int)
		scope.Set("float", Float)
		scope.Set("bool", Bool)
		scope.Set("array", Array)
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
		scope.Set("read", Read)
		scope.Set("write", Write)
		scope.Set("shell", Shell)
		scope.Set("exit", Exit)
	}

	return scope
}

func Print(obj ...interface{}) interface{} {
	var str []string
	for _, v := range obj {
		str = append(str, String(v))
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

func IsTypeArray(obj interface{}) bool {
	switch obj.(type) {
	case []interface{}:
		return true
	default:
		return false
	}
}

func Max(args ...interface{}) interface{} {
	if len(args) == 1 && IsTypeArray(args[0]) {
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
	if len(args) == 1 && IsTypeArray(args[0]) {
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
		str = append(str, String(v))
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

func String(data interface{}) string {
	switch data.(type) {
	case string:
		return data.(string)
	case int:
		return fmt.Sprintf("%d", data.(int))
	case float64:
		return fmt.Sprintf("%f", data.(float64))
	case bool:
		return fmt.Sprintf("%v", data.(bool))
	case []interface{}:
		return fmt.Sprintf("%v", data.([]interface{}))
	case map[string]interface{}:
		return fmt.Sprintf("%v", data.(map[string]interface{}))
	case nil:
		return "null"
	case *Exception:
		return fmt.Sprintf(data.(*Exception).Repr())
	}
	return fmt.Sprintf("%v", data)
}

func Int(data interface{}) interface{} {
	switch data.(type) {
	case int:
		return data.(int)
	case float64:
		return int(data.(float64))
	case string:
		i, err := strconv.Atoi(data.(string))
		if err != nil {
			return 0
		}
		return i
	}
	return 0
}

func Float(data interface{}) interface{} {
	switch data.(type) {
	case float64:
		return data.(float64)
	case int:
		return float64(data.(int))
	case string:
		f, err := strconv.ParseFloat(data.(string), 64)
		if err != nil {
			return 0
		}
		return f
	}
	return 0
}

func Bool(data interface{}) interface{} {
	switch data.(type) {
	case bool:
		return data.(bool)
	case int:
		return data.(int) != 0
	case float64:
		return data.(float64) != 0
	case string:
		return data.(string) != ""
	}
	return false
}

func Array(data interface{}) []interface{} {
	switch data.(type) {
	case []interface{}:
		return data.([]interface{})
	case string:
		return []interface{}{data.(string)}
	case int:
		return []interface{}{data.(int)}
	case float64:
		return []interface{}{data.(float64)}
	case bool:
		return []interface{}{data.(bool)}
	case nil:
		return []interface{}{}
	}
	return []interface{}{}
}

func Shell(cmd interface{}) interface{} {
	out, err := exec.Command(cmd.(string)).Output()
	if err != nil {
		return ""
	}
	return string(out)
}

func Read(filename interface{}) interface{} {
	data, err := os.ReadFile(filename.(string))
	if err != nil {
		return ""
	}
	return string(data)
}

func Write(filename interface{}, data interface{}) {
	content := []byte(String(data))
	if err := os.WriteFile(filename.(string), content, 0644); err != nil {
		panic(err)
	}
}
