package interpret

import (
	"fmt"
	"kylin/utils"
)

type Exception struct {
	Name    string
	Message string
}

func (e *Exception) Error() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Message)
}

func (e *Exception) Call() interface{} {
	utils.Fatal(e.Error())
	return nil
}

func NewException(name string, message string) *Exception {
	return &Exception{name, message}
}

func Raise(name string, message string) interface{} {
	e := NewException(name, message)
	return e.Call()
}

func (i *Interpreter) SetException(name string, message string) {
	i.err = NewException(name, message)
}

func (i *Interpreter) GetException() interface{} {
	return i.err
}

func (i *Interpreter) IsException() bool {
	return i.err != nil
}

func (i *Interpreter) ClearException() {
	i.err = nil
}

func (i *Interpreter) IsCaching() bool {
	return i.caching
}

func (i *Interpreter) SetCaching(caching bool) {
	i.caching = caching
}

func (i *Interpreter) Throw(name string, message string) interface{} {
	mes := fmt.Sprintf("%s at line %d, column %d", message, i.GetCurrentLine(), i.GetCurrentColumn())
	if i.IsCaching() {
		i.SetException(name, mes)
		return nil
	} else {
		return Raise(name, mes)
	}
}

func (i *Interpreter) ThrowError(message string) interface{} {
	return i.Throw("Error", message)
}
