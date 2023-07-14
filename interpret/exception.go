package interpret

import (
	"kylin/include"
)

func (i *Interpreter) SetException(name string, message string, line int, column int) {
	i.err = include.NewException(name, message, line, column)
}

func (i *Interpreter) GetException() interface{} {
	return i.err
}

func (i *Interpreter) IsException() bool {
	return i.err != nil
}

func (i *Interpreter) ClearException() {
	i.err = nil
}

func (i *Interpreter) IsCaching() bool {
	return i.caching
}

func (i *Interpreter) SetCaching(caching bool) {
	i.caching = caching
}

func (i *Interpreter) Throw(name string, message string) interface{} {
	if i.IsCaching() {
		i.SetException(name, message, i.GetCurrentLine(), i.GetCurrentColumn())
		return nil
	} else {
		return include.Raise(name, message, i.GetCurrentLine(), i.GetCurrentColumn())
	}
}

func (i *Interpreter) ThrowError(message string) interface{} {
	return i.Throw("Error", message)
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

	i.SetVariable("error", i.GetException())
	i.ClearException()

	if i.Peek().Type == Catch {
		i.Skip()

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
	return nil
}
