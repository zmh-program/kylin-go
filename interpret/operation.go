package interpret

import "kylin/lexer"

func (i *KyRuntime) ConditionCall() interface{} {
	condition := ToBool(i.ExprNext())
	if condition {
		i.Skip()
		for {
			if i.IsEnd() {
				i.Throw("SyntaxError", "Condition not closed")
			}
			if i.Peek().Type == lexer.RightBrace {
				i.Skip()
				break
			}
			if n := i.Peek(); n.Type == lexer.Continue || n.Type == lexer.Break {
				return n
			}
			i.ExprNext()
			if i.Peek().Type == lexer.Comma {
				i.Skip()
			}
		}
	} else {
		i.Skip()
		for {
			if i.IsEnd() {
				i.Throw("SyntaxError", "Condition not closed")
			}

			if i.Peek().Type == lexer.RightBrace {
				i.Skip()
				break
			}
			i.Skip()
		}
		if i.Peek().Type == lexer.Else {
			i.Skip()
			i.Skip()
			for {
				if i.IsEnd() {
					i.Throw("SyntaxError", "Condition not closed")
				}
				if i.Peek().Type == lexer.RightBrace {
					i.Skip()
					break
				}

				i.ExprNext()
				if i.Peek().Type == lexer.Comma {
					i.Skip()
				}
			}
		} else if i.Peek().Type == lexer.Elif {
			i.Skip()
			i.ConditionCall()
		}
	}
	return nil
}

func (i *KyRuntime) WhileCall() interface{} {
	cursor := i.lexer.Cursor
	line := i.lexer.Line
	condition := ToBool(i.ExprNext())

	if condition {
		i.Skip()
		for {
			if i.IsEnd() {
				i.Throw("SyntaxError", "While not closed")
			}
			if i.Peek().Type == lexer.RightBrace {
				i.Skip()
				break
			}

			if i.Peek().Type == lexer.Break {
				return nil
			} else if i.Peek().Type == lexer.Continue {
				i.Skip()
				break
			}

			i.ExprNext()

			if i.Peek().Type == lexer.Comma {
				i.Skip()
			}
		}

		i.lexer.Cursor = cursor
		i.lexer.Line = line
		i.WhileCall()
	} else {
		i.Skip()
		for {
			if i.IsEnd() {
				i.Throw("SyntaxError", "While not closed")
			}

			if i.Peek().Type == lexer.RightBrace {
				i.Skip()
				break
			}
			i.Skip()
		}
	}
	return nil
}

func (i *KyRuntime) ForCall() interface{} {
	if i.Peek().Type != lexer.Identifier {
		i.Throw("SyntaxError", "For must have a variable")
	}
	param := i.Next().Value

	if i.Peek().Type != lexer.In {
		i.Throw("SyntaxError", "For must have in keyword")
	}
	i.Skip()

	array := i.ExprNext().([]interface{})

	if i.Peek().Type != lexer.LeftBrace {
		i.Throw("SyntaxError", "For must have {")
	}

	i.Skip()

	idx := 0
	cursor := i.lexer.Cursor
	line := i.lexer.Line

	for {
		i.SetVariable(param, array[idx])
		if idx++; idx >= len(array) {
			break
		}

		for {
			if i.IsEnd() {
				i.Throw("SyntaxError", "For not closed")
			}
			if i.Peek().Type == lexer.RightBrace {
				i.Skip()
				break
			}
			if i.Peek().Type == lexer.Break {
				return nil
			} else if i.Peek().Type == lexer.Continue {
				i.Skip()
				break
			}

			i.ExprNext()
			if i.Peek().Type == lexer.Comma {
				i.Skip()
			}
		}

		i.lexer.Cursor = cursor
		i.lexer.Line = line
	}

	return nil
}
