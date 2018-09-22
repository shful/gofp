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
	IsPrefixKnown(prefix string) bool
	IsOWL(prefix string) bool
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
