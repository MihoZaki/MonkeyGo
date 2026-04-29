package lexer

import (
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"

	"github.com/MihoZaki/MonkeyGo/token"
)

type Lexer struct {
	input        []rune
	position     int
	readPosition int
	ch           rune
}

func New(input string) *Lexer {
	// Pre-allocate rune slice for large inputs (avoids intermediate string allocation)
	runes := make([]rune, 0, utf8.RuneCountInString(input))
	for _, rune := range input {
		runes = append(runes, rune)
	}
	l := &Lexer{input: runes}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++

}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = l.makeTwoCharToken(l.ch, '=', token.ASSIGN, token.EQ)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '+':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(l.ch, '=', token.PLUS, token.PLUS_ASSIGN)
		} else if l.peekChar() == '+' {
			tok = l.makeTwoCharToken(l.ch, '+', token.PLUS, token.INCREMENT)
		} else {
			tok = newToken(token.PLUS, l.ch)
		}
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = l.makeTwoCharToken(l.ch, '=', token.BANG, token.NOT_EQ)
	case '/':
		if l.peekChar() == '/' {
			l.skipComment()
			return l.NextToken()
		}
		if l.peekChar() == '*' {
			l.readChar()
			l.readChar()
			if !l.skipMultiLineComment() {
				fmt.Fprintf(os.Stderr, "warning: unclosed multi-line comment\n")
			}
			return l.NextToken()
		}
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}
	l.readChar()
	return tok

}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) skipMultiLineComment() bool {

	for l.ch != 0 {
		if l.ch == '*' && l.peekChar() == '/' {
			l.readChar()
			l.readChar()
			return true
		}
		l.readChar()
	}
	return false
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}
func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

// two character tokens can be: ==, !=, <=, >=, +=, -=, *=, --, ++
func (l *Lexer) makeTwoCharToken(firstChar, expectedSecond rune, singleCharType, twoCharType token.TokenType) token.Token {
	if l.peekChar() == expectedSecond {
		l.readChar()
		return token.Token{Type: twoCharType, Literal: string(firstChar) + string(l.ch)}
	}
	return token.Token{Type: singleCharType, Literal: string(firstChar)}
}
