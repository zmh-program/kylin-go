package interpret

import (
	"kylin/utils"
)

const EnableDecimal = false

type Lexer struct {
	data   string
	cursor int
	line   int
	column int
}

func NewLexer(data string) *Lexer {
	return &Lexer{data: data, cursor: 0, line: 1, column: 0}
}

func (l *Lexer) NextCursor() {
	l.cursor++
	l.column++
}

func (l *Lexer) GetByte() byte {
	return l.data[l.cursor]
}

func (l *Lexer) Next() Token {
	if l.cursor >= len(l.data) {
		return Token{Type: EOF, Value: ""}
	}
	value := l.data[l.cursor]
	l.NextCursor()
	switch value {
	case '+':
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.NextCursor()
			return Token{Type: PlusEquals, Value: "+="}
		}
		return Token{Type: Addition, Value: "+"}
	case '-':
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.NextCursor()
			return Token{Type: MinusEquals, Value: "-="}
		}
		return Token{Type: Subtraction, Value: "-"}
	case '*':
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.NextCursor()
			return Token{Type: TimesEquals, Value: "*="}
		} else if l.cursor < len(l.data) && l.data[l.cursor] == '*' {
			l.NextCursor()
			if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
				l.NextCursor()
				return Token{Type: ExponentEquals, Value: "**="}
			}
			return Token{Type: Exponent, Value: "**"}
		}
		return Token{Type: Multiplication, Value: "*"}
	case '/':
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.NextCursor()
			return Token{Type: DividedEquals, Value: "/="}
		}
		return Token{Type: Division, Value: "/"}
	case '%':
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.NextCursor()
			return Token{Type: ModuloEquals, Value: "%="}
		}
		return Token{Type: Modulo, Value: "%"}
	case '^':
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.NextCursor()
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
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.NextCursor()
			return Token{Type: IsEquals, Value: "=="}
		}
		return Token{Type: Equals, Value: "="}
	case '!':
		if l.cursor < len(l.data) && l.data[l.cursor] == '=' {
			l.NextCursor()
			return Token{Type: NotEquals, Value: "!="}
		}
		return Token{Type: Not, Value: "!"}
	case '&':
		return Token{Type: And, Value: "&"}
	case '|':
		return Token{Type: Or, Value: "|"}
	case '\r':
		return l.Next()
	case ' ':
		return l.Next()
	case '\n':
		l.line++
		l.column = 0
		return l.Next()
	case '\t':
		return l.Next()
	default:
		if utils.IsDigit(value) {
			value, decimal := l.readNumber(string(value))
			if EnableDecimal {
				if decimal {
					return Token{Type: Float, Value: value}
				} else {
					return Token{Type: Integer, Value: value}
				}
			} else {
				return Token{Type: Float, Value: value}
			}
		}
		if utils.IsLetter(value) {
			identifier := l.readIdentifier(string(value))
			switch identifier {
			case "true":
				return Token{Type: True, Value: "true"}
			case "false":
				return Token{Type: False, Value: "false"}
			case "null":
				return Token{Type: Null, Value: "null"}
			case "fn":
				return Token{Type: Function, Value: "fn"}
			case "return":
				return Token{Type: Return, Value: "return"}
			default:
				return Token{Type: Identifier, Value: identifier}
			}
		}
		if utils.IsString(value) {
			return Token{Type: String, Value: l.readString()}
		}
	}
	return Token{Type: EOF, Value: ""}
}

func (l *Lexer) Peek() Token {
	cursor := l.cursor
	token := l.Next()
	l.cursor = cursor
	return token
}

func (l *Lexer) Skip() {
	l.Next()
}

func (l *Lexer) GetNextPtr() *Token {
	token := l.Next()
	return &token
}

func (l *Lexer) IsEnd() bool {
	return l.cursor >= len(l.data)
}

func (l *Lexer) readNumber(number string) (string, bool) {
	decimal := false
	for l.cursor < len(l.data) {
		value := l.data[l.cursor]
		if utils.IsDigit(value) {
			number += string(value)
			l.NextCursor()
			continue
		}
		if value == '.' && !decimal && utils.IsDigit(l.data[l.cursor+1]) {
			decimal = true
			number += string(value)
			l.NextCursor()
			continue
		}
		break
	}
	return number, decimal
}

func (l *Lexer) readIdentifier(identifier string) string {
	var size int
	for l.cursor < len(l.data) {
		size++
		value := l.data[l.cursor]
		if size == 1 && !utils.IsLetter(value) {
			break
		}
		if utils.IsRegular(value) {
			identifier += string(value)
			l.NextCursor()
		} else {
			break
		}
	}
	return identifier
}

func (l *Lexer) readString() string {
	var str string
	for l.cursor < len(l.data) {
		value := l.data[l.cursor]
		if utils.IsString(value) {
			l.NextCursor()
			break
		}
		str += string(value)
		l.NextCursor()
	}
	return str
}
