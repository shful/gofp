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

	p = mock.NewTestParser(`hallo:Welt`)
	prefix, name, err = ParsePrefixedName(p)
	if err != nil {
		t.Fatal(err)
	}
	if prefix != "hallo" {
		t.Fatal(prefix)
	}
	if name != "Welt" {
		t.Fatal(name)
	}
	err = p.ConsumeTokens(parser.EOF)
	if err != nil {
		t.Fatal(err)
	}

	p = mock.NewTestParser(`:ßü_x :abc`)
	prefix, name, err = ParsePrefixedName(p)
	if err != nil {
		t.Fatal(err)
	}
	if prefix != "" {
		t.Fatal(prefix)
	}
	if name != "ßü_x" {
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
	if prefix != "hallo" {
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
	if prefix != "http://www.co-ode.org/ontologies/pizzax/pizza.owl" {
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
}
