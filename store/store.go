package store

// store has interfaces for container types, used by the parser to hold Declarations and Axioms.

// See the storedefaults package for the reference implementation.
//todo: some struct types are left in the interfaces, like NamedIndividualDecl, and that's wrong. For customization, store must reference interfaces only.

import (
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
