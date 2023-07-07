package interpret

import "kylin/utils"

type Lexer struct {
	data   string
	cursor int
}

func NewLexer(data string) *Lexer {
	return &Lexer{data: data, cursor: 0}
}

func (l *Lexer) Next() Token {
	if l.cursor >= len(l.data) {
		return Token{Type: EOF, Value: ""}
	}
	value := l.data[l.cursor]
	l.cursor++
	switch value {
	case '+':
		return Token{Type: Addition, Value: "+"}
	case '-':
		return Token{Type: Subtraction, Value: "-"}
	case '*':
		return Token{Type: Multiplication, Value: "*"}
	case '/':
		return Token{Type: Division, Value: "/"}
	case '(':
		return Token{Type: LeftParenthesis, Value: "("}
	case ')':
		return Token{Type: RightParenthesis, Value: ")"}
	case ' ':
		return Token{Type: Space, Value: " "}
	case '\n':
		return Token{Type: Enter, Value: "\n"}
	case '\t':
		return Token{Type: Space, Value: "\t"}
	default:
		if utils.IsDigit(value) {
			return Token{Type: Number, Value: l.readNumber()}
		}
		if utils.IsLetter(value) {
			return Token{Type: Identifier, Value: l.readIdentifier()}
		}
	}
}

func (l *Lexer) readNumber() string {
	var number string
	for l.cursor < len(l.data) {
		value := l.data[l.cursor]
		if value >= '0' && value <= '9' {
			number += string(value)
			l.cursor++
		} else {
			break
		}
	}
	return number
}

func (l *Lexer) readIdentifier() string {
	var identifier string
	var size int
	for l.cursor < len(l.data) {
		size++
		value := l.data[l.cursor]
		if size == 1 && !utils.IsLetter(value) {
			break
		}
		if utils.IsRegular(value) {
			identifier += string(value)
			l.cursor++
		} else {
			break
		}
	}
	return identifier
}
