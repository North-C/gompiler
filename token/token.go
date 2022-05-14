package token

/*
	define the Token and constants
*/

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"		// an unknown token 
	EOF = "EOF"				// end of file

	// Idetifiers and literals
	IDENT = "IDENT"
	INT = "INT" 	
	STRING = "STRING"
	
	// Operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	
	LT = "<"
	RT = ">"

	EQ = "=="
	NOT_EQ = "!="
	
	// to be define
	AND = "&"
	OR = "|"
	
	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "true"
	FALSE = "fasle"
	IF = "if"
	ELSE = "else"
	RETURN = "return"
)

/* store token types*/
var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"IDENT": IDENT,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

/* get token type according to literal */
func LookupIdent(ident string)TokenType{
	if tok, ok := keywords[ident]; ok{
		return tok
	}
	return IDENT
}

/* carete a Token with {tokenType, ch} */
func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
