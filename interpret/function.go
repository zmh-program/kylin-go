package interpret

import (
	"kylin/include"
	"kylin/module"
	"kylin/utils"
)

type KyFunction struct {
	Name   string
	Params []string
	Body   string
}

func (k *KyFunction) CallWrapper(scope *include.Scope) interface{} {
	return func(args ...interface{}) interface{} {
		interpreter := &Interpreter{
			lexer:  NewLexer(k.Body),
			scope:  include.NewScope(scope),
			module: module.NewManager(),
		}

		for i, param := range k.Params {
			interpreter.SetVariable(param, args[i])
		}
		return interpreter.Run()
	}
}

func (i *Interpreter) ReadFunctionName() string {
	name := ""
	for {
		if i.IsEnd() {
			i.Throw("SyntaxError", "Function name not closed")
		}
		if i.Peek().Type == LeftParenthesis {
			break
		}
		name += i.Peek().Value
		i.Skip()
	}
	return name
}

func (i *Interpreter) ReadFunctionParams() []string {
	params := make([]string, 0)
	if i.Peek().Type != LeftParenthesis {
		i.Throw("SyntaxError", "Function params must start with '('")
	}
	i.Skip()
	for {
		if i.IsEnd() {
			i.Throw("SyntaxError", "Function params not closed")
		}
		if i.Peek().Type == RightParenthesis {
			i.Skip()
			break
		}
		param := i.Peek()
		if param.Type != Identifier {
			i.Throw("SyntaxError", "Function params must be identifier")
		}
		i.Skip()
		params = append(params, param.Value)
		if i.Peek().Type == Comma {
			i.Skip()
		}
	}
	return params
}

func (i *Interpreter) ReadFunctionBody() string {
	body := ""
	i.Skip()
	for {
		i.lexer.NextCursor()
		if i.IsEnd() {
			i.Throw("SyntaxError", "Function body not closed")
		}

		if c := i.lexer.GetByte(); c == '}' {
			i.lexer.NextCursor()
			break
		} else {
			body += string(c)
		}
	}
	return body
}

func (i *Interpreter) ReadFunction() *KyFunction {
	function := &KyFunction{
		Name:   i.ReadFunctionName(),
		Params: i.ReadFunctionParams(),
		Body:   i.ReadFunctionBody(),
	}
	i.SetVariable(
		function.Name,
		function.CallWrapper(i.GetScope()),
	)
	return function
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

	if i.IsException() {
		return false, nil
	}
	resp := utils.CallFunc(i.GetVariable(token.Value), param)
	return true, i.CountCall(resp)
}
