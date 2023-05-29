package token

import "strconv"

type Kind int

const (
	EOF Kind = iota
	ILLEGAL
	IDENT
	INT
	ASSIGN
	PLUS
	MINUS
	ASTERISK
	SLASH
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	FUNCTION
	LET
	BANG
	LT
	GT
	TRUE
	FALSE
	IF
	ELSE
	RETURN
	EQ
	NOT_EQ
	STRING
	LBRACKET
	RBRACKET
	COLON
)

var kinds = [...]string{
	EOF:       "EOF",
	ILLEGAL:   "ILLEGAL",
	IDENT:     "IDENT",
	INT:       "INT",
	ASSIGN:    "ASSIGN",
	PLUS:      "PLUS",
	MINUS:     "MINUS",
	ASTERISK:  "ASTERISK",
	SLASH:     "SLASH",
	COMMA:     "COMMA",
	SEMICOLON: "SEMICOLON",
	LPAREN:    "LPAREN",
	RPAREN:    "RPAREN",
	LBRACE:    "LBRACE",
	RBRACE:    "RBRACE",
	FUNCTION:  "FUNCTION",
	LET:       "LET",
	BANG:      "BANG",
	LT:        "LT",
	GT:        "GT",
	TRUE:      "TRUE",
	FALSE:     "FALSE",
	IF:        "IF",
	ELSE:      "ELSE",
	RETURN:    "RETURN",
	EQ:        "EQ",
	NOT_EQ:    "NOT_EQ",
	STRING:    "STRING",
	LBRACKET:  "LBRACKET",
	RBRACKET:  "RBRACKET",
	COLON:     "COLON",
}

func (k Kind) String() string {
	s := ""
	if 0 <= k && k < Kind(len(kinds)) {
		s = kinds[k]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(k)) + ")"
	}

	return s
}

var keywords = map[string]Kind{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) Kind {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

type Token struct {
	Kind    Kind
	Literal string
}
