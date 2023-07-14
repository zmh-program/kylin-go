package interpret

import (
	"kylin/include"
	"kylin/module"
	"kylin/utils"
	"log"
	"strconv"
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

func (i *Interpreter) FunctionCall(token *Token) (bool, interface{}) {
	if token.Type != Identifier {
		return false, nil
	}

	if i.IsEnd() {
		return false, nil
	}

	if i.Peek().Type != LeftParenthesis {
		return false, nil
	}
	i.Skip()

	param := make([]interface{}, 0)
	for {
		if i.IsEnd() {
			return false, nil
		}
		if i.Peek().Type == RightParenthesis {
			i.Skip()
			break
		}
		param = append(param, i.ExprNext())
		if i.Peek().Type == Comma {
			i.Skip()
		}
	}

	resp := utils.CallFunc(i.GetVariable(token.Value), param)
	return true, i.CountCall(resp)
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

func (i *Interpreter) CountCall(token interface{}) interface{} {
	var value interface{}
	switch (token).(type) {
	case *Token:
		token := token.(*Token)
		switch token.Type {
		case Identifier:
			value = i.GetVariable(token.Value)
		case Integer:
			value = utils.MustGet(strconv.ParseInt(token.Value, 10, 64))
		case Float:
			value = utils.MustGet(strconv.ParseFloat(token.Value, 64))
		default:
			value = token.Value
		}
	default:
		value = token
	}

	peek := i.Peek()
	switch peek.Type {
	case Addition:
		i.Skip()
		return value.(float64) + i.ExprNext().(float64)
	case Subtraction:
		i.Skip()
		return value.(float64) - i.ExprNext().(float64)
	case Multiplication:
		i.Skip()
		return value.(float64) * i.ExprNext().(float64)
	case Division:
		i.Skip()
		return value.(float64) / i.ExprNext().(float64)
	case Modulo:
		i.Skip()
		return float64(int64(value.(float64)) % int64(i.ExprNext().(float64)))
	case Exponent:
		i.Skip()
		return utils.Pow(value.(float64), i.ExprNext().(float64))
	default:
		return value
	}
}

func (i *Interpreter) ParenthesisCall() interface{} {
	i.Skip()
	r := i.ExprNext()
	if i.Peek().Type != RightParenthesis {
		panic("Parenthesis not closed")
	}
	i.Skip()
	return r
}

func (i *Interpreter) ReadArray() []interface{} {
	array := make([]interface{}, 0)
	for {
		if i.IsEnd() {
			panic("Array not closed")
		}
		if i.Peek().Type == RightBracket {
			i.Skip()
			break
		}
		array = append(array, i.ExprNext())
		if i.Peek().Type == Comma {
			i.Skip()
		}
	}
	return array
}

func (i *Interpreter) ReadObject() map[string]interface{} {
	object := make(map[string]interface{})
	for {
		if i.IsEnd() {
			log.Fatalln("Object not closed")
		}
		if i.Peek().Type == RightBrace {
			i.Skip()
			break
		}
		key := i.Peek()
		if key.Type != String {
			log.Fatalln("Object key must be string")
		}
		i.Skip()
		if i.Peek().Type != Colon {
			log.Fatalln("Object key must be followed by colon")
		}
		i.Skip()
		object[key.Value] = i.ExprNext()
		if i.Peek().Type == Comma {
			i.Skip()
		}
	}
	return object
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
