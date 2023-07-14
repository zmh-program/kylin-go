package interpret

import (
	"kylin/utils"
	"log"
	"strconv"
)

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
	case And:
		i.Skip()
		return utils.ToBool(value) && utils.ToBool(i.ExprNext())
	case Or:
		i.Skip()
		return utils.ToBool(value) || utils.ToBool(i.ExprNext())
	case IsEquals:
		i.Skip()
		return value == i.ExprNext()
	case NotEquals:
		i.Skip()
		return value != i.ExprNext()
	case GreaterThan:
		i.Skip()
		return value.(float64) > i.ExprNext().(float64)
	case LessThan:
		i.Skip()
		return value.(float64) < i.ExprNext().(float64)
	case GreaterThanOrEqual:
		i.Skip()
		return value.(float64) >= i.ExprNext().(float64)
	case LessThanOrEqual:
		i.Skip()
		return value.(float64) <= i.ExprNext().(float64)
	default:
		return value
	}
}

func (i *Interpreter) ParenthesisCall() []interface{} {
	i.Skip()

	buffer := make([]interface{}, 0)
	for {
		if i.IsEnd() {
			log.Fatal("Unexpected end of file")
		}
		if i.Peek().Type == RightParenthesis {
			i.Skip()
			break
		}

		buffer = append(buffer, i.ExprNext())
	}
	return buffer
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
