package lexer_test

import (
	"strings"
	"testing"

	"github.com/h8liu/e8/stay/lexer"
	. "github.com/h8liu/e8/stay/tokens"
)

type r struct {
	t   int
	lit string
}

func TestLexer(t *testing.T) {
	o := func(s string, exp ...*r) {
		lex := lexer.New(strings.NewReader(s))
		i := 0
		for {
			to, _, lit := lex.Scan()
			if to == EOF {
				break
			}
			// t.Logf("#%d, %s(%s)", i, lit, TokenStr(to))
			if i >= len(exp) || exp[i].t != to || exp[i].lit != lit {
				t.Errorf("lex `%s`: #%d: %s(%s)",
					s, i, lit, TokenStr(to),
				)
			}
			i++
		}

		if i != len(exp) {
			t.Errorf("lex `%s`: ntoken exp %d, got %d", s, len(exp), i)
		}

		if lex.Err() != nil {
			t.Errorf("lex `%s`: got error %s", s, lex.Err())
		}
	}

	m := func(t int, lit string) *r { return &r{t, lit} }
	n := func(t int) *r { return &r{t, TokenStr(t)} }
	id := func(s string) *r { return &r{Ident, s} }
	// eof := n(EOF)
	// sc := n(Semicolon)

	o("")
	o("   ")
	o("  =", n(Assign))
	o("  =    \t", n(Assign))
	o("=", n(Assign))
	o("==", n(Eq))
	o("3.24 = 3", m(Float, "3.24"), n(Assign), m(Int, "3"))
	o("fun", id("fun"))
	o("0X3 >= 0334", m(Int, "0X3"), n(Geq), m(Int, "0334"))
	o("3.e5", m(Float, "3.e5"))
	o("3e5", m(Float, "3e5"))
	o("3E5", m(Float, "3E5"))
	o("3D5", m(Illegal, "3D5"))
	o(".7e5", m(Float, ".7e5"))
	o("a3", id("a3"))
	o("_A3.come()", id("_A3"), n(Dot),
		id("come"), n(Lparen), n(Rparen),
	)
	o("A3(); B3(); C3();",
		id("A3"), n(Lparen), n(Rparen), n(Semicolon),
		id("B3"), n(Lparen), n(Rparen), n(Semicolon),
		id("C3"), n(Lparen), n(Rparen), n(Semicolon),
	)
	o("$", m(Illegal, "$"))
}
