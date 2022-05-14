package lexer

import (
	"gompiler/token"
)

/*
	Define Lexer structure,Only support ASCII, How to support Unicode???
*/
type Lexer struct {
	input        string
	position     int  // current position in input, point to character last read
	readPosition int  // current reading position in input, point to next char
	ch           byte // current char under examination
}

/* create a Lexer from input */
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/* read char from input */
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

/* return next char in input[] */
func (l *Lexer) peekChar() byte {
	if l.readPosition > len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

/* read an identifier form input */
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

/* read number from input */
func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

/* skip whitespace */
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/* read a token from input of Lexer, return Token */
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) +
				string(l.ch)}
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)
	case '/':
		tok = token.NewToken(token.SLASH, l.ch)
	case '<':
		tok = token.NewToken(token.LT, l.ch)
	case '>':
		tok = token.NewToken(token.RT, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) +
				string(l.ch)}
		} else {
			tok = token.NewToken(token.BANG, l.ch)
		}
	case '[':
		tok = token.NewToken(token.LBRACKET, l.ch)
	case ']':
		tok = token.NewToken(token.RBRACKET, l.ch)
	case ':':
		tok = token.NewToken(token.COLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

/* if `ch` is a Letter */
func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_' || ch == '!' || ch == '?'
}

/* if `ch` is a digit */
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
