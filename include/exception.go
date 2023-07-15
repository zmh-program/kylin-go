package include

import (
	"fmt"
	"kylin/utils"
)

type Exception struct {
	Name    string
	Message string
	Line    int
	Column  int
}

func (e *Exception) Error() string {
	return fmt.Sprintf("%s: %s at line %d, column %d", e.Name, e.Message, e.Line, e.Column)
}

func (e *Exception) Call() interface{} {
	utils.Fatal(e.Error())
	return nil
}

func (e *Exception) Repr() string {
	return fmt.Sprintf("%s(type=\"%s\", message=\"%s\", line=%d, column=%d)", e.Name, e.Name, e.Message, e.Line, e.Column)
}

func NewException(name string, message string, line int, column int) *Exception {
	return &Exception{name, message, line, column}
}

func Raise(name string, message string, line int, column int) interface{} {
	e := NewException(name, message, line, column)
	return e.Call()
}
