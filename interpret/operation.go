package interpret

import (
	"kylin/utils"
	"log"
)

func (i *Interpreter) ConditionCall() interface{} {
	condition := utils.ToBool(i.ExprNext())
	if condition {
		i.Skip()
		for {
			if i.IsEnd() {
				log.Fatalln("Condition not closed")
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
	} else {
		i.Skip()
		for {
			if i.IsEnd() {
				log.Fatalln("Condition not closed")
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
					log.Fatalln("Condition not closed")
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
				log.Fatalln("While not closed")
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

		i.lexer.cursor = cursor
		i.WhileCall()
	} else {
		i.Skip()
		for {
			if i.IsEnd() {
				log.Fatalln("While not closed")
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
