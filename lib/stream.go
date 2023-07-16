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
	indent := strings.Repeat(" ", indentLevel)
	res := "\n"

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		res += fmt.Sprintf("%s%s: %v\n", indent, field.Name, format(value.Interface(), indentLevel+1))
	}
	return res
}

func formatMap(m interface{}, indentLevel int) string {
	v := reflect.ValueOf(m)
	keys := v.MapKeys()
	indent := strings.Repeat(" ", indentLevel)
	res := ""

	for _, key := range keys {
		value := v.MapIndex(key)

		res += fmt.Sprintf("%s%v: %v\n", indent, format(key.Interface(), indentLevel+1), format(value.Interface(), indentLevel+1))
	}
	return res
}

func formatSlice(s interface{}, indentLevel int) string {
	v := reflect.ValueOf(s)
	indent := strings.Repeat(" ", indentLevel)
	res := "\n"

	for i := 0; i < v.Len(); i++ {
		value := v.Index(i)

		res += fmt.Sprintf("%s[%d]: %v\n", indent, i, format(value.Interface(), indentLevel+1))
	}
	return res
}

func format(v interface{}, indentLevel int) string {
	res := ""

	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Array:
		res += formatSlice(v, indentLevel)
	case reflect.Map:
		res += formatMap(v, indentLevel)
	case reflect.Struct:
		res += formatStruct(v, indentLevel)
	default:
		res += fmt.Sprintf("%v", v)
	}
	return res
}

func Print(s interface{}) {
	fmt.Println(format(s, 0))
}
