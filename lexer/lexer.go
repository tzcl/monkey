// Package lexer implements a lexer for Monkey.
package lexer

import (
	"github.com/dmolesUC3/emoji"
	"github.com/tzcl/monkey/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int  // points to current char
	readPosition int  // points after current char (allows us to peek)
	r            rune // current rune being processed
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readRune()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.r {
	case '=':
		if l.peekRune() == '=' {
			t = l.makeTwoRuneToken(token.EQ)
		} else {
			t = l.makeToken(token.ASSIGN)
		}
	case '+':
		t = l.makeToken(token.PLUS)
	case '-':
		t = l.makeToken(token.MINUS)
	case '!':
		if l.peekRune() == '=' {
			t = l.makeTwoRuneToken(token.NOT_EQ)
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
		if unicode.IsLetter(l.r) || isEmoji(l.r) {
			t.Literal = l.readIdentifier()
			t.Type = token.IdentType(t.Literal)
			return t
		} else if unicode.IsDigit(l.r) {
			t.Literal = l.readNumber()
			t.Type = token.INT
			return t
		} else {
			t = l.makeToken(token.ILLEGAL)
		}
	}

	l.readRune()
	return t
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.r) {
		l.readRune()
	}
}

func (l *Lexer) readRune() {
	size := 1
	if l.readPosition >= len(l.input) {
		l.r = 0 // represents "NUL"
	} else {
		l.r, size = utf8.DecodeRuneInString(l.input[l.readPosition:])
	}
	l.position = l.readPosition
	l.readPosition += size
}

func (l *Lexer) peekRune() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		r, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
		return r
	}
}

func (l *Lexer) makeToken(tokenType token.TokenType) token.Token {
	return token.Token{Type: tokenType, Literal: string(l.r)}
}

func (l *Lexer) makeTwoRuneToken(tokenType token.TokenType) token.Token {
	r := l.r
	l.readRune()
	return token.Token{Type: tokenType, Literal: string(r) + string(l.r)}
}

// readIdentifier reads the input until it reaches a non-letter character
// TODO: generalise these functions (pass read a predicate)
func (l *Lexer) readIdentifier() string {
	position := l.position
	for unicode.IsLetter(l.r) || isEmoji(l.r) {
		l.readRune()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(l.r) {
		l.readRune()
	}
	return l.input[position:l.position]
}

func isEmoji(r rune) bool {
	// NOTE: need to manually check for digits (they are considered emojis)
	if unicode.IsDigit(r) {
		return false
	}
	return emoji.IsEmoji(r)
}
