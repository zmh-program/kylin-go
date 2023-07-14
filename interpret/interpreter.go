package interpret

import (
	"kylin/include"
	"kylin/module"
	"kylin/utils"
)

type Interpreter struct {
	lexer  *Lexer
	scope  *include.Scope
	buffer interface{}
	ret    interface{}
	module *module.Manager
}

func NewInterpreter(path string, parent *include.Scope) *Interpreter {
	data := utils.ReadKylinFile(path)
	return &Interpreter{
		lexer:  NewLexer(data),
		scope:  include.NewScope(parent),
		module: module.NewManager(),
	}
}

func (i *Interpreter) GetVariable(name string) interface{} {
	return i.scope.Get(name)
}

func (i *Interpreter) SetVariable(name string, value interface{}) {
	i.scope.Set(name, value)
}

func (i *Interpreter) GetScope() *include.Scope {
	return i.scope
}

func (i *Interpreter) SetScope(scope *include.Scope) {
	i.scope = scope
}

func (i *Interpreter) Next() Token {
	return i.lexer.Next()
}

func (i *Interpreter) Peek() Token {
	return i.lexer.Peek()
}

func (i *Interpreter) Skip() {
	i.lexer.Skip()
}

func (i *Interpreter) GetCurrentLine() int {
	return i.lexer.line
}

func (i *Interpreter) GetCurrentColumn() int {
	return i.lexer.column
}

func (i *Interpreter) GetNextPtr() *Token {
	return i.lexer.GetNextPtr()
}

func (i *Interpreter) Expr(token *Token) interface{} {
	switch token.Type {
	case Integer:
		return i.CountCall(token)
	case Float:
		return i.CountCall(token)
	case True:
		return true
	case False:
		return false
	case Null:
		return nil
	case String:
		return token.Value
	case Identifier:
		if i.AssignCall(token) {
			return nil
		}
		if check, resp := i.FunctionCall(token); check {
			return resp
		}
		return i.CountCall(token)
	case LeftParenthesis:
		return i.ParenthesisCall()
	case LeftBracket:
		return i.ReadArray()
	case LeftBrace:
		return i.ReadObject()
	case Function:
		return i.ReadFunction()
	case Return:
		return i.SetReturn(i.ExprNext())
	case Subtraction:
		return -i.ExprNext().(float64)
	case If:
		return i.ConditionCall()
	case EOF:
		return nil
	}
	return token
}

func (i *Interpreter) GetBuffer() interface{} {
	return i.buffer
}

func (i *Interpreter) SetBuffer(token interface{}) {
	i.buffer = token
}

func (i *Interpreter) GetReturn() interface{} {
	return i.ret
}

func (i *Interpreter) SetReturn(token interface{}) interface{} {
	i.ret = token
	return token
}

func (i *Interpreter) IsEnd() bool {
	return i.lexer.IsEnd()
}

func (i *Interpreter) ExprNext() interface{} {
	res := i.Expr(i.GetNextPtr())
	i.SetBuffer(res)
	return res
}

func (i *Interpreter) Run() interface{} {
	for !i.IsEnd() {
		i.ExprNext()
	}
	return i.GetReturn()
}
