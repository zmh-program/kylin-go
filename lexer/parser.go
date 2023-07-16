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
	return &Parser{data: lexer.ReadAll()}
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
	for p.HasNext() {
		if p.Peek().Type == LeftBrace {
			break
		}
		_for.Condition = append(_for.Condition, p.Next())
	}
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

func (p *Parser) Parse() ExecSequence {
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
	}
	return ExecSequence{}
}

func (p *Parser) ParseAll() []ExecSequence {
	stack := make([]ExecSequence, 0)
	for p.HasNext() {
		if p.Next().Type == Sep {
			continue
		}
		stack = append(stack, p.Parse())
	}
	fmt.Println(stack)
	return stack
}
