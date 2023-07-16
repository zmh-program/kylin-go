package lib

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func ReadFile(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func ReadKylinFile(module string) string {
	if !strings.HasSuffix(module, ".ky") {
		module += ".ky"
	}
	data, err := ReadFile(module)
	if err != nil {
		Fatal(fmt.Sprintf("File is not exist: %s", module))
	}
	return data
}

func Fatal(err ...interface{}) {
	message := fmt.Sprint(err...)
	_, _ = fmt.Fprintf(os.Stderr, message+"\n\n")
	os.Exit(1)
}

func Debug[T comparable](data T) T {
	fmt.Println(data)
	return data
}

func formatStruct(s interface{}, indentLevel int) string {
	v := reflect.ValueOf(s)
	t := v.Type()
	indent := strings.Repeat(" ", indentLevel*4)
	res := ""

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if value.Kind() == reflect.Struct {
			res += fmt.Sprintf("%s%s:\n", indent, field.Name)
			res += formatStruct(value.Interface(), indentLevel+1)
		} else {
			res += fmt.Sprintf("%s%s: %v\n", indent, field.Name, value.Interface())
		}
	}
	return res
}

func formatMap(m interface{}, indentLevel int) string {
	v := reflect.ValueOf(m)
	keys := v.MapKeys()
	indent := strings.Repeat(" ", indentLevel*4)
	res := ""

	for _, key := range keys {
		value := v.MapIndex(key)

		if value.Kind() == reflect.Map {
			res += fmt.Sprintf("%s%v:\n", indent, key.Interface())
			res += formatMap(value.Interface(), indentLevel+1)
		} else {
			res += fmt.Sprintf("%s%v: %v\n", indent, key.Interface(), value.Interface())
		}
	}
	return res
}

func formatSlice(s interface{}, indentLevel int) string {
	v := reflect.ValueOf(s)
	indent := strings.Repeat(" ", indentLevel*4)
	res := ""

	for i := 0; i < v.Len(); i++ {
		value := v.Index(i)

		if value.Kind() == reflect.Map {
			res += fmt.Sprintf("%s[%d]:\n", indent, i)
			res += formatMap(value.Interface(), indentLevel+1)
		} else if value.Kind() == reflect.Struct {
			res += fmt.Sprintf("%s[%d]:\n", indent, i)
			res += formatStruct(value.Interface(), indentLevel+1)
		} else {
			res += fmt.Sprintf("%s[%d]: %v\n", indent, i, value.Interface())
		}
	}
	return res
}

func format(v interface{}, indentLevel int) string {
	if v == nil {
		return "null"
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Struct:
		return formatStruct(v, indentLevel)
	case reflect.Slice, reflect.Array:
		return formatSlice(v, indentLevel)
	case reflect.Map:
		return formatMap(v, indentLevel)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func Print(s interface{}) {
	fmt.Println(format(s, 0))
}
