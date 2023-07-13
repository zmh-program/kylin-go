package module

import "kylin/include"

type TokenType int

type Token interface {
	GetType() TokenType
	GetValue() string
}

type Interpreter interface {
	GetVariable(name string) interface{}
	SetVariable(name string, value interface{})
	GetScope() *include.Scope
	SetScope(scope *include.Scope)
	Next() Token
	Peek() Token
	Skip()
	GetCurrentLine() int
	GetCurrentColumn() int
	Run()
}
