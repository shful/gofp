package defaults

import (
	"github.com/shful/gofp/owlfunctional/annotations"
	"github.com/shful/gofp/owlfunctional/axioms"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/store"
)

//todo: Is there sense in separating Axioms and Decls, or should be join both into a KnowledgeBase (AxiomStore) as owlapi does ?
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

func (s *AxiomStore) StoreAnnotationAssertion(a annotations.AnnotationAssertion) {
	s.allAnnotationAssertions = append(s.allAnnotationAssertions, a)
}
func (s *AxiomStore) StoreAsymmetricObjectProperty(a meta.ObjectPropertyExpression) {
	s.allAsymmetricObjectProperties = append(s.allAsymmetricObjectProperties, a)
}

func (s *AxiomStore) StoreClassAssertion(a axioms.ClassAssertion) {
	s.allClassAssertions = append(s.allClassAssertions, a)
}

func (s *AxiomStore) StoreDataPropertyAssertion(a axioms.DataPropertyAssertion) {
	s.allDataPropertyAssertions = append(s.allDataPropertyAssertions, a)
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

func (s *AxiomStore) StoreInverseObjectProperties(a axioms.InverseObjectProperties) {
	s.allInverseObjectProperties = append(s.allInverseObjectProperties, a)
}

func (s *AxiomStore) StoreIrreflexiveObjectProperty(a meta.ObjectPropertyExpression) {
	s.allIrreflexiveObjectProperties = append(s.allIrreflexiveObjectProperties, a)
}

func (s *AxiomStore) StoreDataPropertyDomain(a axioms.DataPropertyDomain) {
	s.allDataPropertyDomains = append(s.allDataPropertyDomains, a)
}

func (s *AxiomStore) StoreDataPropertyRange(a axioms.DataPropertyRange) {
	s.allDataPropertyRanges = append(s.allDataPropertyRanges, a)
}

func (s *AxiomStore) StoreDisjointClasses(a axioms.DisjointClasses) {
	s.allDisjointClasses = append(s.allDisjointClasses, a)
}

func (s *AxiomStore) StoreDifferentIndividuals(a axioms.DifferentIndividuals) {
	s.allDifferentIndividuals = append(s.allDifferentIndividuals, a)
}

func (s *AxiomStore) StoreEquivalentClasses(a axioms.EquivalentClasses) {
	s.allEquivalentClasses = append(s.allEquivalentClasses, a)
}

func (s *AxiomStore) StoreObjectPropertyDomain(a axioms.ObjectPropertyDomain) {
	s.allObjectPropertyDomains = append(s.allObjectPropertyDomains, a)
}

func (s *AxiomStore) StoreObjectPropertyRange(a axioms.ObjectPropertyRange) {
	s.allObjectPropertyRanges = append(s.allObjectPropertyRanges, a)
}

func (s *AxiomStore) StoreReflexiveObjectProperty(a meta.ObjectPropertyExpression) {
	s.allReflexiveObjectProperties = append(s.allReflexiveObjectProperties, a)
}

func (s *AxiomStore) StoreSubClassOf(a axioms.SubClassOf) {
	s.allSubClassOfs = append(s.allSubClassOfs, a)
}

func (s *AxiomStore) StoreSubDataPropertyOf(a axioms.SubDataPropertyOf) {
	s.allSubDataPropertyOfs = append(s.allSubDataPropertyOfs, a)
}

func (s *AxiomStore) StoreSubObjectPropertyOf(a axioms.SubObjectPropertyOf) {
	s.allSubObjectPropertyOfs = append(s.allSubObjectPropertyOfs, a)
}

func (s *AxiomStore) StoreSymmetricObjectProperty(a meta.ObjectPropertyExpression) {
	s.allSymmetricObjectProperties = append(s.allSymmetricObjectProperties, a)
}

func (s *AxiomStore) StoreTransitiveObjectProperty(a meta.ObjectPropertyExpression) {
	s.allTransitiveObjectProperties = append(s.allTransitiveObjectProperties, a)
}
