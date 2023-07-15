package interpret

import (
	"log"
	"math"
	"strconv"
)

func ToBool(data interface{}) bool {
	if data == nil {
		return false
	}
	switch data.(type) {
	case bool:
		return data.(bool)
	case int:
		return data.(int) != 0
	case float64:
		return data.(float64) != 0
	case string:
		return data.(string) != ""
	case []interface{}:
		return len(data.([]interface{})) != 0
	case map[string]interface{}:
		return len(data.(map[string]interface{})) != 0
	default:
		return true
	}
}

func (i *KyRuntime) ReadArray() []interface{} {
	array := make([]interface{}, 0)
	for {
		if i.IsEnd() {
			i.Throw("SyntaxError", "Array not closed")
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

func (i *KyRuntime) ReadObject() map[string]interface{} {
	object := make(map[string]interface{})
	for {
		if i.IsEnd() {
			i.Throw("SyntaxError", "Object not closed")
		}
		if i.Peek().Type == RightBrace {
			i.Skip()
			break
		}
		key := i.Peek()
		if key.Type != String {
			i.Throw("SyntaxError", "Object key must be string")
		}
		i.Skip()
		if i.Peek().Type != Colon {
			i.Throw("SyntaxError", "Object key must be string")
		}
		i.Skip()
		object[key.Value] = i.ExprNext()
		if i.Peek().Type == Comma {
			i.Skip()
		}
	}
	return object
}

func (i *KyRuntime) MustGet(data interface{}, err error) interface{} {
	if err != nil {
		i.Throw("TypeError", err.Error())
	}
	return data
}

func (i *KyRuntime) CountCall(token interface{}) interface{} {
	var value interface{}
	switch (token).(type) {
	case *Token:
		token := token.(*Token)
		switch token.Type {
		case Identifier:
			value = i.GetVariable(token.Value)
		case Integer:
			value = i.MustGet(strconv.ParseInt(token.Value, 10, 64))
		case Float:
			value = i.MustGet(strconv.ParseFloat(token.Value, 64))
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
		return math.Pow(value.(float64), i.ExprNext().(float64))
	case And:
		i.Skip()
		return ToBool(value) && ToBool(i.ExprNext())
	case Or:
		i.Skip()
		return ToBool(value) || ToBool(i.ExprNext())
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

func (i *KyRuntime) ParenthesisCall() []interface{} {
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

func (i *KyRuntime) AssignCall(token *Token) bool {
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
		i.SetVariable(token.Value, math.Pow(i.GetVariable(token.Value).(float64), i.ExprNext().(float64)))
		return true
	default:
		return false
	}
}
