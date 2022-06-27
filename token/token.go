package token

// NOTE: int or byte might be more efficient but strings are easy to work with
type TokenType string

// NOTE: A more complex lexer would attach filenames and line numbers to tokens
// to better track down lexing and parsing errors
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func IdentType(ident string) TokenType {
	if t, ok := keywords[ident]; ok {
		return t
	}
	return IDENT
}
