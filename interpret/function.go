package interpret

type KyFunction struct {
	Name   string
	Params []string
	Body   string
}

func NewKyFunction(name string, params []string) *KyFunction {
	return &KyFunction{
		Name:   name,
		Params: params,
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
	return function
}
