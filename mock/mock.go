// mock are functions and types for unit tests.
package mock

import (
	"strings"

	"reifenberg.de/gofp/helper"
	"reifenberg.de/gofp/owlfunctional/declarations"
	"reifenberg.de/gofp/owlfunctional/parser"
	"reifenberg.de/gofp/tech"
)

type MockDeclarations struct {
	helper.DeclarationsImpl
}

var _ tech.Declarations = (*MockDeclarations)(nil)

type MockPrefixes struct {
	data map[string]string
}

var _ tech.Prefixes = (*MockPrefixes)(nil)

func (s *MockPrefixes) IsPrefixKnown(prefix string) bool {
	_, ok := s.data[prefix]
	return ok
}

func (s *MockPrefixes) IsOWL(prefix string) bool {
	val, _ := s.data[prefix]
	return val == `<http://www.w3.org/2002/07/owl#>`
}

// Builder builds new mock data.
type Builder struct {
	decls    *MockDeclarations
	prefixes *MockPrefixes
}

func NewBuilder() *Builder {
	return &Builder{
		decls:    &MockDeclarations{DeclarationsImpl: *helper.NewDeclarationsImpl()},
		prefixes: &MockPrefixes{map[string]string{}},
	}
}

func (s *Builder) Get() (tech.Declarations, tech.Prefixes) {
	return s.decls, s.prefixes
}

func (s *Builder) AddPrefixes(prefixes ...string) *Builder {
	for _, prefix := range prefixes {
		s.prefixes.data[prefix] = "longname-for-" + prefix
	}
	return s
}

// AddPrefixOWL adds this often used prefix into the mock data.
func (s *Builder) AddPrefixOWL() *Builder {
	s.prefixes.data["owl"] = `<http://www.w3.org/2002/07/owl#>`
	return s
}

func (s *Builder) AddClassDecl(prefix string, name string) *Builder {
	s.decls.AddClassDeclForTest(
		parser.FmtPrefixedName(prefix, name),
		&declarations.ClassDecl{Declaration: declarations.Declaration{Prefix: prefix, Name: name}},
	)
	return s
}

func (s *Builder) AddDataPropertyDecl(prefix string, name string) *Builder {
	s.decls.AddDataPropertyDeclForTest(
		parser.FmtPrefixedName(prefix, name),
		&declarations.DataPropertyDecl{Declaration: declarations.Declaration{Prefix: prefix, Name: name}},
	)
	return s
}

func (s *Builder) AddObjectPropertyDecl(prefix string, name string) *Builder {
	s.decls.AddObjectPropertyDeclForTest(
		parser.FmtPrefixedName(prefix, name),
		&declarations.ObjectPropertyDecl{Declaration: declarations.Declaration{Prefix: prefix, Name: name}},
	)
	return s
}

func NewTestParser(owl string) *parser.Parser {
	return parser.NewParser(strings.NewReader(owl), "Testparser")
}
