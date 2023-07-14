package interpret

import (
	"fmt"
	"kylin/utils"
)

func (i *Interpreter) ConditionCall() interface{} {
	condition := utils.ToBool(i.ExprNext())
	if condition {
		i.Skip()
		for {
			if i.IsEnd() {
				panic("Condition not closed")
			}
			if i.Peek().Type == RightParenthesis {
				i.Skip()
				break
			}
			fmt.Println(i.Peek())
			i.ExprNext()
			if i.Peek().Type == Comma {
				i.Skip()
			}
		}
	} else {
		i.Skip()
		for {
			if i.IsEnd() {
				panic("Condition not closed")
			}
			if i.Peek().Type == RightParenthesis {
				i.Skip()
				break
			}
		}
		if i.Peek().Type == Else {
			i.Skip()
			for {
				if i.IsEnd() {
					panic("Condition not closed")
				}
				if i.Peek().Type == RightParenthesis {
					i.Skip()
					break
				}

				i.ExprNext()
				if i.Peek().Type == Comma {
					i.Skip()
				}
			}
		} else if i.Peek().Type == Elif {
			i.ConditionCall()
		}
	}
	return nil
}
