package parser

import (
	"fmt"
	"strings"
	"testing"
)

func assertToks(t *testing.T, extoks []Token, p *Parser, lineNo1, colNo1 int, currentLineHead string) {
	var pos ParserPosition
	for _, extok := range extoks {
		var tok Token
		var lit string
		tok, lit, pos = p.ScanIgnoreWSAndComment()
		if tok != extok {
			t.Fatalf("expected %v, got %v (%v)", Tokenname(extok), Tokenname(tok), lit)
		} else {
			fmt.Println(Tokenname(tok), "-", lit)
		}
	}

	// check position of the last token only
	// currentlineHead and column must be start of the last token.
	if pos.GetCurrentLineHead() != currentLineHead {
		t.Fatal(currentLineHead, pos.String())
	}

	if pos.LineNo1() != lineNo1 {
		t.Fatal(lineNo1, pos)
	}

	if pos.ColNo1() != colNo1 {
		t.Fatal(colNo1, pos)
	}
}

func TestScanIgnoreWSAndEOLAndComment(t *testing.T) {
	var parserTestString = `
Prefix(rdfs:=<http://www.w3.org/2000/01/rdf-schema#>) # Linecomment

Ontology(<urn:absolute:reifenberg.de>
# Class: :HighCaloriePizza (:HighCaloriePizza)

EquivalentClasses(:HighCaloriePizza) # shortened
)`

	var p *Parser
	p = NewParser(strings.NewReader(parserTestString), "Testdata")
	TokenLog = true

	if p.LineNo() != 0 {
		t.Fatal(p.LineNo())
	}

	if l := p.Pos().currentLineHead; l != `` {
		t.Fatal(l)
	}

	fmt.Println("Ontology")
	assertToks(t,
		[]Token{
			Prefix, B1, IDENT, COLON, EQUALS, IRI, B2,
			Ontology, B1, IRI,
		},
		p,
		4,
		10,
		`Ontology(`,
	)

	if p.LineNo() != 3 {
		t.Fatal(p.LineNo())
	}

	if p.PBal() != 1 {
		t.Fatal(p.PBal())
	}

	fmt.Println("B2")
	assertToks(t,
		[]Token{
			EquivalentClasses, B1, COLON, IDENT, B2,
		},
		p,
		7,
		36,
		`EquivalentClasses(:HighCaloriePizza`,
	)

	if p.LineNo() != 6 {
		t.Fatal(p.LineNo())
	}

	if p.PBal() != 1 {
		t.Fatal(p.PBal())
	}

	assertToks(t,
		[]Token{
			B2,
		},
		p,
		8,
		1,
		``,
	)

	if p.LineNo() != 7 {
		t.Fatal(p.LineNo())
	}

	if p.PBal() != 0 {
		t.Fatal(p.PBal())
	}
}
