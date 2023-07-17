package lexer

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"kylin/i18n"
	"kylin/lib"
	"os"
	"strconv"
)

type Parser struct {
	data   []Token
	cursor int
}

func NewParser(data string, i18n *i18n.Manager) *Parser {
	lexer := NewLexer(data, i18n)
	return &Parser{data: lexer.ReadAll(), cursor: 0}
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

func (p *Parser) SkipN(n int) {
	p.cursor += n
}

func (p *Parser) Assert(t TokenType) {
	if p.Next().Type != t {
		lib.Fatal("Syntax error: ", p.Get().Value, " is not ", t)
	}
}

func (p *Parser) ExprArray() ArrayType {
	array := make([]Value, 0)
	for p.HasNext() {
		p.Skip()
		if p.Get().Type == RightBracket {
			p.Skip()
			break
		}
		array = append(array, p.Expr())
		if p.Get().Type == Comma {
			p.Skip()
		}
	}
	return array
}

func (p *Parser) ExprMap() MapType {
	m := make(MapType)
	for p.HasNext() {
		p.Skip()
		if p.Get().Type == RightBrace {
			p.Skip()
			break
		}
		key := p.Get()
		p.Assert(Colon)
		p.Skip()
		value := p.Expr()
		m[key.Value] = value
		if p.Get().Type == Comma {
			p.Skip()
		}
	}
	return m
}

func (p *Parser) Expr() Value {
	token := p.Get()
	switch token.Type {
	case Integer:
		value, _ := strconv.Atoi(token.Value)
		return Value{Type: Integer, Value: value}
	case Float:
		value, _ := strconv.ParseFloat(token.Value, 64)
		return Value{Type: Float, Value: value}
	case String:
		return Value{Type: String, Value: token.Value}
	case Identifier:
		if p.Peek().Type == LeftParenthesis {
			return Value{Type: Function, Value: p.ParseFunctionCall(token).Data}
		}
		return Value{Type: Identifier, Value: token.Value}
	case LeftBracket:
		return Value{Type: Array, Value: p.ExprArray()}
	case LeftBrace:
		return Value{Type: Map, Value: p.ExprMap()}
	case True:
		return Value{Type: True, Value: true}
	case False:
		return Value{Type: False, Value: false}
	case Null:
		return Value{Type: Null, Value: nil}
	}
	return Value{Type: token.Type, Value: token.Value}
}

func (p *Parser) ParseCount() ExecSequence {
	count := CountStruct{}
	for p.HasNext() {
		p.Skip()
		if lib.InArray(p.Get().Type, []TokenType{Sep, RightParenthesis, Comma, RightBracket, RightBrace}) {
			p.Skip()
			break
		}
		count.Value = append(count.Value, p.Expr())
	}
	return ExecSequence{Type: CountSequence, Data: count}
}

func (p *Parser) ParseFunctionCall(n Token) ExecSequence {
	call := FunctionCallStruct{}
	call.Name = n.Value
	p.Skip()
	for p.HasNext() {
		if p.Get().Type == RightParenthesis {
			p.Skip()
			break
		}
		call.Args = append(call.Args, p.ParseCount())
		if p.Peek().Type == EOF {
			return ExecSequence{Type: FunctionCallSequence, Data: call}
		}
		if p.Get().Type == Comma {
			p.Skip()
		}
	}
	return ExecSequence{Type: FunctionCallSequence, Data: call}
}

func (p *Parser) ParseIdentifier(n Token) ExecSequence {
	if p.Peek().Type == LeftParenthesis {
		return p.ParseFunctionCall(n)
	} else if lib.InArray(p.Peek().Type, AssignType) {
		p.Skip()
		assign := AssignStruct{
			Variable: n.Value,
			Type:     p.Get().Type,
			Value:    p.ParseCount(),
		}
		return ExecSequence{Type: AssignSequence, Data: assign}
	}
	return ExecSequence{Type: IdentifierSequence, Data: CountStruct{
		Value: []Value{p.Expr()},
	}}
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
	for p.HasNext() && p.Get().Type == Sep {
		p.Skip()
	}
	return
}

func (p *Parser) GetBufferStack() []Token {
	buffer := []Token{p.Get()}
	cursor := p.cursor
	for p.HasNext() {
		if p.Peek().Type == Sep || p.Peek().Type == EOF {
			break
		}
		p.Skip()
		buffer = append(buffer, p.Peek())
	}
	p.cursor = cursor
	return buffer
}

func (p *Parser) Parse() ExecSequence {
	p.BreakSep()
	n := p.Get()
	switch n.Type {
	case Identifier:
		return p.ParseIdentifier(n)
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
	return p.ParseCount()
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
	return stack
}

func (p *Parser) Compile() []byte {
	return CompileStack(p.ParseAll())
}

func (p *Parser) FSWrite(path string) {
	err := os.WriteFile(path, p.Compile(), 0644)
	if err != nil {
		panic(err)
	}
}

func CompileStack(stack []ExecSequence) []byte {
	text, err := json.Marshal(stack)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	if _, err := writer.Write(text); err != nil {
		panic(err)
	}
	if err := writer.Close(); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
