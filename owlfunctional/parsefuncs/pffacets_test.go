package parsefuncs

import (
	"testing"

	"github.com/shful/gofp/mock"
	"github.com/shful/gofp/owlfunctional/facets"
	"github.com/shful/gofp/owlfunctional/parser"
	"github.com/shful/gofp/tech"
)

func TestParseDatatypeRestriction(t *testing.T) {

	var p *parser.Parser
	var err error
	var expr facets.DatatypeRestriction

	decls, prefixes := mock.NewBuilder().AddOWLStandardPrefixes().
		AddClassDecl(*tech.MustNewFragmentedIRI("xxx#", "CheeseTopping")).
		Get()

	p = mock.NewTestParser(`DatatypeRestriction(xsd:integer xsd:minInclusive "400"^^xsd:integer)`)

	expr, err = parseDatatypeRestriction(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}

	dn := expr.DN
	if dn.(*facets.BuiltinDatatype).DatatypeIRI != "http://www.w3.org/2001/XMLSchema#integer" {
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
