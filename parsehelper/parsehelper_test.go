package parsehelper

import (
	"testing"

	"github.com/shful/gofp/mock"
	"github.com/shful/gofp/owlfunctional/parser"
)

func TestParsePrefixedName(t *testing.T) {
	var p *parser.Parser
	var prefix, name string
	var err error

	p = mock.NewTestParser(`::hallo:Welt`)
	prefix, name, err = ParsePrefixedName(p)
	if err == nil {
		t.Fatal()
	}

	p = mock.NewTestParser(`hallo:Welt-1`)
	prefix, name, err = ParsePrefixedName(p)
	if err != nil {
		t.Fatal(err)
	}
	if prefix != "hallo" {
		t.Fatal(prefix)
	}
	if name != "Welt-1" {
		t.Fatal(name)
	}
	err = p.ConsumeTokens(parser.EOF)
	if err != nil {
		t.Fatal(err)
	}

	p = mock.NewTestParser(`:ßü_-x. :abc`)
	prefix, name, err = ParsePrefixedName(p)
	if err != nil {
		t.Fatal(err)
	}
	if prefix != "" {
		t.Fatal(prefix)
	}
	if name != "ßü_-x" {
		t.Fatal(name)
	}
}

func TestParseIRIWithFragment(t *testing.T) {
	var p *parser.Parser
	var prefix, fragment string
	var err error

	p = mock.NewTestParser(`<>`)
	prefix, fragment, err = ParseIRIWithFragment(p)
	if err == nil {
		t.Fatal(err)
	}

	p = mock.NewTestParser(`<hallo#Welt>`)
	prefix, fragment, err = ParseIRIWithFragment(p)
	if err != nil {
		t.Fatal(err)
	}
	if prefix != "hallo#" {
		t.Fatal("Prefix=" + prefix + " Fragment=" + fragment)
	}
	if fragment != "Welt" {
		t.Fatal("Prefix=" + prefix + " Fragment=" + fragment)
	}

	p = mock.NewTestParser(`<http://www.co-ode.org/ontologies/pizzax/pizza.owl#VegetarianPizzaEquivalent2>`)
	prefix, fragment, err = ParseIRIWithFragment(p)
	if err != nil {
		t.Fatal(err)
	}
	if prefix != "http://www.co-ode.org/ontologies/pizzax/pizza.owl#" {
		t.Fatal("Prefix=" + prefix + " Fragment=" + fragment)
	}
	if fragment != "VegetarianPizzaEquivalent2" {
		t.Fatal("Prefix=" + prefix + " Fragment=" + fragment)
	}
}

func TestParseUnprefixedIRI(t *testing.T) {
	var p *parser.Parser
	var iri string
	var err error

	p = mock.NewTestParser(`<http://test.de#Hello>`)
	iri, err = ParseUnprefixedIRI(p)
	if err != nil {
		t.Fatal(err)
	}
	if iri != "http://test.de#Hello" {
		t.Fatal(iri)
	}

	p = mock.NewTestParser(`<http://test.de#Hel-lo>`)
	iri, err = ParseUnprefixedIRI(p)
	if err != nil {
		t.Fatal(err)
	}
	if iri != "http://test.de#Hel-lo" {
		t.Fatal(iri)
	}
}


func TestProlongIDENT(t *testing.T) {
	var p *parser.Parser
	var suffix string

	p = mock.NewTestParser(``)
	suffix = prolongIDENT(p)
	if suffix != "" {
		t.Fatal(suffix)
	}

	p = mock.NewTestParser(` abc`)
	suffix = prolongIDENT(p)
	if suffix != "" {
		t.Fatal(suffix)
	}

	p = mock.NewTestParser(`-- abc`)
	suffix = prolongIDENT(p)
	if suffix != "--" {
		t.Fatal(suffix)
	}

	p = mock.NewTestParser(`-ß0-,`)
	suffix = prolongIDENT(p)
	if suffix != "-ß0-" {
		t.Fatal(suffix)
	}
	
	p = mock.NewTestParser(`Ä-ä-123`)
	suffix = prolongIDENT(p)
	if suffix != "Ä-ä-123" {
		t.Fatal(suffix)
	}
}