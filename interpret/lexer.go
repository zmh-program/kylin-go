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
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.cursor++
			return Token{Type: PlusEquals, Value: "+="}
		}
		return Token{Type: Addition, Value: "+"}
	case '-':
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.cursor++
			return Token{Type: MinusEquals, Value: "-="}
		}
		return Token{Type: Subtraction, Value: "-"}
	case '*':
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.cursor++
			return Token{Type: TimesEquals, Value: "*="}
		}
		return Token{Type: Multiplication, Value: "*"}
	case '/':
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.cursor++
			return Token{Type: DividedEquals, Value: "/="}
		}
		return Token{Type: Division, Value: "/"}
	case '%':
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.cursor++
			return Token{Type: ModuloEquals, Value: "%="}
		}
		return Token{Type: Modulo, Value: "%"}
	case '^':
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.cursor++
			return Token{Type: ExponentEquals, Value: "^="}
		}
		return Token{Type: Exponent, Value: "^"}
	case '(':
		return Token{Type: LeftParenthesis, Value: "("}
	case ')':
		return Token{Type: RightParenthesis, Value: ")"}
	case '[':
		return Token{Type: LeftBracket, Value: "["}
	case ']':
		return Token{Type: RightBracket, Value: "]"}
	case '{':
		return Token{Type: LeftBrace, Value: "{"}
	case '}':
		return Token{Type: RightBrace, Value: "}"}
	case ',':
		return Token{Type: Comma, Value: ","}
	case '.':
		return Token{Type: Period, Value: "."}
	case ':':
		return Token{Type: Colon, Value: ":"}
	case ';':
		return Token{Type: Semicolon, Value: ";"}
	case '=':
		return Token{Type: Equals, Value: "="}
	case '!':
		return Token{Type: Not, Value: "!"}
	case '&':
		return Token{Type: And, Value: "&"}
	case '|':
		return Token{Type: Or, Value: "|"}
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
	return Token{Type: EOF, Value: ""}
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
