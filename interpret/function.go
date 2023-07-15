package interpret

import (
	"kylin/i18n"
	"kylin/include"
	"kylin/lexer"
	"kylin/module"
	"reflect"
)

type KyFunction struct {
	Name   string
	Params []string
	Body   string
}

func (k *KyFunction) CallWrapper(scope *include.Scope, i18n *i18n.Manager) interface{} {
	return func(args ...interface{}) interface{} {
		runtime := &KyRuntime{
			lexer:  lexer.NewLexer(k.Body, i18n),
			scope:  include.NewScope(scope),
			module: module.NewManager(),
			i18n:   i18n,
		}

		for i, param := range k.Params {
			runtime.SetVariable(param, args[i])
		}
		return runtime.Run()
	}
}

func (i *KyRuntime) ReadFunctionName() string {
	name := ""
	for {
		if i.IsEnd() {
			i.Throw("SyntaxError", "Function name not closed")
		}
		if i.Peek().Type == lexer.LeftParenthesis {
			break
		}
		name += i.Peek().Value
		i.Skip()
	}
	return name
}

func (i *KyRuntime) ReadFunctionParams() []string {
	params := make([]string, 0)
	if i.Peek().Type != lexer.LeftParenthesis {
		i.Throw("SyntaxError", "Function params must start with '('")
	}
	i.Skip()
	for {
		if i.IsEnd() {
			i.Throw("SyntaxError", "Function params not closed")
		}
		if i.Peek().Type == lexer.RightParenthesis {
			i.Skip()
			break
		}
		param := i.Peek()
		if param.Type != lexer.Identifier {
			i.Throw("SyntaxError", "Function params must be identifier")
		}
		i.Skip()
		params = append(params, param.Value)
		if i.Peek().Type == lexer.Comma {
			i.Skip()
		}
	}
	return params
}

func (i *KyRuntime) ReadFunctionBody() string {
	body := ""
	i.Skip()
	for {
		i.lexer.NextCursor()
		if i.IsEnd() {
			i.Throw("SyntaxError", "Function body not closed")
		}

		if c := i.lexer.GetRune(); c == '}' {
			i.lexer.NextCursor()
			break
		} else {
			body += string(c)
		}
	}
	return body
}

func (i *KyRuntime) ReadFunction() *KyFunction {
	function := &KyFunction{
		Name:   i.ReadFunctionName(),
		Params: i.ReadFunctionParams(),
		Body:   i.ReadFunctionBody(),
	}
	i.SetVariable(
		function.Name,
		function.CallWrapper(i.GetScope(), i.GetI18n()),
	)
	return function
}

func (i *KyRuntime) FunctionCall(token *lexer.Token) (bool, interface{}) {
	if token.Type != lexer.Identifier {
		return false, nil
	}

	if i.IsEnd() {
		return false, nil
	}

	if i.Peek().Type != lexer.LeftParenthesis {
		return false, nil
	}
	i.Skip()

	param := make([]interface{}, 0)
	for {
		if i.IsEnd() {
			return false, nil
		}
		if i.Peek().Type == lexer.RightParenthesis {
			i.Skip()
			break
		}
		param = append(param, i.ExprNext())
		if i.Peek().Type == lexer.Comma {
			i.Skip()
		}
	}

	if i.IsException() {
		return false, nil
	}
	resp := CallFunc(i.GetVariable(token.Value), param)
	return true, i.CountCall(resp)
}

func CallFunc(_fn interface{}, _args []interface{}) interface{} {
	fn := reflect.ValueOf(_fn)

	args := make([]reflect.Value, len(_args))
	for i, arg := range _args {
		if arg == nil {
			args[i] = reflect.ValueOf(new(interface{})).Elem()
			continue
		}
		args[i] = reflect.ValueOf(arg)
	}

	resp := fn.Call(args)
	if len(resp) == 0 {
		return nil
	} else if len(resp) == 1 {
		return resp[0].Interface()
	} else {
		results := make([]interface{}, len(resp))
		for i, r := range resp {
			results[i] = r.Interface()
		}
		return results
	}
}
