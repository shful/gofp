package parsehelper

import (
	"testing"

	"reifenberg.de/gofp/mock"
	"reifenberg.de/gofp/owlfunctional/parser"
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
