package interpret

import "kylin/utils"

type Interpreter struct {
	lexer *Lexer
	scope *Scope
}

func NewInterpreter(module string, parent *Scope) *Interpreter {
	data := utils.ReadKylinFile(module)
	return &Interpreter{
		lexer: NewLexer(data),
		scope: NewScope(parent),
	}
}

func (i *Interpreter) GetVariable(name string) interface{} {
	return i.scope.Get(name)
}

func (i *Interpreter) SetVariable(name string, value interface{}) {
	i.scope.Set(name, value)
}

func (i *Interpreter) Expr(token *Token) interface{} {
	if token.Type == Number {
		return token.Value
	}
	if token.Type == Identifier {
		return i.GetVariable(token.Value)
	}
}
