package parser

import (
	"testing"
)

func TestColNo1WithTabsize(t *testing.T) {
	var p ParserPosition

	p = ParserPosition{currentLineHead: ``}
	if c := p.ColNo1WithTabsize(0); c != 1 {
		t.Fatal(c)
	}
	if c := p.ColNo1WithTabsize(1); c != 1 {
		t.Fatal(c)
	}
	if c := p.ColNo1WithTabsize(7); c != 1 {
		t.Fatal(c)
	}

	p = ParserPosition{currentLineHead: `xäö`}
	if c := p.ColNo1WithTabsize(0); c != 4 {
		t.Fatal(c)
	}

	p = ParserPosition{currentLineHead: `	x	äö	`}
	if c := p.ColNo1WithTabsize(1); c != 7 {
		t.Fatal(c)
	}
	if c := p.ColNo1WithTabsize(4); c != 16 {
		t.Fatal(c)
	}
}
