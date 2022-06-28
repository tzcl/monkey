// Package lexer implements a lexer for Monkey.
package lexer

import "github.com/tzcl/monkey/token"

type Lexer struct {
	input        string
	position     int  // points to current char
	readPosition int  // points after current char (allows us to peek)
	ch           byte // current char being processed
	// ch should be a rune to support Unicode
	// TODO: support this! Should be able to use Unicode and emojis
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			t = l.makeTwoCharToken(token.EQ)
		} else {
			t = l.makeToken(token.ASSIGN)
		}
	case '+':
		t = l.makeToken(token.PLUS)
	case '-':
		t = l.makeToken(token.MINUS)
	case '!':
		if l.peekChar() == '=' {
			t = l.makeTwoCharToken(token.NOT_EQ)
		} else {
			t = l.makeToken(token.BANG)
		}
	case '*':
		t = l.makeToken(token.ASTERISK)
	case '/':
		t = l.makeToken(token.SLASH)
	case '<':
		t = l.makeToken(token.LT)
	case '>':
		t = l.makeToken(token.GT)
	case ',':
		t = l.makeToken(token.COMMA)
	case ';':
		t = l.makeToken(token.SEMICOLON)
	case '(':
		t = l.makeToken(token.LPAREN)
	case ')':
		t = l.makeToken(token.RPAREN)
	case '{':
		t = l.makeToken(token.LBRACE)
	case '}':
		t = l.makeToken(token.RBRACE)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.IdentType(t.Literal)
			return t
		} else if isDigit(l.ch) {
			t.Literal = l.readNumber()
			t.Type = token.INT
			return t
		} else {
			t = l.makeToken(token.ILLEGAL)
		}
	}

	l.readChar()
	return t
}

// TODO: I feel like regex would be cleaner?
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// TODO: only supports ASCII
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII for "NUL"
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) makeToken(tokenType token.TokenType) token.Token {
	return token.Token{Type: tokenType, Literal: string(l.ch)}
}

func (l *Lexer) makeTwoCharToken(tokenType token.TokenType) token.Token {
	ch := l.ch
	l.readChar()
	return token.Token{Type: tokenType, Literal: string(ch) + string(l.ch)}
}

// readIdentifier reads the input until it reaches a non-letter character
// TODO: generalise these functions (pass read a predicate)
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
