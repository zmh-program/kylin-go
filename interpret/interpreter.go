package interpret

import (
	"kylin/i18n"
	"kylin/include"
	"kylin/lib"
	"kylin/module"
)

type KyRuntime struct {
	lexer   *Lexer
	scope   *include.Scope
	buffer  interface{}
	err     interface{}
	ret     interface{}
	caching bool
	module  *module.Manager
	i18n    *i18n.Manager
}

func NewRuntime(path string, parent *include.Scope, i18n *i18n.Manager) *KyRuntime {
	data := lib.ReadKylinFile(path)
	return &KyRuntime{
		lexer:  NewLexer(data, i18n),
		scope:  include.NewScope(parent),
		module: module.NewManager(),
		i18n:   i18n,
	}
}

func (i *KyRuntime) GetVariable(name string) interface{} {
	if val, ok := i.scope.Get(name); ok {
		return val
	}
	i.Throw("ReferenceError", "Variable "+name+" not defined")
	return nil
}

func (i *KyRuntime) SetVariable(name string, value interface{}) {
	i.scope.Set(name, value)
}

func (i *KyRuntime) GetScope() *include.Scope {
	return i.scope
}

func (i *KyRuntime) SetScope(scope *include.Scope) {
	i.scope = scope
}

func (i *KyRuntime) GetModule() *module.Manager {
	return i.module
}

func (i *KyRuntime) SetModule(module *module.Manager) {
	i.module = module
}

func (i *KyRuntime) GetI18n() *i18n.Manager {
	return i.i18n
}

func (i *KyRuntime) SetI18n(i18n *i18n.Manager) {
	i.i18n = i18n
}

func (i *KyRuntime) Next() Token {
	return i.lexer.Next()
}

func (i *KyRuntime) Peek() Token {
	return i.lexer.Peek()
}

func (i *KyRuntime) Skip() {
	i.lexer.Skip()
}

func (i *KyRuntime) GetCurrentLine() int {
	return i.lexer.line
}

func (i *KyRuntime) GetCurrentColumn() int {
	return i.lexer.column
}

func (i *KyRuntime) GetNextPtr() *Token {
	return i.lexer.GetNextPtr()
}

func (i *KyRuntime) Expr(token *Token) interface{} {
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
	case While:
		return i.WhileCall()
	case For:
		return i.ForCall()
	case Try:
		return i.ExceptionCall()
	case Use:
		return i.UseCall()
	case EOF:
		return nil
	}
	return token
}

func (i *KyRuntime) GetBuffer() interface{} {
	return i.buffer
}

func (i *KyRuntime) SetBuffer(token interface{}) {
	i.buffer = token
}

func (i *KyRuntime) GetReturn() interface{} {
	return i.ret
}

func (i *KyRuntime) SetReturn(token interface{}) interface{} {
	i.ret = token
	return token
}

func (i *KyRuntime) IsEnd() bool {
	return !i.lexer.HasNext()
}

func (i *KyRuntime) ExprNext() interface{} {
	res := i.Expr(i.GetNextPtr())
	i.SetBuffer(res)
	return res
}

func (i *KyRuntime) Run() interface{} {
	for !i.IsEnd() {
		i.ExprNext()
	}
	return i.GetReturn()
}
