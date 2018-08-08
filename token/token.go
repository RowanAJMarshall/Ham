package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"var":    VAR,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	EQ     = "=="
	NOT_EQ = "!="

	// Identifiers and literals
	IDENT = "IDENT" // add, doThing, ConnectionManagerFactoryImpl
	INT   = "INT"   // 502

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LESSTHAN    = "<"
	GREATERTHAN = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN  = "("
	RPAREN  = ")"
	LBRACE  = "{"
	RBRACE  = "}"
	LSQUARE = "["
	RSQUARE = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	VAL      = "VAL"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
