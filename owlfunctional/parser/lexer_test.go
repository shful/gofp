package parser

import (
	"fmt"
	"strings"
	"testing"
)

type tl struct {
	Tok Token
	Lit string
}

func TestScanObjectIntersectionOf(t *testing.T) {
	var s string
	var scanner *Scanner
	var tok Token = ILLEGAL
	var lit string

	s = `EquivalentClasses(:Teenager 
		ObjectIntersectionOf(:Person
			DataSomeValuesFrom(:hasAge 
			facets.DatatypeRestriction(xsd:integer 
				xsd:maxExclusive "20"^^xsd:integer 
				xsd:minExclusive "12"^^xsd:integer))))`

	scanner = NewScanner(strings.NewReader(s))
	for tok != EOF {
		tok, lit = scanner.scan()
		fmt.Println(Tokenname(tok), ":", lit)
	}
}

func TestScanComment(t *testing.T) {
	var s string

	s = `# Some Comment with ( and ) and so
Declaration()# A partial Line comment with " included`
	assertTokLits(t,
		[]tl{
			tl{LINECOMMENT, "# Some Comment with ( and ) and so"},
			tl{EOL, ""},
			tl{Declaration, ""},
			tl{B1, ""},
			tl{B2, ""},
			tl{LINECOMMENT, "# A partial Line comment with \" included"},
			tl{EOF, ""},
		},
		NewScanner(strings.NewReader(s)),
	)
}

// assertTokLit expects the given Token and Literal from the Scanner.
// Use "" to ignore the literal result.
func assertTokLit(t *testing.T, extok Token, exlit string, s *Scanner) {
	tok, lit := s.scan()
	// fmt.Printf("got (%v/%v)", Tokenname(tok), lit)
	if tok != extok {
		t.Fatalf("got %v/%v, expected %v/%v.", Tokenname(tok), lit, Tokenname(extok), exlit)
	}
	if exlit != "" && lit != exlit {
		t.Fatalf("got %v/%v, expected %v/%v.", Tokenname(tok), lit, Tokenname(extok), exlit)
	}
}

func assertTokLits(t *testing.T, tls []tl, s *Scanner) {
	for _, tl := range tls {
		assertTokLit(t, tl.Tok, tl.Lit, s)
		fmt.Printf("%v ok(%v)", Tokenname(tl.Tok), tl.Lit)
	}
}

func TestScanEOL(t *testing.T) {
	var s string

	s = "\n"
	assertTokLits(t,
		[]tl{tl{EOL, ""}, tl{EOF, ""}},
		NewScanner(strings.NewReader(s)),
	)

	s = "\r"
	assertTokLits(t,
		[]tl{tl{EOL, ""}, tl{EOF, ""}},
		NewScanner(strings.NewReader(s)),
	)

	s = "\n\n"
	assertTokLits(t,
		[]tl{tl{EOL, ""}, tl{EOL, ""}, tl{EOF, ""}},
		NewScanner(strings.NewReader(s)),
	)

	s = "\n\r"
	assertTokLits(t,
		[]tl{tl{EOL, ""}, tl{EOL, ""}, tl{EOF, ""}},
		NewScanner(strings.NewReader(s)),
	)

	s = "\r\n"
	assertTokLits(t,
		[]tl{tl{EOL, ""}, tl{EOF, ""}},
		NewScanner(strings.NewReader(s)),
	)

	s = "\r\n\r\n\r"
	assertTokLits(t,
		[]tl{tl{EOL, ""}, tl{EOL, ""}, tl{EOL, ""}, tl{EOF, ""}},
		NewScanner(strings.NewReader(s)),
	)
}

func TestScanDoubleCircum(t *testing.T) {
	var s string

	s = `^^`
	assertTokLits(t,
		[]tl{tl{DOUBLECIRCUM, ""}, tl{EOF, ""}},
		NewScanner(strings.NewReader(s)),
	)

	s = `^ ^`
	assertTokLits(t,
		[]tl{tl{ILLEGAL, "^"}, tl{WS, " "}, tl{ILLEGAL, "^"}, tl{EOF, ""}},
		NewScanner(strings.NewReader(s)),
	)
}

func TestIntegerLiteral(t *testing.T) {
	var s string

	s = `0`
	assertTokLits(t,
		[]tl{tl{INTLIT, ""}, tl{EOF, ""}},
		NewScanner(strings.NewReader(s)),
	)

	s = `999 `
	assertTokLits(t,
		[]tl{tl{INTLIT, "999"}, tl{WS, " "}, tl{EOF, ""}},
		NewScanner(strings.NewReader(s)),
	)

	s = `010`
	assertTokLits(t,
		[]tl{tl{INTLIT, "010"}, tl{EOF, ""}},
		NewScanner(strings.NewReader(s)),
	)

}
