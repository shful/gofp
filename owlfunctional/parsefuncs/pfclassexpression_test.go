package parsefuncs

import (
	"testing"

	"reifenberg.de/gofp/mock"
	"reifenberg.de/gofp/owlfunctional/classexpression"
	"reifenberg.de/gofp/owlfunctional/declarations"
	"reifenberg.de/gofp/owlfunctional/meta"
	"reifenberg.de/gofp/owlfunctional/parser"
)

func TestParseThingAndNothing(t *testing.T) {
	// parse a predeclared named class

	var p *parser.Parser
	var err error
	var expr meta.ClassExpression

	decls, prefixes := mock.NewBuilder().AddPrefixOWL().Get()

	p = mock.NewTestParser(`owl:Thing`)
	expr, err = ParseClassExpression(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := expr.(*classexpression.OWLThing); !ok {
		t.Fatal(expr)
	}

	p = mock.NewTestParser(`owl:Nothing`)
	expr, err = ParseClassExpression(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := expr.(*classexpression.OWLNothing); !ok {
		t.Fatal(expr)
	}
}

func TestParseClassDecl(t *testing.T) {
	// parse a predeclared named class

	var p *parser.Parser
	var err error
	var expr meta.ClassExpression

	decls, prefixes := mock.NewBuilder().AddPrefixes("", "xsd").AddClassDecl("", "CheeseTopping").Get()

	p = mock.NewTestParser(`:CheeseTopping`)
	expr, err = ParseClassExpression(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}

	var x *declarations.ClassDecl
	var ok bool
	if x, ok = expr.(*declarations.ClassDecl); !ok {
		t.Fatal(x)
	}
	if x.PrefixedName() != `:CheeseTopping` {
		t.Fatal(x)
	}
}
func TestParseClassExpressionCN(t *testing.T) {
	var p *parser.Parser
	var err error

	decls, prefixes := mock.NewBuilder().AddPrefixes("").AddClassDecl("", "CheeseTopping").Get()

	p = mock.NewTestParser(`:CheeseTopping`)
	parser.TokenLog = true

	var expr meta.ClassExpression
	expr, err = ParseClassExpression(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	x := expr.(*declarations.ClassDecl)
	if x.PrefixedName() != ":CheeseTopping" {
		t.Fatal(x)
	}
	err = p.ConsumeTokens(parser.EOF)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseObjectMinCardinality(t *testing.T) {
	var p *parser.Parser
	var err error

	decls, prefixes := mock.NewBuilder().AddPrefixes("").AddObjectPropertyDecl("", "hasTopping").Get()

	// var o tech.Declarations = mock.MockDecls{}
	// o.Prefixes[""] = "localprefix"
	// o.objectPropertyDecls[`:hasTopping`] = &ObjectPropertyDecl{Declaration: Declaration{Prefix: "", Name: "hasTopping"}}

	p = mock.NewTestParser(`ObjectMinCardinality(3 :hasTopping)`)
	parser.TokenLog = true

	var expr meta.ClassExpression
	expr, err = parseObjectMinCardinality(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	x := expr.(*classexpression.ObjectMinCardinality)
	if x.N != 3 {
		t.Fatal(x)
	}
	err = p.ConsumeTokens(parser.EOF)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseClassExpression_ObjectMinCardinality(t *testing.T) {
	var p *parser.Parser
	var err error
	decls, prefixes := mock.NewBuilder().AddPrefixes("").AddObjectPropertyDecl("", "hasTopping").Get()

	p = mock.NewTestParser(`ObjectMinCardinality(3 :hasTopping)`)
	parser.TokenLog = true

	var expr meta.ClassExpression
	expr, err = ParseClassExpression(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	x := expr.(*classexpression.ObjectMinCardinality)
	if x.N != 3 {
		t.Fatal(x)
	}
	err = p.ConsumeTokens(parser.EOF)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseObjectIntersectionOf(t *testing.T) {
	var p *parser.Parser
	var err error
	decls, prefixes := mock.NewBuilder().AddPrefixes("").AddClassDecl("", "Pizza").AddClassDecl("", "InterestingPizza").AddObjectPropertyDecl("", "hasTopping").Get()

	p = mock.NewTestParser(`ObjectIntersectionOf(:Pizza ObjectMinCardinality(3 :hasTopping))`)
	parser.TokenLog = true

	var expr meta.ClassExpression
	expr, err = parseObjectIntersectionOf(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	x := expr.(*classexpression.ObjectIntersectionOf)
	if len(x.Cs) != 2 {
		t.Fatal(x.Cs)
	}

	err = p.ConsumeTokens(parser.EOF)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseObjectIntersectionOf2(t *testing.T) {
	var p *parser.Parser
	var err error
	decls, prefixes := mock.NewBuilder().AddPrefixes("", "xsd").AddClassDecl("", "Pizza").AddClassDecl("", "InterestingPizza").AddDataPropertyDecl("", "hasCaloricContentValue").Get()

	p = mock.NewTestParser(`ObjectIntersectionOf(:Pizza DataHasValue(:hasCaloricContentValue "150"^^xsd:int))`)
	parser.TokenLog = true

	var expr meta.ClassExpression
	expr, err = parseObjectIntersectionOf(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	x := expr.(*classexpression.ObjectIntersectionOf)
	if len(x.Cs) != 2 {
		t.Fatal(x.Cs)
	}

	err = p.ConsumeTokens(parser.EOF)
	if err != nil {
		t.Fatal(err)
	}
}
