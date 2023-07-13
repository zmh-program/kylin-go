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

func (i *Interpreter) GetScope() *Scope {
	return i.scope
}

func (i *Interpreter) SetScope(scope *Scope) {
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

func (i *Interpreter) AssignCall(token *Token) bool {
	if token.Type != Identifier {
		return false
	}
	if i.IsEnd() {
		return false
	}

	peek := i.Peek()
	switch peek.Type {
	case Equals:
		i.Skip()
		r := i.ExprNext()
		fmt.Println(r)
		i.SetVariable(token.Value, r)
		return true
	case PlusEquals:
		i.Skip()
		i.SetVariable(token.Value, i.GetVariable(token.Value).(float64)+i.ExprNext().(float64))
		return true
	case MinusEquals:
		i.Skip()
		i.SetVariable(token.Value, i.GetVariable(token.Value).(float64)-i.ExprNext().(float64))
		return true
	case TimesEquals:
		i.Skip()
		i.SetVariable(token.Value, i.GetVariable(token.Value).(float64)*i.ExprNext().(float64))
		return true
	case DividedEquals:
		i.Skip()
		i.SetVariable(token.Value, i.GetVariable(token.Value).(float64)/i.ExprNext().(float64))
		return true
	case ModuloEquals:
		i.Skip()
		i.SetVariable(token.Value, float64(int64(i.GetVariable(token.Value).(float64))%int64(i.ExprNext().(float64))))
		return true
	case ExponentEquals:
		i.Skip()
		i.SetVariable(token.Value, utils.Pow(i.GetVariable(token.Value).(float64), i.ExprNext().(float64)))
		return true
	default:
		return false
	}
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
		if !i.AssignCall(token) {
			return i.GetVariable(token.Value)
		}
		return nil
	case LeftParenthesis:
		return i.ExprNext()
	case RightParenthesis:
		return i.ExprNext()
	case Addition:
		//fmt.Println(i.GetBuffer().(*Token).Value)
		return i.GetBuffer().(float64) + i.ExprNext().(float64)
	case Subtraction:
		return i.GetBuffer().(float64) - i.ExprNext().(float64)
	case Multiplication:
		return i.GetBuffer().(float64) * i.ExprNext().(float64)
	case Division:
		return i.GetBuffer().(float64) / i.ExprNext().(float64)
	case Modulo:
		return int(i.GetBuffer().(float64)) % int(i.ExprNext().(float64))
	case Exponent:
		return utils.Pow(i.GetBuffer().(float64), i.ExprNext().(float64))
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

func (i *Interpreter) ExprNext() interface{} {
	res := i.Expr(i.GetNextPtr())
	i.SetBuffer(res)
	return res
}

func (i *Interpreter) Run() {
	for !i.IsEnd() {
		i.ExprNext()
	}
}
