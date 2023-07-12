package interpret

import (
	"fmt"
	"kylin/utils"
	"strconv"
)

type Interpreter struct {
	lexer  *Lexer
	scope  *Scope
	buffer interface{}
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
		if !i.IsEnd() && i.lexer.Peek().Type == Equals {
			i.lexer.Skip()
			i.SetVariable(token.Value, i.Expr(i.lexer.GetNextPtr()))
			return nil
		} else {
			return i.GetVariable(token.Value)
		}
	case LeftParenthesis:
		return i.Expr(i.lexer.GetNextPtr())
	case Addition:
		fmt.Println(i.GetBuffer(), "hi")
		return i.GetBuffer().(float64) + i.Expr(i.lexer.GetNextPtr()).(float64)
	case Subtraction:
		return i.GetBuffer().(float64) - i.Expr(i.lexer.GetNextPtr()).(float64)
	case Multiplication:
		return i.GetBuffer().(float64) * i.Expr(i.lexer.GetNextPtr()).(float64)
	case Division:
		return i.GetBuffer().(float64) / i.Expr(i.lexer.GetNextPtr()).(float64)
	case Modulo:
		return int(i.GetBuffer().(float64)) % int(i.Expr(i.lexer.GetNextPtr()).(float64))
	case Exponent:
		return utils.Pow(i.GetBuffer().(float64), i.Expr(i.lexer.GetNextPtr()).(float64))
	case PlusEquals:
		i.GetBuffer()
	}
	return token
}

func (i *Interpreter) GetBuffer() interface{} {
	return i.buffer
}

func (i *Interpreter) SetBuffer(token interface{}) {
	i.buffer = token
}

func ReadUntilEnter(lexer *Lexer) string {
	var result string
	for {
		token := lexer.Next()
		if token.Type == Enter || token.Type == EOF {
			break
		}
		result += token.Value
	}
	return result
}

func (i *Interpreter) IsEnd() bool {
	return i.lexer.IsEnd()
}

func (i *Interpreter) Run() {
	for {
		i.SetBuffer(i.Expr(i.lexer.GetNextPtr()))
		fmt.Println(i.GetBuffer())
		if i.IsEnd() {
			break
		}
	}
}
