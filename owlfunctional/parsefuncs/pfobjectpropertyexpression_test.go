package parsefuncs

import (
	"testing"

	"github.com/shful/gofp/mock"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/owlfunctional/parser"
	"github.com/shful/gofp/owlfunctional/properties"
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
