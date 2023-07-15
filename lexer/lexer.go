package lexer

import (
	"kylin/i18n"
	"unicode"
)

const EnableDecimal = false

type Lexer struct {
	data   []rune
	i18n   *i18n.Manager
	Cursor int
	Line   int
	Column int
}

func IsLetter(n rune) bool {
	return unicode.IsLetter(n)
}
func IsDigit(n rune) bool {
	return n >= '0' && n <= '9'
}

func IsRegularSymbol(n rune) bool {
	return n == '_' || n == '$'
}

func IsRegular(n rune) bool {
	return IsLetter(n) || IsRegularSymbol(n) || IsDigit(n)
}

func IsString(n rune) bool {
	return n == '"' || n == '\''
}

func NewLexer(data string, i18n *i18n.Manager) *Lexer {
	return &Lexer{data: []rune(data), Cursor: 0, Line: 1, Column: 0, i18n: i18n}
}

func (l *Lexer) NextCursor() {
	l.Cursor++
	l.Column++
}

func (l *Lexer) GetRune() rune {
	return l.data[l.Cursor]
}

func (l *Lexer) Next() Token {
	if l.Cursor >= len(l.data) {
		return Token{Type: EOF}
	}
	value := l.data[l.Cursor]
	l.NextCursor()
	switch value {
	case '+':
		if l.GetNext('=') {
			return Token{Type: PlusEquals}
		}
		return Token{Type: Addition}
	case '-':
		if l.GetNext('=') {
			return Token{Type: MinusEquals}
		}
		return Token{Type: Subtraction}
	case '*':
		if l.GetNext('=') {
			return Token{Type: TimesEquals}
		} else if l.GetNext('*') {
			if l.GetNext('=') {
				return Token{Type: ExponentEquals}
			}
			return Token{Type: Exponent}
		}
		return Token{Type: Multiplication}
	case '^':
		if l.GetNext('=') {
			l.NextCursor()
			return Token{Type: ExponentEquals}
		}
		return Token{Type: Exponent}
	case '/':
		if l.GetNext('=') {
			return Token{Type: DividedEquals}
		}
		return Token{Type: Division}
	case '%':
		if l.GetNext('=') {
			return Token{Type: ModuloEquals}
		}
		return Token{Type: Modulo}
	case '(':
		return Token{Type: LeftParenthesis}
	case ')':
		return Token{Type: RightParenthesis}
	case '[':
		return Token{Type: LeftBracket}
	case ']':
		return Token{Type: RightBracket}
	case '{':
		return Token{Type: LeftBrace}
	case '}':
		return Token{Type: RightBrace}
	case ',':
		return Token{Type: Comma}
	case '.':
		return Token{Type: Period}
	case ':':
		return Token{Type: Colon}
	case '=':
		if l.GetNext('=') {
			return Token{Type: IsEquals}
		}
		return Token{Type: Equals}
	case '!':
		if l.GetNext('=') {
			return Token{Type: NotEquals}
		}
		return Token{Type: Not}
	case '&':
		return Token{Type: And}
	case '|':
		return Token{Type: Or}
	case ' ':
		return l.Next()
	case ';', '\r', '\t':
		return Token{Type: Sep}
	case '\n':
		l.Line++
		l.Column = 0
		return Token{Type: Sep}
	case '>':
		if l.GetNext('=') {
			l.NextCursor()
			return Token{Type: GreaterThanOrEqual}
		}
		return Token{Type: GreaterThan}
	case '<':
		if l.GetNext('=') {
			l.NextCursor()
			return Token{Type: LessThanOrEqual}
		}
		return Token{Type: LessThan}
	default:
		if IsDigit(value) {
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
		if IsLetter(value) {
			identifier := l.readIdentifier(string(value))
			identifier = l.i18n.GetKeyword(identifier)
			if token, ok := KeywordMap[identifier]; ok {
				return Token{Type: token}
			}
			return Token{Type: Identifier, Value: identifier}
		}
		if IsString(value) {
			return Token{Type: String, Value: l.readString()}
		}
	}
	return Token{Type: EOF}
}

func (l *Lexer) Peek() Token {
	cursor, line, column := l.Cursor, l.Line, l.Column
	token := l.Next()
	l.Cursor, l.Line, l.Column = cursor, line, column
	return token
}

func (l *Lexer) Skip() {
	l.Next()
}

func (l *Lexer) GetNextPtr() *Token {
	token := l.Next()
	return &token
}

func (l *Lexer) HasNext() bool {
	return l.Cursor < len(l.data)
}

func (l *Lexer) GetNext(token rune) bool {
	if l.HasNext() && l.data[l.Cursor] == token {
		l.NextCursor()
		return true
	}
	return false
}

func (l *Lexer) readNumber(number string) (string, bool) {
	decimal := false
	for l.Cursor < len(l.data) {
		value := l.data[l.Cursor]
		if IsDigit(value) {
			number += string(value)
			l.NextCursor()
			continue
		}
		if value == '.' && !decimal && IsDigit(l.data[l.Cursor+1]) {
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
	for l.Cursor < len(l.data) {
		size++
		value := l.data[l.Cursor]
		if size == 1 && !IsLetter(value) {
			break
		}
		if IsRegular(value) {
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
	for l.Cursor < len(l.data) {
		value := l.data[l.Cursor]
		if IsString(value) {
			l.NextCursor()
			break
		}
		str += string(value)
		l.NextCursor()
	}
	return str
}

func (l *Lexer) ReadAll() []Token {
	var tokens []Token
	for l.HasNext() {
		tokens = append(tokens, l.Next())
	}
	return tokens
}
