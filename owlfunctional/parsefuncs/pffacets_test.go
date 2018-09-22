package parsefuncs

import (
	"testing"

	"reifenberg.de/gofp/mock"
	"reifenberg.de/gofp/owlfunctional/facets"
	"reifenberg.de/gofp/owlfunctional/parser"
)

func TestParseDatatypeRestriction(t *testing.T) {

	var p *parser.Parser
	var err error
	var expr facets.DatatypeRestriction

	decls, prefixes := mock.NewBuilder().AddPrefixes("xsd").AddClassDecl("", "CheeseTopping").Get()

	p = mock.NewTestParser(`DatatypeRestriction(xsd:integer xsd:minInclusive "400"^^xsd:integer)`)

	expr, err = parseDatatypeRestriction(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}

	dn := expr.DN
	if dn.(*facets.BuiltinDatatype).DatatypeIRI != "xsd:integer" {
		t.Fatal()
	}
	fvPairs := expr.FVPairs
	if len(fvPairs) != 1 {
		t.Fatal(fvPairs)
	}
	fvPair := fvPairs[0]
	if fvPair.V.Value != "400" {
		t.Fatal(fvPair)
	}
	if err = p.ConsumeTokens(parser.EOF); err != nil {
		t.Fatal(err)
	}
}
