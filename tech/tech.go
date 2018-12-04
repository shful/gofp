// tech are types not required by OWL but internally.
package tech

import (
	"fmt"

	"reifenberg.de/gofp/owlfunctional/declarations"
)

type Declarations interface {
	GetAnnotationPropertyDecl(prefix, name string) (*declarations.AnnotationPropertyDecl, bool)
	GetClassDecl(prefix, name string) (*declarations.ClassDecl, bool)
	GetDataPropertyDecl(prefix, name string) (*declarations.DataPropertyDecl, bool)
	GetDatatypeDecl(prefix, name string) (*declarations.DatatypeDecl, bool)
	GetNamedIndividualDecl(prefix, name string) (*declarations.NamedIndividualDecl, bool)
	GetObjectPropertyDecl(prefix, name string) (*declarations.ObjectPropertyDecl, bool)
}

type Prefixes interface {
	// ResolvePrefix returns the IRI part which is associated with the prefix
	// error is prefix was unknown.
	ResolvePrefix(prefix string) (string, error)
	IsPrefixKnown(prefix string) bool //todo remove because ResolvePrefix is enough
	IsOWL(prefix string) bool         //todo eventually replace by an IRI check with already resolved prefix
}

// IRI resembles an IRI that OWL uses as identifier.
// Here, it is stored in two pieces - the fragment resp."Name" of the ontology element,
// and everything before, without the separating hash (#)
type IRI struct {
	Name string // e.g."Thing"
	Head string // e.g."http://www.w3.org/2002/07/owl"
}

func NewIRI(head, name string) *IRI {
	return &IRI{Head: head, Name: name}
}

func (s IRI) String() string {
	return s.Head + "#" + s.Name
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
