package interpret

import (
	"kylin/include"
	"kylin/lexer"
)

func (i *KyRuntime) SetException(name string, message string, line int, column int) {
	i.err = include.NewException(name, message, line, column)
}

func (i *KyRuntime) GetException() interface{} {
	return i.err
}

func (i *KyRuntime) IsException() bool {
	return i.err != nil
}

func (i *KyRuntime) ClearException() {
	i.err = nil
}

func (i *KyRuntime) IsCaching() bool {
	return i.caching
}

func (i *KyRuntime) SetCaching(caching bool) {
	i.caching = caching
}

func (i *KyRuntime) Throw(name string, message string) interface{} {
	if i.IsCaching() {
		i.SetException(name, message, i.GetCurrentLine(), i.GetCurrentColumn())
		return nil
	} else {
		return include.Raise(name, message, i.GetCurrentLine(), i.GetCurrentColumn())
	}
}

func (i *KyRuntime) ThrowError(message string) interface{} {
	return i.Throw("Error", message)
}

func (i *KyRuntime) ExceptionCall() interface{} {
	i.SetCaching(true)
	defer i.SetCaching(false)

	if i.Peek().Type != lexer.LeftBrace {
		i.Throw("SyntaxError", "Exception must have {")
	}
	i.Skip()

	for {
		if i.IsEnd() {
			i.Throw("SyntaxError", "Exception not closed")
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
		if i.IsException() {
			for {
				if i.IsEnd() {
					i.Throw("SyntaxError", "Exception not closed")
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
			break
		}
	}

	i.SetVariable("error", i.GetException())
	i.ClearException()

	if i.Peek().Type == lexer.Catch {
		i.Skip()

		if i.Peek().Type != lexer.LeftBrace {
			i.Throw("SyntaxError", "Catch must have {")
		}
		i.Skip()

		for {
			if i.IsEnd() {
				i.Throw("SyntaxError", "Catch not closed")
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
	}
	return nil
}
