package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Indentifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUX    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	COMMA     = ","
	SEMICOLON = ";"

	LEFTPAREN  = "("
	RIGHTPAREN = ")"
	LEFTBRACE  = "{"
	RIGHTBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	// bool
	TRUE  = "TRUE"
	FALSE = "FALSE"

	// Eq
	EQ     = "=="
	NOT_EQ = "!="
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdentifier(identifer string) TokenType {
	if tok, ok := keywords[identifer]; ok {
		return tok
	}
	return IDENT
}
