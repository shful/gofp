// tech are types not required by OWL but internally.
package tech

import (
	"fmt"
	"strings"

	"github.com/shful/gofp/owlfunctional/annotations"
	"github.com/shful/gofp/owlfunctional/axioms"
	"github.com/shful/gofp/owlfunctional/declarations"
	"github.com/shful/gofp/owlfunctional/meta"
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
	GetAnnotationPropertyDecl(ident IRI) (*declarations.AnnotationPropertyDecl, bool)
	GetClassDecl(ident IRI) (*declarations.ClassDecl, bool)
	GetDataPropertyDecl(ident IRI) (*declarations.DataPropertyDecl, bool)
	GetDatatypeDecl(ident IRI) (*declarations.DatatypeDecl, bool)
	GetNamedIndividualDecl(ident IRI) (*declarations.NamedIndividualDecl, bool)
	GetObjectPropertyDecl(ident IRI) (*declarations.ObjectPropertyDecl, bool)

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
	StoreAnnotationPropertyDecl(ident IRI, decl *declarations.AnnotationPropertyDecl)
	StoreClassDecl(ident IRI, decl *declarations.ClassDecl)
	StoreDataPropertyDecl(ident IRI, decl *declarations.DataPropertyDecl)
	StoreDatatypeDecl(ident IRI, decl *declarations.DatatypeDecl)
	StoreNamedIndividualDecl(ident IRI, decl *declarations.NamedIndividualDecl)
	StoreObjectPropertyDecl(ident IRI, decl *declarations.ObjectPropertyDecl)
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

type Prefixes interface {
	// ResolvePrefix returns the IRI part which is associated with the prefix
	// false if prefix was unknown.
	ResolvePrefix(prefix string) (resolved string, ok bool)
}

//todo: Eventually remove the IRI type, it comes from my early wrong idea of a special role of fragments (RR)
// IRI resembles an IRI that OWL uses as identifier.
// In case the IRI has a fragment, it is stored in two pieces - the fragment (without hash sign), and everything before.
// The intention is that the part before the fragment, in some cases, has a meaning for itself,
// e.g. to answer the question "is this an element of the OWL namespace". While not beeing of great value, storing these two pices sometimes saves some Strings.startsWith - operations.
type IRI struct {
	// Fragment is empty for an IRI without fragment.
	// Example: for http://www.w3.org/2002/07/owl#Thing, the Fragment ist "Thing" - without the separating Hash.
	Fragment string // e.g."Thing"

	// Head + Fragment forms the whole IRI String.
	// In case there's no Fragment, Head is the whole IRI.
	// In case there is a fragment, Head MUST end with Hash (#).
	Head string // e.g."http://www.w3.org/2002/07/owl#"
}

// MustNewFragmentedIRI expects a head ending with "#".
// Panics otherwise.
func MustNewFragmentedIRI(head, fragment string) *IRI {
	if !strings.HasSuffix(head, "#") {
		panic(fmt.Sprintf("fragmented IRI needs head with Suffix '#'. (Got head=%v and fragment=%v)", head, fragment))
	}
	return &IRI{Head: head, Fragment: fragment}
}

// NewIRIFromString separates the fragment from the first part (Head), if the given value has a fragment.
// Otherwise, Fragment remains empty.
// error if val is no valid IRI. Note that most error conditions are not checked.
func NewIRIFromString(val string) (*IRI, error) {
	parts := strings.Split(val, "#")
	switch len(parts) {
	case 1: // no "#"
		return &IRI{Head: parts[0]}, nil
	case 2: // had "#" : keep the # at the end of Head
		return &IRI{Head: parts[0] + "#", Fragment: parts[1]}, nil
	default:
		return nil, fmt.Errorf("invalid IRI string with multiple # (%v)", val)
	}
}

func (s *IRI) String() string {
	return s.Head + s.Fragment
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
