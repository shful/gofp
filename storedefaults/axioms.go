package storedefaults

import (
	"github.com/shful/gofp/owlfunctional/annotations"
	"github.com/shful/gofp/owlfunctional/assertions"
	"github.com/shful/gofp/owlfunctional/axioms"
	"github.com/shful/gofp/owlfunctional/individual"
	"github.com/shful/gofp/owlfunctional/literal"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/store"
)

// AxiomStore holds all axioms and declarations of a single ontology, as read by the parser. It's the "raw" data. i.e. has no inferred knowledge.
type AxiomStore struct {

	// Axioms
	allAnnotationAssertions              []annotations.AnnotationAssertion
	allAnnotationPropertyDomains         []annotations.AnnotationPropertyDomain
	allAnnotationPropertyRanges          []annotations.AnnotationPropertyRange
	allAsymmetricObjectProperties        []meta.ObjectPropertyExpression
	allClassAssertions                   []axioms.ClassAssertion
	allDataPropertyAssertions            []axioms.DataPropertyAssertion
	allDataPropertyDomains               []axioms.DataPropertyDomain
	allDataPropertyRanges                []axioms.DataPropertyRange
	allDisjointClasses                   []axioms.DisjointClasses
	allDifferentIndividuals              []axioms.DifferentIndividuals
	allEquivalentClasses                 []axioms.EquivalentClasses
	allFunctionalDataProperties          []meta.DataProperty
	allFunctionalObjectProperties        []meta.ObjectPropertyExpression
	allInverseFunctionalObjectProperties []meta.ObjectPropertyExpression
	allInverseObjectProperties           []axioms.InverseObjectProperties
	allIrreflexiveObjectProperties       []meta.ObjectPropertyExpression
	allNegativeObjectPropertyAssertions  []assertions.NegativeObjectPropertyAssertion
	allObjectPropertyAssertions          []assertions.ObjectPropertyAssertion
	allObjectPropertyDomains             []axioms.ObjectPropertyDomain
	allObjectPropertyRanges              []axioms.ObjectPropertyRange
	allReflexiveObjectProperties         []meta.ObjectPropertyExpression
	allSubAnnotationPropertyOfs          []annotations.SubAnnotationPropertyOf
	allSubClassOfs                       []axioms.SubClassOf
	allSubDataPropertyOfs                []axioms.SubDataPropertyOf
	allSubObjectPropertyOfs              []axioms.SubObjectPropertyOf
	allSymmetricObjectProperties         []meta.ObjectPropertyExpression
	allTransitiveObjectProperties        []meta.ObjectPropertyExpression
}

var _ AllAxioms = (*AxiomStore)(nil)
var _ store.AxiomStore = (*AxiomStore)(nil)

func NewAxiomStore() *AxiomStore {
	return &AxiomStore{}
}

func (s *AxiomStore) AllAnnotationAssertions() []annotations.AnnotationAssertion {
	return s.allAnnotationAssertions
}

func (s *AxiomStore) AllAnnotationPropertyDomains() []annotations.AnnotationPropertyDomain {
	return s.allAnnotationPropertyDomains
}

func (s *AxiomStore) AllAnnotationPropertyRanges() []annotations.AnnotationPropertyRange {
	return s.allAnnotationPropertyRanges
}

func (s *AxiomStore) AllAsymmetricObjectProperties() []meta.ObjectPropertyExpression {
	return s.allAsymmetricObjectProperties
}

func (s *AxiomStore) AllClassAssertions() []axioms.ClassAssertion {
	return s.allClassAssertions
}

func (s *AxiomStore) AllDataPropertyAssertions() []axioms.DataPropertyAssertion {
	return s.allDataPropertyAssertions
}

func (s *AxiomStore) AllFunctionalDataProperties() []meta.DataProperty {
	return s.allFunctionalDataProperties
}

func (s *AxiomStore) AllFunctionalObjectProperties() []meta.ObjectPropertyExpression {
	return s.allFunctionalObjectProperties
}

func (s *AxiomStore) AllInverseFunctionalObjectProperties() []meta.ObjectPropertyExpression {
	return s.allInverseFunctionalObjectProperties
}

func (s *AxiomStore) AllInverseObjectProperties() []axioms.InverseObjectProperties {
	return s.allInverseObjectProperties
}

func (s *AxiomStore) AllIrreflexiveObjectProperties() []meta.ObjectPropertyExpression {
	return s.allIrreflexiveObjectProperties
}

func (s *AxiomStore) AllDataPropertyDomains() []axioms.DataPropertyDomain {
	return s.allDataPropertyDomains
}

