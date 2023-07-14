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
	}

	return nil
}

func (i *Interpreter) ExceptionCall() interface{} {
	i.SetCaching(true)
	defer i.SetCaching(false)

	if i.Peek().Type != LeftBrace {
		i.Throw("SyntaxError", "Exception must have {")
	}
	i.Skip()

	for {
		if i.IsEnd() {
			i.Throw("SyntaxError", "Exception not closed")
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
		if i.IsException() {
			for {
				if i.IsEnd() {
					i.Throw("SyntaxError", "Exception not closed")
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
			break
		}
	}

	if i.Peek().Type == Catch {
		i.Skip()
		i.SetVariable("error", i.GetException())

		if i.Peek().Type != LeftBrace {
			i.Throw("SyntaxError", "Catch must have {")
		}
		i.Skip()

		for {
			if i.IsEnd() {
				i.Throw("SyntaxError", "Catch not closed")
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
	}
	i.ClearException()
	return nil
}
