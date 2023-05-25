package ast

import (
	"testing"

	"github.com/jcbbb/monkey/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Kind: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Kind: token.IDENT, Literal: "x"},
					Value: "x",
				},
				Value: &Identifier{
					Token: token.Token{Kind: token.IDENT, Literal: "y"},
					Value: "y",
				},
			},
		},
	}

	if program.String() != "let x = y;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