func (s *AxiomStore) AllDataPropertyRanges() []axioms.DataPropertyRange {
	return s.allDataPropertyRanges
}

func (s *AxiomStore) AllDisjointClasses() []axioms.DisjointClasses {
	return s.allDisjointClasses
}

func (s *AxiomStore) AllDifferentIndividuals() []axioms.DifferentIndividuals {
	return s.allDifferentIndividuals
}

func (s *AxiomStore) AllEquivalentClasses() []axioms.EquivalentClasses {
	return s.allEquivalentClasses
}

func (s *AxiomStore) AllNegativeObjectPropertyAssertions() []assertions.NegativeObjectPropertyAssertion {
	return s.allNegativeObjectPropertyAssertions
}

func (s *AxiomStore) AllObjectPropertyAssertions() []assertions.ObjectPropertyAssertion {
	return s.allObjectPropertyAssertions
}

func (s *AxiomStore) AllObjectPropertyDomains() []axioms.ObjectPropertyDomain {
	return s.allObjectPropertyDomains
}

func (s *AxiomStore) AllObjectPropertyRanges() []axioms.ObjectPropertyRange {
	return s.allObjectPropertyRanges
}

func (s *AxiomStore) AllReflexiveObjectProperties() []meta.ObjectPropertyExpression {
	return s.allReflexiveObjectProperties
}

func (s *AxiomStore) AllSubClassOfs() []axioms.SubClassOf {
	return s.allSubClassOfs
}

func (s *AxiomStore) AllSubDataPropertyOfs() []axioms.SubDataPropertyOf {
	return s.allSubDataPropertyOfs
}

func (s *AxiomStore) AllSubObjectPropertyOfs() []axioms.SubObjectPropertyOf {
	return s.allSubObjectPropertyOfs
}

func (s *AxiomStore) AllSymmetricObjectProperties() []meta.ObjectPropertyExpression {
	return s.allSymmetricObjectProperties
}

func (s *AxiomStore) AllTransitiveObjectProperties() []meta.ObjectPropertyExpression {
	return s.allTransitiveObjectProperties
}

func (s *AxiomStore) StoreAnnotationAssertion(A meta.AnnotationProperty, S string, t string, anns []meta.Annotation) {
	s.allAnnotationAssertions = append(s.allAnnotationAssertions, annotations.AnnotationAssertion{A: A, S: S, T: t})
}

func (s *AxiomStore) StoreAnnotationPropertyDomain(A meta.AnnotationProperty, U string, anns []meta.Annotation) {
	s.allAnnotationPropertyDomains = append(s.allAnnotationPropertyDomains, annotations.AnnotationPropertyDomain{A: A, U: U})
}

func (s *AxiomStore) StoreAnnotationPropertyRange(A meta.AnnotationProperty, U string, anns []meta.Annotation) {
	s.allAnnotationPropertyRanges = append(s.allAnnotationPropertyRanges, annotations.AnnotationPropertyRange{A: A, U: U})
}

func (s *AxiomStore) StoreAsymmetricObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation) {
	s.allAsymmetricObjectProperties = append(s.allAsymmetricObjectProperties, P)
}

func (s *AxiomStore) StoreClassAssertion(C meta.ClassExpression, a individual.Individual, anns []meta.Annotation) {
	s.allClassAssertions = append(s.allClassAssertions, axioms.ClassAssertion{C: C, A: a})
}

func (s *AxiomStore) StoreDataPropertyAssertion(R meta.DataProperty, a individual.Individual, v literal.OWLLiteral, anns []meta.Annotation) {
	s.allDataPropertyAssertions = append(s.allDataPropertyAssertions, axioms.DataPropertyAssertion{R: R, A: a, V: v})
}

func (s *AxiomStore) StoreFunctionalDataProperty(a meta.DataProperty, anns []meta.Annotation) {
	s.allFunctionalDataProperties = append(s.allFunctionalDataProperties, a)
}

func (s *AxiomStore) StoreFunctionalObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation) {
	s.allFunctionalObjectProperties = append(s.allFunctionalObjectProperties, P)
}

func (s *AxiomStore) StoreInverseFunctionalObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation) {
	s.allInverseFunctionalObjectProperties = append(s.allInverseFunctionalObjectProperties, P)
}

func (s *AxiomStore) StoreInverseObjectProperties(P1, P2 meta.ObjectPropertyExpression, anns []meta.Annotation) {
	s.allInverseObjectProperties = append(s.allInverseObjectProperties, axioms.InverseObjectProperties{P1: P1, P2: P2})
}

