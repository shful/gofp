package defaults

import (
	"github.com/shful/gofp/owlfunctional/annotations"
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
	allAsymmetricObjectProperties        []meta.ObjectPropertyExpression
	allClassAssertions                   []axioms.ClassAssertion
	allDataPropertyAssertions            []axioms.DataPropertyAssertion
	allFunctionalDataProperties          []meta.DataProperty
	allFunctionalObjectProperties        []meta.ObjectPropertyExpression
	allInverseFunctionalObjectProperties []meta.ObjectPropertyExpression
	allInverseObjectProperties           []axioms.InverseObjectProperties
	allIrreflexiveObjectProperties       []meta.ObjectPropertyExpression
	allDataPropertyDomains               []axioms.DataPropertyDomain
	allDataPropertyRanges                []axioms.DataPropertyRange
	allDisjointClasses                   []axioms.DisjointClasses
	allDifferentIndividuals              []axioms.DifferentIndividuals
	allEquivalentClasses                 []axioms.EquivalentClasses
	allObjectPropertyDomains             []axioms.ObjectPropertyDomain
	allObjectPropertyRanges              []axioms.ObjectPropertyRange
	allReflexiveObjectProperties         []meta.ObjectPropertyExpression
	allSubClassOfs                       []axioms.SubClassOf
	allSubDataPropertyOfs                []axioms.SubDataPropertyOf
	allSubObjectPropertyOfs              []axioms.SubObjectPropertyOf
	allSymmetricObjectProperties         []meta.ObjectPropertyExpression
	allTransitiveObjectProperties        []meta.ObjectPropertyExpression
}

var _ store.Axioms = (*AxiomStore)(nil)
var _ store.AxiomStore = (*AxiomStore)(nil)

func NewAxiomStore() *AxiomStore {
	return &AxiomStore{}
}

func (s *AxiomStore) AllAnnotationAssertions() []annotations.AnnotationAssertion {
	return s.allAnnotationAssertions
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

func (s *AxiomStore) StoreAnnotationAssertion(A string, S string, t string) {
	s.allAnnotationAssertions = append(s.allAnnotationAssertions, annotations.AnnotationAssertion{A: A, S: S, T: t})
}
func (s *AxiomStore) StoreAsymmetricObjectProperty(a meta.ObjectPropertyExpression) {
	s.allAsymmetricObjectProperties = append(s.allAsymmetricObjectProperties, a)
}

func (s *AxiomStore) StoreClassAssertion(C meta.ClassExpression, a individual.Individual) {
	s.allClassAssertions = append(s.allClassAssertions, axioms.ClassAssertion{C: C, A: a})
}

func (s *AxiomStore) StoreDataPropertyAssertion(R meta.DataProperty, a individual.Individual, v literal.OWLLiteral) {
	s.allDataPropertyAssertions = append(s.allDataPropertyAssertions, axioms.DataPropertyAssertion{R: R, A: a, V: v})
}

func (s *AxiomStore) StoreFunctionalDataProperty(a meta.DataProperty) {
	s.allFunctionalDataProperties = append(s.allFunctionalDataProperties, a)
}

func (s *AxiomStore) StoreFunctionalObjectProperty(a meta.ObjectPropertyExpression) {
	s.allFunctionalObjectProperties = append(s.allFunctionalObjectProperties, a)
}

func (s *AxiomStore) StoreInverseFunctionalObjectProperty(a meta.ObjectPropertyExpression) {
	s.allInverseFunctionalObjectProperties = append(s.allInverseFunctionalObjectProperties, a)
}

func (s *AxiomStore) StoreInverseObjectProperties(P1, P2 meta.ObjectPropertyExpression) {
	s.allInverseObjectProperties = append(s.allInverseObjectProperties, axioms.InverseObjectProperties{P1: P1, P2: P2})
}

func (s *AxiomStore) StoreIrreflexiveObjectProperty(a meta.ObjectPropertyExpression) {
	s.allIrreflexiveObjectProperties = append(s.allIrreflexiveObjectProperties, a)
}

func (s *AxiomStore) StoreDataPropertyDomain(R meta.DataProperty, C meta.ClassExpression) {
	s.allDataPropertyDomains = append(s.allDataPropertyDomains, axioms.DataPropertyDomain{R: R, C: C})
}

func (s *AxiomStore) StoreDataPropertyRange(R meta.DataProperty, D meta.DataRange) {
	s.allDataPropertyRanges = append(s.allDataPropertyRanges, axioms.DataPropertyRange{R: R, D: D})
}

func (s *AxiomStore) StoreDisjointClasses(Cs []meta.ClassExpression) {
	s.allDisjointClasses = append(s.allDisjointClasses, axioms.DisjointClasses{DisjointClasses: Cs})
}

func (s *AxiomStore) StoreDifferentIndividuals(as []individual.Individual) {
	s.allDifferentIndividuals = append(s.allDifferentIndividuals, axioms.DifferentIndividuals{As: as})
}

func (s *AxiomStore) StoreEquivalentClasses(Cs []meta.ClassExpression) {
	s.allEquivalentClasses = append(s.allEquivalentClasses, axioms.EquivalentClasses{EquivalentClasses: Cs})
}

func (s *AxiomStore) StoreObjectPropertyDomain(P meta.ObjectPropertyExpression, C meta.ClassExpression) {
	s.allObjectPropertyDomains = append(s.allObjectPropertyDomains, axioms.ObjectPropertyDomain{P: P, C: C})
}

func (s *AxiomStore) StoreObjectPropertyRange(P meta.ObjectPropertyExpression, C meta.ClassExpression) {
	s.allObjectPropertyRanges = append(s.allObjectPropertyRanges, axioms.ObjectPropertyRange{P: P, C: C})
}

func (s *AxiomStore) StoreReflexiveObjectProperty(a meta.ObjectPropertyExpression) {
	s.allReflexiveObjectProperties = append(s.allReflexiveObjectProperties, a)
}

func (s *AxiomStore) StoreSubClassOf(C1, C2 meta.ClassExpression) {
	s.allSubClassOfs = append(s.allSubClassOfs, axioms.SubClassOf{C1: C1, C2: C2})
}

func (s *AxiomStore) StoreSubDataPropertyOf(P1, P2 meta.DataProperty) {
	s.allSubDataPropertyOfs = append(s.allSubDataPropertyOfs, axioms.SubDataPropertyOf{P1: P1, P2: P2})
}

func (s *AxiomStore) StoreSubObjectPropertyOf(P1, P2 meta.ObjectPropertyExpression) {
	s.allSubObjectPropertyOfs = append(s.allSubObjectPropertyOfs, axioms.SubObjectPropertyOf{P1: P1, P2: P2})
}

func (s *AxiomStore) StoreSymmetricObjectProperty(a meta.ObjectPropertyExpression) {
	s.allSymmetricObjectProperties = append(s.allSymmetricObjectProperties, a)
}

func (s *AxiomStore) StoreTransitiveObjectProperty(a meta.ObjectPropertyExpression) {
	s.allTransitiveObjectProperties = append(s.allTransitiveObjectProperties, a)
}
