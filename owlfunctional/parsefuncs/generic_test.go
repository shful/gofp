package parsefuncs

import (
	"testing"

	"reifenberg.de/gofp/tech"

	"reifenberg.de/gofp/mock"
	"reifenberg.de/gofp/owlfunctional/declarations"
	"reifenberg.de/gofp/owlfunctional/meta"
	"reifenberg.de/gofp/owlfunctional/parser"
)

func TestParseNRD(t *testing.T) {
	var p *parser.Parser
	var err error
	var n int
	var R meta.DataProperty
	var D meta.DataRange
	var isQualified bool

	decls, prefixes := mock.NewBuilder().AddOWLStandardPrefixes().AddPrefixes("").AddDataPropertyDecl(*tech.NewIRI("longname-for-", "hasPercent")).Get()

	// qualified - with D
	p = mock.NewTestParser(`(13 :hasPercent xsd:integer)`)
	n, R, D, isQualified, err = parseNRD(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if n != 13 {
		t.Fatal(n)
	}
	x := R.(*declarations.DataPropertyDecl)
	if x.IRI != "longname-for-#hasPercent" {
		t.Fatal(x.IRI)
	}
	if !D.IsNamedDatatype() {
		t.Fatal(R)
	}
	if !isQualified {
		t.Fatal(isQualified)
	}

	// unqualified
	p = mock.NewTestParser(`(13 :hasPercent)`)
	n, R, D, isQualified, err = parseNRD(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if n != 13 {
		t.Fatal(n)
	}
	if isQualified {
		t.Fatal(isQualified)
	}
}
