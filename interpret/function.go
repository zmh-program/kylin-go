package interpret

import (
	"kylin/include"
	"kylin/module"
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
		interpreter.Run()
		return interpreter.buffer
	}
}

func (i *Interpreter) ReadFunctionName() string {
	name := ""
	for {
		if i.IsEnd() {
			panic("Function not closed")
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
		panic("Function must be followed by parenthesis")
	}
	i.Skip()
	for {
		if i.IsEnd() {
			panic("Function not closed")
		}
		if i.Peek().Type == RightParenthesis {
			i.Skip()
			break
		}
		param := i.Peek()
		if param.Type != Identifier {
			panic("Function param must be identifier")
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
	for {
		i.lexer.NextCursor()
		if i.IsEnd() {
			panic("Function body not closed")
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
