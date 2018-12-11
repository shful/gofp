// tech are types not required by OWL but internally.
package tech

import (
	"fmt"

	"reifenberg.de/gofp/owlfunctional/declarations"
)

type Declarations interface {
	GetAnnotationPropertyDecl(ident IRI) (*declarations.AnnotationPropertyDecl, bool)
	GetClassDecl(ident IRI) (*declarations.ClassDecl, bool)
	GetDataPropertyDecl(ident IRI) (*declarations.DataPropertyDecl, bool)
	GetDatatypeDecl(ident IRI) (*declarations.DatatypeDecl, bool)
	GetNamedIndividualDecl(ident IRI) (*declarations.NamedIndividualDecl, bool)
	GetObjectPropertyDecl(ident IRI) (*declarations.ObjectPropertyDecl, bool)
}

type Prefixes interface {
	// ResolvePrefix returns the IRI part which is associated with the prefix
	// false if prefix was unknown.
	ResolvePrefix(prefix string) (resolved string, ok bool)
	IsPrefixKnown(prefix string) bool //todo remove because ResolvePrefix is enough
	IsOWL(prefix string) bool         //todo eventually replace by an IRI check with already resolved prefix
}

// IRI resembles an IRI that OWL uses as identifier.
// Here, it is stored in two pieces - the fragment resp."Name" of the ontology element,
// and everything before, without the separating hash (#)
type IRI struct {
	// Name cannot be empty for a valid IRI.
	Name string // e.g."Thing"

	// Head can be empty. That means, Prefix must be given, and that can be resolved to Head.
	Head string // e.g."http://www.w3.org/2002/07/owl"

	// Prefix, if not empty, can be resolved to Head.
	Prefix string // e.g. "owl"
}

func NewIRI(head, name string) *IRI {
	return &IRI{Head: head, Name: name}
}

// NewIRIWithPrefix constructs an IRI where Head is unset but can later be resolved from Prefix.
func NewIRIWithPrefix(prefix, name string) *IRI {
	return &IRI{Prefix: prefix, Name: name}
}

func (s *IRI) String() string {
	return s.Head + "#" + s.Name
}

func (s *IRI) NeedsResolution() bool {
	return s.Prefix != "" && s.Head == ""
}

func (s *IRI) ResolveTo(name string) {
	s.Name = name
}

//todo IsOWL - functions belong somewhere else
func (s *IRI) IsOWL() bool {
	return s.Head == "http://www.w3.org/2002/07/owl"
}

func (s *IRI) IsOWLThing() bool {
	return s.IsOWL() && s.Name == "Thing"
}

// ZeroBasedPosWord is "first" for 0, then "second" ... 4th ... and so on
func ZeroBasedPosWord(i int) string {
	switch i {
	case 0:
		return "first"
	case 1:
		return "second"
	case 2:
		return "third"
	default:
		return fmt.Sprintf("%dth", i)
	}
}
