package interpret

import (
	"kylin/utils"
)

func (i *Interpreter) ConditionCall() interface{} {
	condition := utils.ToBool(i.ExprNext())
	if condition {
		i.Skip()
		for {
			if i.IsEnd() {
				i.Throw("SyntaxError", "Condition not closed")
			}
			if i.Peek().Type == RightBrace {
				i.Skip()
				break
			}
			if n := i.Peek(); n.Type == Continue || n.Type == Break {
				return n
			}
			i.ExprNext()
			if i.Peek().Type == Comma {
				i.Skip()
			}
		}
	} else {
		i.Skip()
		for {
			if i.IsEnd() {
				i.Throw("SyntaxError", "Condition not closed")
			}

			if i.Peek().Type == RightBrace {
				i.Skip()
				break
			}
			i.Skip()
		}
		if i.Peek().Type == Else {
			i.Skip()
			i.Skip()
			for {
				if i.IsEnd() {
					i.Throw("SyntaxError", "Condition not closed")
				}
				if i.Peek().Type == RightBrace {
					i.Skip()
					break
				}

				i.ExprNext()
				if i.Peek().Type == Comma {
					i.Skip()
				}
			}
		} else if i.Peek().Type == Elif {
			i.Skip()
			i.ConditionCall()
		}
	}
	return nil
}

func (i *Interpreter) WhileCall() interface{} {
	cursor := i.lexer.cursor
	line := i.lexer.line
	condition := utils.ToBool(i.ExprNext())

	if condition {
		i.Skip()
		for {
			if i.IsEnd() {
				i.Throw("SyntaxError", "While not closed")
			}
			if i.Peek().Type == RightBrace {
				i.Skip()
				break
			}

			if i.Peek().Type == Break {
				return nil
			} else if i.Peek().Type == Continue {
				i.Skip()
				break
			}

			i.ExprNext()

			if i.Peek().Type == Comma {
				i.Skip()
			}
		}

		i.lexer.cursor = cursor
		i.lexer.line = line
		i.WhileCall()
	} else {
		i.Skip()
		for {
			if i.IsEnd() {
				i.Throw("SyntaxError", "While not closed")
			}

			if i.Peek().Type == RightBrace {
				i.Skip()
				break
			}
			i.Skip()
		}
	}
	return nil
}

func (i *Interpreter) ForCall() interface{} {
	if i.Peek().Type != Identifier {
		i.Throw("SyntaxError", "For must have a variable")
	}
	param := i.Next().Value

	if i.Peek().Type != In {
		i.Throw("SyntaxError", "For must have in keyword")
	}
	i.Skip()

	array := i.ExprNext().([]interface{})

	if i.Peek().Type != LeftBrace {
		i.Throw("SyntaxError", "For must have {")
	}

	i.Skip()

	idx := 0
	cursor := i.lexer.cursor
	line := i.lexer.line

	for {
		i.SetVariable(param, array[idx])
		if idx++; idx >= len(array) {
			break
		}

		for {
			if i.IsEnd() {
				i.Throw("SyntaxError", "For not closed")
			}
			if i.Peek().Type == RightBrace {
				i.Skip()
				break
			}
			if i.Peek().Type == Break {
				return nil
			} else if i.Peek().Type == Continue {
				i.Skip()
				break
			}

			i.ExprNext()
			if i.Peek().Type == Comma {
				i.Skip()
			}
		}

		i.lexer.cursor = cursor
		i.lexer.line = line
	}

	return nil
}
