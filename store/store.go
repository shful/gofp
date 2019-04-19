package store

// store has interfaces to read and write Declarations and Axioms.
// The intentions is to parse into custom structures, optionally.

import (
	"github.com/shful/gofp/owlfunctional/annotations"
	"github.com/shful/gofp/owlfunctional/axioms"
	"github.com/shful/gofp/owlfunctional/declarations"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/tech"
)

// Decls gives read access to all Declarations that were parsed yet.
type Decls interface {
	AnnotationPropertyDecl(ident string) (*declarations.AnnotationPropertyDecl, bool)
	ClassDecl(ident string) (*declarations.ClassDecl, bool)
	DataPropertyDecl(ident string) (*declarations.DataPropertyDecl, bool)
	DatatypeDecl(ident string) (*declarations.DatatypeDecl, bool)
	NamedIndividualDecl(ident string) (*declarations.NamedIndividualDecl, bool)
	ObjectPropertyDecl(ident string) (*declarations.ObjectPropertyDecl, bool)

	// old IRI-based variants of the methods above. Only one of the variants must be kept:
	GetAnnotationPropertyDecl(ident tech.IRI) (*declarations.AnnotationPropertyDecl, bool)
	GetClassDecl(ident tech.IRI) (*declarations.ClassDecl, bool)
	GetDataPropertyDecl(ident tech.IRI) (*declarations.DataPropertyDecl, bool)
	GetDatatypeDecl(ident tech.IRI) (*declarations.DatatypeDecl, bool)
	GetNamedIndividualDecl(ident tech.IRI) (*declarations.NamedIndividualDecl, bool)
	GetObjectPropertyDecl(ident tech.IRI) (*declarations.ObjectPropertyDecl, bool)

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
	StoreAnnotationPropertyDecl(ident tech.IRI, decl *declarations.AnnotationPropertyDecl)
	StoreClassDecl(ident tech.IRI, decl *declarations.ClassDecl)
	StoreDataPropertyDecl(ident tech.IRI, decl *declarations.DataPropertyDecl)
	StoreDatatypeDecl(ident tech.IRI, decl *declarations.DatatypeDecl)
	StoreNamedIndividualDecl(ident tech.IRI, decl *declarations.NamedIndividualDecl)
	StoreObjectPropertyDecl(ident tech.IRI, decl *declarations.ObjectPropertyDecl)
}

// AxiomStore takes all possible axioms and encapsulates the data structures to store them.
type AxiomStore interface {
	StoreAnnotationAssertion(annotations.AnnotationAssertion)
	StoreAsymmetricObjectProperty(meta.ObjectPropertyExpression)
	StoreClassAssertion(axioms.ClassAssertion)
	StoreDataPropertyAssertion(axioms.DataPropertyAssertion)
	StoreFunctionalDataProperty(meta.DataProperty)
	StoreFunctionalObjectProperty(meta.ObjectPropertyExpression)
	StoreInverseFunctionalObjectProperty(meta.ObjectPropertyExpression)
	StoreInverseObjectProperties(axioms.InverseObjectProperties)
	StoreIrreflexiveObjectProperty(meta.ObjectPropertyExpression)
	StoreDataPropertyDomain(axioms.DataPropertyDomain)
	StoreDataPropertyRange(axioms.DataPropertyRange)
	StoreDisjointClasses(axioms.DisjointClasses)
	StoreDifferentIndividuals(axioms.DifferentIndividuals)
	StoreEquivalentClasses(axioms.EquivalentClasses)
	StoreObjectPropertyDomain(axioms.ObjectPropertyDomain)
	StoreObjectPropertyRange(axioms.ObjectPropertyRange)
	StoreReflexiveObjectProperty(meta.ObjectPropertyExpression)
	StoreSubClassOf(axioms.SubClassOf)
	StoreSubDataPropertyOf(axioms.SubDataPropertyOf)
	StoreSubObjectPropertyOf(axioms.SubObjectPropertyOf)
	StoreSymmetricObjectProperty(meta.ObjectPropertyExpression)
	StoreTransitiveObjectProperty(meta.ObjectPropertyExpression)
}
