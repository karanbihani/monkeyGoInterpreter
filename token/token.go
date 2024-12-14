package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers and literals

	IDENT = "IDENT" // add, fooobar ect
	INT = "INT" // Integer 1,2,3,4

	// Operators

	ASSIGN = "="
	PLUS = "+"

	// Delimiters

	COMMA = ","
	SEMICOLON = ";"

	LPARAN = "("
	RPARAN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords

	FUNCTION = "FUNCTION"
	LET = "LET"
)