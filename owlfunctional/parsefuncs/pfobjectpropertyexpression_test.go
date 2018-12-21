package parsefuncs

import (
	"testing"

	"reifenberg.de/gofp/mock"
	"reifenberg.de/gofp/owlfunctional/meta"
	"reifenberg.de/gofp/owlfunctional/parser"
	"reifenberg.de/gofp/owlfunctional/properties"
)

func TestParseObjectPropertyExpression(t *testing.T) {
	var p *parser.Parser
	var err error
	decls, prefixes := mock.NewBuilder().AddPrefixes("").AddOWLStandardPrefixes().Get()

	p = mock.NewTestParser(`owl:topObjectProperty`)
	var expr meta.ObjectPropertyExpression
	expr, err = ParseObjectPropertyExpression(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := expr.(*properties.OWLTopObjectProperty); !ok {
		t.Fatal(expr)
	}
}