func (s *AxiomStore) StoreIrreflexiveObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation) {
	s.allIrreflexiveObjectProperties = append(s.allIrreflexiveObjectProperties, P)
}

func (s *AxiomStore) StoreDataPropertyDomain(R meta.DataProperty, C meta.ClassExpression, anns []meta.Annotation) {
	s.allDataPropertyDomains = append(s.allDataPropertyDomains, axioms.DataPropertyDomain{R: R, C: C})
}

func (s *AxiomStore) StoreDataPropertyRange(R meta.DataProperty, D meta.DataRange, anns []meta.Annotation) {
	s.allDataPropertyRanges = append(s.allDataPropertyRanges, axioms.DataPropertyRange{R: R, D: D})
}

func (s *AxiomStore) StoreDisjointClasses(Cs []meta.ClassExpression, anns []meta.Annotation) {
	s.allDisjointClasses = append(s.allDisjointClasses, axioms.DisjointClasses{DisjointClasses: Cs})
}

func (s *AxiomStore) StoreDifferentIndividuals(as []individual.Individual, anns []meta.Annotation) {
	s.allDifferentIndividuals = append(s.allDifferentIndividuals, axioms.DifferentIndividuals{As: as})
}

func (s *AxiomStore) StoreEquivalentClasses(Cs []meta.ClassExpression, anns []meta.Annotation) {
	s.allEquivalentClasses = append(s.allEquivalentClasses, axioms.EquivalentClasses{EquivalentClasses: Cs})
}

func (s *AxiomStore) StoreNegativeObjectPropertyAssertion(P meta.ObjectPropertyExpression, a1 individual.Individual, a2 individual.Individual) {
	s.allNegativeObjectPropertyAssertions = append(s.allNegativeObjectPropertyAssertions, assertions.NegativeObjectPropertyAssertion{P, a1, a2})
}

func (s *AxiomStore) StoreObjectPropertyAssertion(PN string, a1 individual.Individual, a2 individual.Individual) {
	s.allObjectPropertyAssertions = append(s.allObjectPropertyAssertions, assertions.ObjectPropertyAssertion{PN, a1, a2})
}

func (s *AxiomStore) StoreObjectPropertyDomain(P meta.ObjectPropertyExpression, C meta.ClassExpression, anns []meta.Annotation) {
	s.allObjectPropertyDomains = append(s.allObjectPropertyDomains, axioms.ObjectPropertyDomain{P: P, C: C})
}

func (s *AxiomStore) StoreObjectPropertyRange(P meta.ObjectPropertyExpression, C meta.ClassExpression, anns []meta.Annotation) {
	s.allObjectPropertyRanges = append(s.allObjectPropertyRanges, axioms.ObjectPropertyRange{P: P, C: C})
}

func (s *AxiomStore) StoreReflexiveObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation) {
	s.allReflexiveObjectProperties = append(s.allReflexiveObjectProperties, P)
}

func (s *AxiomStore) StoreSubAnnotationPropertyOf(A1, A2 string, anns []meta.Annotation) {
	s.allSubAnnotationPropertyOfs = append(s.allSubAnnotationPropertyOfs, annotations.SubAnnotationPropertyOf{A1: A1, A2: A2})
}

func (s *AxiomStore) StoreSubClassOf(Csub, Csuper meta.ClassExpression, anns []meta.Annotation) {
	s.allSubClassOfs = append(s.allSubClassOfs, axioms.SubClassOf{C1: Csub, C2: Csuper})
}

func (s *AxiomStore) StoreSubDataPropertyOf(P1, P2 meta.DataProperty, anns []meta.Annotation) {
	s.allSubDataPropertyOfs = append(s.allSubDataPropertyOfs, axioms.SubDataPropertyOf{P1: P1, P2: P2})
}

func (s *AxiomStore) StoreSubObjectPropertyOf(P1, P2 meta.ObjectPropertyExpression, anns []meta.Annotation) {
	s.allSubObjectPropertyOfs = append(s.allSubObjectPropertyOfs, axioms.SubObjectPropertyOf{P1: P1, P2: P2})
}

func (s *AxiomStore) StoreSymmetricObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation) {
	s.allSymmetricObjectProperties = append(s.allSymmetricObjectProperties, P)
}

func (s *AxiomStore) StoreTransitiveObjectProperty(P meta.ObjectPropertyExpression, anns []meta.Annotation) {
	s.allTransitiveObjectProperties = append(s.allTransitiveObjectProperties, P)
}

type DefaultAxiom struct {
	annotations []meta.Annotation
}

func (s *DefaultAxiom) Annotations() []meta.Annotation {
	return s.annotations
}
