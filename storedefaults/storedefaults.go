package storedefaults

// storedefaults is the default implementation for the "store" types
// which hold allparsed data.

import (
	"github.com/shful/gofp/owlfunctional/annotations"
	"github.com/shful/gofp/owlfunctional/axioms"
	"github.com/shful/gofp/owlfunctional/declarations"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/store"
)

// AllAxioms are the methods to get slices of all parsed Axioms.
type AllAxioms interface {
	AllAnnotationAssertions() []annotations.AnnotationAssertion
	AllAsymmetricObjectProperties() []meta.ObjectPropertyExpression
	AllClassAssertions() []axioms.ClassAssertion
	AllDataPropertyAssertions() []axioms.DataPropertyAssertion
	AllFunctionalDataProperties() []meta.DataProperty
	AllFunctionalObjectProperties() []meta.ObjectPropertyExpression
	AllInverseFunctionalObjectProperties() []meta.ObjectPropertyExpression
	AllInverseObjectProperties() []axioms.InverseObjectProperties
	AllIrreflexiveObjectProperties() []meta.ObjectPropertyExpression
	AllDataPropertyDomains() []axioms.DataPropertyDomain
	AllDataPropertyRanges() []axioms.DataPropertyRange
	AllDisjointClasses() []axioms.DisjointClasses
	AllDifferentIndividuals() []axioms.DifferentIndividuals
	AllEquivalentClasses() []axioms.EquivalentClasses
	AllObjectPropertyDomains() []axioms.ObjectPropertyDomain
	AllObjectPropertyRanges() []axioms.ObjectPropertyRange
	AllReflexiveObjectProperties() []meta.ObjectPropertyExpression
	AllSubClassOfs() []axioms.SubClassOf
	AllSubDataPropertyOfs() []axioms.SubDataPropertyOf
	AllSubObjectPropertyOfs() []axioms.SubObjectPropertyOf
	AllSymmetricObjectProperties() []meta.ObjectPropertyExpression
	AllTransitiveObjectProperties() []meta.ObjectPropertyExpression
}

// AllDecls are the methods to get slices of all parsed Declarations.
type AllDecls interface {
	// All (as-slice) - methods:
	AllAnnotationPropertyDecls() []*declarations.AnnotationPropertyDecl
	AllClassDecls() []*declarations.ClassDecl
	AllDataPropertyDecls() []*declarations.DataPropertyDecl
	AllDatatypeDecls() []*declarations.DatatypeDecl
	AllNamedIndividualDecls() []*declarations.NamedIndividualDecl
	AllObjectPropertyDecls() []*declarations.ObjectPropertyDecl
}

// K are all Knowledge Get methods assembled in one interface.
type K interface {
	AllDecls
	AllAxioms
	store.Decls
}

// DefaultK is the default implementation for the Knowledge interface K.
type DefaultK struct {
	AxiomStore
	DeclStore
}

func NewDefaultK() *DefaultK {
	return &DefaultK{
		AxiomStore: *NewAxiomStore(),
		DeclStore:  *NewDeclStore(),
	}
}

var _ K = (*DefaultK)(nil)
