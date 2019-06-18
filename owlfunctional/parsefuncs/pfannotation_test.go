package parsefuncs

import (
	"testing"

	"github.com/shful/gofp/mock"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/owlfunctional/parser"
)

func TestParseAnnotations(t *testing.T) {
	// parse a predeclared named class

	var p *parser.Parser
	var err error
	var exprs []meta.Annotation

	decls, prefixes := mock.NewBuilder().AddOWLStandardPrefixes().AddPrefixes("oboInOwl").Get()
	decls.ExplicitDecls=false

	p = mock.NewTestParser(``)
	exprs, err = ParseAnnotations(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if len(exprs) != 0 {
		t.Fatal(exprs)
	}

	p = mock.NewTestParser(`
Annotation(oboInOwl:hasDbXref "GOC:ai"^^xsd:string) Annotation(oboInOwl:hasDbXref "GOC:vw"^^xsd:string) obo:IAO_0000115`) //from: geneontology
	exprs, err = ParseAnnotations(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if len(exprs) != 2 {
		t.Fatal(exprs)
	}
	if exprs[0].T() != `"GOC:ai"^^http://www.w3.org/2001/XMLSchema#string` {
		//todo: makes it sense to have T as raw form with literaltype + eventually lang tags as .T, or better provide "GOC:ai" only?
		t.Fatal(exprs[0].T())
	}
}
