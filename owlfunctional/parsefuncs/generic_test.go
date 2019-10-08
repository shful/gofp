package parsefuncs

import (
	"testing"

	"github.com/shful/gofp/tech"

	"github.com/shful/gofp/mock"
	"github.com/shful/gofp/owlfunctional/decl"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/owlfunctional/parser"
)

func TestParseNRD(t *testing.T) {
	var p *parser.Parser
	var err error
	var n int
	var R meta.DataProperty
	var D meta.DataRange
	var isQualified bool

	decls, prefixes := mock.NewBuilder().AddOWLStandardPrefixes().AddPrefixes("").AddDataPropertyDecl(*tech.MustNewFragmentedIRI("longname-for-#", "hasPercent")).Get()

	// qualified - with D
	p = mock.NewTestParser(`13 :hasPercent xsd:integer)`)
	n, R, D, isQualified, err = parseNRD(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if n != 13 {
		t.Fatal(n)
	}
	x := R.(*decl.DataPropertyDecl)
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
	p = mock.NewTestParser(`13 :hasPercent)`)
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
