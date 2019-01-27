package parsefuncs

import (
	"testing"

	"github.com/shful/gofp/tech"

	"github.com/shful/gofp/mock"
	"github.com/shful/gofp/owlfunctional/classexpression"
	"github.com/shful/gofp/owlfunctional/declarations"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/owlfunctional/parser"
)

func TestParseThingAndNothing(t *testing.T) {
	// parse a predeclared named class

	var p *parser.Parser
	var err error
	var expr meta.ClassExpression

	decls, prefixes := mock.NewBuilder().AddOWLStandardPrefixes().Get()

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

func TestParseClassDeclWithNonemptyPrefix(t *testing.T) {
	// parse a predeclared named class

	var p *parser.Parser
	var err error
	var expr meta.ClassExpression

	decls, prefixes := mock.NewBuilder().AddPrefixes("abc").
		AddClassDecl(*tech.MustNewFragmentedIRI("longname-for-abc#", "CheeseTopping")).
		Get()

	p = mock.NewTestParser(`abc:CheeseTopping`)
	expr, err = ParseClassExpression(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}

	var x *declarations.ClassDecl
	var ok bool
	if x, ok = expr.(*declarations.ClassDecl); !ok {
		t.Fatal(x)
	}
	if x.IRI != `longname-for-abc#CheeseTopping` {
		t.Fatal(x.IRI)
	}
}

// TestParseClassDeclWithFullIRI tests IRIs with full notation. Not prefixes.
func TestParseClassDeclWithFullIRI(t *testing.T) {
	// parse a predeclared named class

	var p *parser.Parser
	var err error
	var expr meta.ClassExpression

	decls, prefixes := mock.NewBuilder().AddPrefixes("abc").
		AddClassDecl(*tech.MustNewFragmentedIRI("longname-for-abc#", "CheeseTopping")).
		AddClassDecl(tech.IRI{Fragment: "", Head: `http://www.example.org/(*§!_)someWildÜRI/`}).
		Get()

	p = mock.NewTestParser(`<longname-for-abc#CheeseTopping>`)
	expr, err = ParseClassExpression(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	var x *declarations.ClassDecl
	var ok bool
	if x, ok = expr.(*declarations.ClassDecl); !ok {
		t.Fatal(x)
	}
	if x.IRI != `longname-for-abc#CheeseTopping` {
		t.Fatal(x.IRI)
	}

	p = mock.NewTestParser(`<http://www.w3.org/2002/07/owl#Nothing>`)
	expr, err = ParseClassExpression(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := expr.(*classexpression.OWLNothing); !ok {
		t.Fatal(expr)
	}

	p = mock.NewTestParser(`<http://www.example.org/(*§!_)someWildÜRI/>`)
	expr, err = ParseClassExpression(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := expr.(*declarations.ClassDecl); !ok {
		t.Fatal(expr)
	}
}

func TestParseClassDeclWithEmptyPrefix(t *testing.T) {
	var p *parser.Parser
	var err error

	decls, prefixes := mock.NewBuilder().AddPrefixes("").
		AddClassDecl(*tech.MustNewFragmentedIRI("longname-for-#", "CheeseTopping")).
		Get()

	p = mock.NewTestParser(`:CheeseTopping`)
	parser.TokenLog = true

	var expr meta.ClassExpression
	expr, err = ParseClassExpression(p, decls, prefixes)
	if err != nil {
		t.Fatal(err)
	}
	x := expr.(*declarations.ClassDecl)
	if x.IRI != "longname-for-#CheeseTopping" {
		t.Fatal(x.IRI)
	}
	err = p.ConsumeTokens(parser.EOF)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseObjectMinCardinality(t *testing.T) {
	var p *parser.Parser
	var err error

	decls, prefixes := mock.NewBuilder().AddPrefixes("").
		AddObjectPropertyDecl(*tech.MustNewFragmentedIRI("longname-for-#", "hasTopping")).
		Get()

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
	decls, prefixes := mock.NewBuilder().AddPrefixes("").
		AddObjectPropertyDecl(*tech.MustNewFragmentedIRI("longname-for-#", "hasTopping")).
		Get()

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
	decls, prefixes := mock.NewBuilder().AddPrefixes("").
		AddClassDecl(*tech.MustNewFragmentedIRI("longname-for-#", "Pizza")).
		AddClassDecl(*tech.MustNewFragmentedIRI("xxx#", "InterestingPizza")).
		AddObjectPropertyDecl(*tech.MustNewFragmentedIRI("longname-for-#", "hasTopping")).
		Get()

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
	decls, prefixes := mock.NewBuilder().
		AddPrefixes("", "xsd").
		AddClassDecl(*tech.MustNewFragmentedIRI("longname-for-#", "Pizza")).
		AddClassDecl(*tech.MustNewFragmentedIRI("xxx#", "InterestingPizza")).
		AddDataPropertyDecl(*tech.MustNewFragmentedIRI("longname-for-#", "hasCaloricContentValue")).
		Get()

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
