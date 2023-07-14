package interpret

import (
	"fmt"
	"log"
)

type Exception struct {
	Name    string
	Message string
}

func (e *Exception) Error() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Message)
}

func (e *Exception) Call() interface{} {
	log.Fatalln(e.Error())
	return nil
}

func NewException(name string, message string) *Exception {
	return &Exception{name, message}
}

func Raise(name string, message string) interface{} {
	e := NewException(name, message)
	return e.Call()
}
