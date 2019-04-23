// mock are functions and types for unit tests.
package mock

import (
	"strings"

	"github.com/shful/gofp/owlfunctional/parser"
	"github.com/shful/gofp/store"
	"github.com/shful/gofp/storedefaults"
	"github.com/shful/gofp/tech"
)

type MockDeclarations struct {
	storedefaults.DeclStore
}

var _ store.Decls = (*MockDeclarations)(nil)

type MockPrefixes struct {
	data map[string]string
}

var _ tech.Prefixes = (*MockPrefixes)(nil)

func (s *MockPrefixes) ResolvePrefix(prefix string) (res string, ok bool) {
	res, ok = s.data[prefix]
	return
}

// Builder builds new mock data.
type Builder struct {
	decls    *MockDeclarations
	prefixes *MockPrefixes
}

func NewBuilder() *Builder {
	return &Builder{
		decls:    &MockDeclarations{DeclStore: *storedefaults.NewDeclStore()},
		prefixes: &MockPrefixes{map[string]string{}},
	}
}

func (s *Builder) Get() (store.Decls, tech.Prefixes) {
	return s.decls, s.prefixes
}

// AddPrefixes adds one automatic long value for each prefix P into the prefixes map.
// Each has the form "longname-for-P#"
func (s *Builder) AddPrefixes(prefixes ...string) *Builder {
	for _, prefix := range prefixes {
		s.prefixes.data[prefix] = "longname-for-" + prefix + "#"
	}
	return s
}

// AddOWLStandardPrefixes adds these often used prefixes into the mock data.
func (s *Builder) AddOWLStandardPrefixes() *Builder {
	s.prefixes.data["owl"] = `http://www.w3.org/2002/07/owl#`
	s.prefixes.data["rdf"] = `http://www.w3.org/1999/02/22-rdf-syntax-ns#`
	s.prefixes.data["rdfs"] = `http://www.w3.org/2000/01/rdf-schema#`
	s.prefixes.data["xml"] = `http://www.w3.org/XML/1998/namespace`
	s.prefixes.data["xsd"] = `http://www.w3.org/2001/XMLSchema#`
	return s
}

func (s *Builder) AddClassDecl(ident tech.IRI) *Builder {
	s.decls.StoreClassDecl(ident.String())
	return s
}

func (s *Builder) AddDataPropertyDecl(ident tech.IRI) *Builder {
	s.decls.StoreDataPropertyDecl(ident.String())
	return s
}

func (s *Builder) AddObjectPropertyDecl(ident tech.IRI) *Builder {
	s.decls.StoreObjectPropertyDecl(ident.String())
	return s
}

func NewTestParser(owl string) *parser.Parser {
	return parser.NewParser(strings.NewReader(owl), "Testparser")
}
