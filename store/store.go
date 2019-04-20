package store

// store has interfaces for read-and-write container types, which hold Declarations and Axioms.
// The intentions is to optionally parse into custom container types.
// See gofp/ontologies/defaults/ for the reference implementation.

import (
	"github.com/shful/gofp/owlfunctional/annotations"
	"github.com/shful/gofp/owlfunctional/axioms"
	"github.com/shful/gofp/owlfunctional/declarations"
	"github.com/shful/gofp/owlfunctional/individual"
	"github.com/shful/gofp/owlfunctional/literal"
	"github.com/shful/gofp/owlfunctional/meta"
)

// Decls gives read access to all Declarations that were parsed yet.
type Decls interface {
	// By Key - methods:
	AnnotationPropertyDecl(ident string) (*declarations.AnnotationPropertyDecl, bool)
	ClassDecl(ident string) (*declarations.ClassDecl, bool)
	DataPropertyDecl(ident string) (*declarations.DataPropertyDecl, bool)
	DatatypeDecl(ident string) (*declarations.DatatypeDecl, bool)
	NamedIndividualDecl(ident string) (*declarations.NamedIndividualDecl, bool)
	ObjectPropertyDecl(ident string) (*declarations.ObjectPropertyDecl, bool)

	// All (as-slice) - methods:
	AllAnnotationPropertyDecls() []*declarations.AnnotationPropertyDecl
	AllClassDecls() []*declarations.ClassDecl
	AllDataPropertyDecls() []*declarations.DataPropertyDecl
	AllDatatypeDecls() []*declarations.DatatypeDecl
	AllNamedIndividualDecls() []*declarations.NamedIndividualDecl
	AllObjectPropertyDecls() []*declarations.ObjectPropertyDecl
}

type Axioms interface {
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

type DeclStore interface {
	StoreAnnotationPropertyDecl(iri string)
	StoreClassDecl(iri string)
	StoreDataPropertyDecl(iri string)
	StoreDatatypeDecl(iri string)
	StoreNamedIndividualDecl(iri string)
	StoreObjectPropertyDecl(iri string)
}

// AxiomStore takes all possible axioms and encapsulates the data structures to store them.
type AxiomStore interface {
	StoreAnnotationAssertion(A string, S string, t string)
	StoreAsymmetricObjectProperty(meta.ObjectPropertyExpression)
	StoreClassAssertion(C meta.ClassExpression, a individual.Individual)
	StoreDataPropertyAssertion(R meta.DataProperty, a individual.Individual, v literal.OWLLiteral)
	StoreFunctionalDataProperty(meta.DataProperty)
	StoreFunctionalObjectProperty(meta.ObjectPropertyExpression)
	StoreInverseFunctionalObjectProperty(meta.ObjectPropertyExpression)
	StoreInverseObjectProperties(P1, P2 meta.ObjectPropertyExpression)
	StoreIrreflexiveObjectProperty(meta.ObjectPropertyExpression)
	StoreDataPropertyDomain(R meta.DataProperty, C meta.ClassExpression)
	StoreDataPropertyRange(R meta.DataProperty, D meta.DataRange)
	StoreDisjointClasses(Cs []meta.ClassExpression)
	StoreDifferentIndividuals(as []individual.Individual)
	StoreEquivalentClasses(Cs []meta.ClassExpression)
	StoreObjectPropertyDomain(P meta.ObjectPropertyExpression, C meta.ClassExpression)
	StoreObjectPropertyRange(P meta.ObjectPropertyExpression, C meta.ClassExpression)
	StoreReflexiveObjectProperty(meta.ObjectPropertyExpression)
	StoreSubClassOf(C1, C2 meta.ClassExpression) //todo comment which is sup, which is super
	StoreSubDataPropertyOf(P1, P2 meta.DataProperty)
	StoreSubObjectPropertyOf(P1, P2 meta.ObjectPropertyExpression)
	StoreSymmetricObjectProperty(meta.ObjectPropertyExpression)
	StoreTransitiveObjectProperty(meta.ObjectPropertyExpression)
}
