package interpret

import (
	"fmt"
	"kylin/utils"
	"strconv"
)

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
	switch token.Type {
	case Integer:
		return utils.MustGet(strconv.Atoi(token.Value))
	case Float:
		return utils.MustGet(strconv.ParseFloat(token.Value, 64))
	case String:
		return token.Value
	case Identifier:
		return i.scope.Get(token.Value)
	}
	return token.Value
}

func (i *Interpreter) IsEnd() bool {
	return i.lexer.IsEnd()
}

func (i *Interpreter) Run() {
	for {
		token := i.lexer.Next()
		fmt.Println(i.Expr(&token))

		if i.IsEnd() {
			break
		}
	}
}
