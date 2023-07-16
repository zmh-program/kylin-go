package lexer

import (
	"fmt"
	"kylin/i18n"
	"kylin/lib"
)

type Parser struct {
	data   []Token
	cursor int
}

func NewParser(data string, i18n *i18n.Manager) *Parser {
	lexer := NewLexer(data, i18n)
	return &Parser{data: lexer.ReadAll(), cursor: -1}
}

func (p *Parser) HasNext() bool {
	return p.cursor+1 < len(p.data)
}

func (p *Parser) Get() Token {
	return p.data[p.cursor]
}

func (p *Parser) Next() Token {
	if p.HasNext() {
		p.cursor++
		return p.Get()
	}
	return Token{}
}

func (p *Parser) Peek() Token {
	if p.HasNext() {
		return p.data[p.cursor+1]
	}
	return Token{}
}

func (p *Parser) Skip() {
	p.cursor++
}

func (p *Parser) Assert(t TokenType) {
	if p.Next().Type != t {
		lib.Fatal("Syntax error: ", p.Get().Value, " is not ", t)
	}
}

func (p *Parser) ParseCount() ExecSequence {
	count := CountStruct{}
	for p.HasNext() {
		if p.Peek().Type == Sep {
			p.Skip()
			break
		}
		count.Value = append(count.Value, p.Next())
	}
	return ExecSequence{Type: CountSequence, Data: count}
}

func (p *Parser) ParseAssign() ExecSequence {
	assign := AssignStruct{
		Variable: p.Get().Value,
		Type:     p.Next().Type,
		Value:    p.ParseCount(),
	}
	return ExecSequence{Type: AssignSequence, Data: assign}
}

func (p *Parser) ParseFunction() ExecSequence {
	function := FunctionStruct{}
	p.Assert(LeftParenthesis)
	for p.HasNext() {
		if p.Peek().Type == RightParenthesis {
			p.Skip()
			break
		}
		function.Args = append(function.Args, p.Next().Value)
		if p.Peek().Type == Comma {
			p.Skip()
		}
	}
	p.Assert(LeftBrace)
	for p.HasNext() {
		if p.Peek().Type == RightBrace {
			p.Skip()
			break
		}
		function.Body = append(function.Body, p.Parse())
		if p.Peek().Type == Sep {
			p.Skip()
		}
	}
	return ExecSequence{Type: FunctionSequence, Data: function}
}

func (p *Parser) ParseWhile() ExecSequence {
	while := WhileStruct{}
	for p.HasNext() {
		if p.Peek().Type == LeftBrace {
			break
		}
		while.Condition = append(while.Condition, p.Next())
	}
	p.Assert(LeftBrace)
	for p.HasNext() {
		if p.Peek().Type == RightBrace {
			p.Skip()
			break
		}
		while.Body = append(while.Body, p.Parse())
		if p.Peek().Type == Sep {
			p.Skip()
		}
	}
	return ExecSequence{Type: WhileSequence, Data: while}
}

func (p *Parser) ParseFor() ExecSequence {
	_for := ForStruct{}
	_for.Variable = p.Next().Value
	p.Assert(In)
	_for.Array = p.ParseCount()
	p.Assert(LeftBrace)
	for p.HasNext() {
		if p.Peek().Type == RightBrace {
			p.Skip()
			break
		}
		_for.Body = append(_for.Body, p.Parse())
		if p.Peek().Type == Sep {
			p.Skip()
		}
	}
	return ExecSequence{Type: ForSequence, Data: _for}
}

func (p *Parser) ParseTry() ExecSequence {
	_try := TryStruct{}
	p.Assert(LeftBrace)
	for p.HasNext() {
		if p.Peek().Type == RightBrace {
			p.Skip()
			break
		}
		_try.Body = append(_try.Body, p.Parse())
		if p.Peek().Type == Sep {
			p.Skip()
		}
	}

	if p.Peek().Type == Catch {
		p.Skip()
		for p.HasNext() {
			if p.Peek().Type == LeftBrace {
				break
			}
			_try.Catch = append(_try.Catch, p.Parse())
			if p.Peek().Type == Sep {
				p.Skip()
			}
		}
	}
	return ExecSequence{Type: TrySequence, Data: _try}
}

func (p *Parser) ParseElif() ElifStruct {
	_elif := ElifStruct{}
	for p.HasNext() {
		if p.Peek().Type == LeftBrace {
			break
		}
		_elif.Condition = append(_elif.Condition, p.Next())
		if p.Peek().Type == Sep {
			p.Skip()
		}
	}
	p.Assert(LeftBrace)
	for p.HasNext() {
		if p.Peek().Type == RightBrace {
			p.Skip()
			break
		}
		_elif.Body = append(_elif.Body, p.Parse())
		if p.Peek().Type == Sep {
			p.Skip()
		}
	}
	return _elif
}

func (p *Parser) ParseIf() ExecSequence {
	_if := IfStruct{}
	for p.HasNext() {
		if p.Peek().Type == LeftBrace {
			break
		}
		_if.Condition = append(_if.Condition, p.Next())
	}
	p.Assert(LeftBrace)
	for p.HasNext() {
		if p.Peek().Type == RightBrace {
			p.Skip()
			break
		}
		_if.Body = append(_if.Body, p.Parse())
		if p.Peek().Type == Sep {
			p.Skip()
		}
	}

	for p.Peek().Type == Elif {
		p.Skip()
		_if.Elif = append(_if.Elif, p.ParseElif())
	}

	if p.Peek().Type == Else {
		p.Skip()
		for p.HasNext() {
			if p.Peek().Type == LeftBrace {
				break
			}
			_if.Else = append(_if.Else, p.Parse())
			if p.Peek().Type == Sep {
				p.Skip()
			}
		}
	}
	return ExecSequence{Type: IfSequence, Data: _if}
}

func (p *Parser) BreakSep() {
	for p.HasNext() && p.Peek().Type == Sep {
		p.Skip()
	}
	return
}

func (p *Parser) Parse() ExecSequence {
	p.BreakSep()
	n := p.Next()
	switch n.Type {
	case Identifier:
		if p.Peek().Type == Sep {
			return ExecSequence{Type: IdentifierSequence, Data: CountStruct{
				Value: []Token{n},
			}}
		}
		return p.ParseAssign()
	case Function:
		return p.ParseFunction()
	case While:
		return p.ParseWhile()
	case For:
		return p.ParseFor()
	case Try:
		return p.ParseTry()
	case If:
		return p.ParseIf()
	}
	return ExecSequence{}
}

func (p *Parser) ParseAll() []ExecSequence {
	stack := make([]ExecSequence, 0)
	for p.HasNext() {
		if p.Peek().Type == Sep {
			p.Skip()
			continue
		}
		if res := p.Parse(); res.Data != nil {
			stack = append(stack, res)
		}
	}
	for _, v := range stack {
		fmt.Println(v)
	}
	fmt.Println(len(stack))
	return stack
}
