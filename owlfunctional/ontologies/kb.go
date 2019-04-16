package ontologies

import (
	"github.com/shful/gofp/owlfunctional/annotations"
	"github.com/shful/gofp/owlfunctional/axioms"
	"github.com/shful/gofp/owlfunctional/declarations"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/tech"
)

// KB holds all axioms and declarations of a single ontology, as read by the parser. It's the "raw" data. i.e. has no inferred knowledge.
// The terminology "KB" (Knowledge Base) is from the java OWLAPI project.
type KB struct {
	// Declarations result each in a set[IRI string]:
	// Currently, we require explicit declaration before usage. However, OWL does not require that:
	// Although declarations are not always required, they can be used to catch obvious errors in ontologies.(https://www.w3.org/2007/OWL/wiki/Syntax#Declaration_Consistency)
	AllAnnotationPropertyDecls map[string]*declarations.AnnotationPropertyDecl
	AllClassDecls              map[string]*declarations.ClassDecl
	AllDataPropertyDecls       map[string]*declarations.DataPropertyDecl
	AllDatatypeDecls           map[string]*declarations.DatatypeDecl
	AllNamedIndividualDecls    map[string]*declarations.NamedIndividualDecl
	AllObjectPropertyDecls     map[string]*declarations.ObjectPropertyDecl

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

var _ tech.Declarations = (*KB)(nil)
var _ tech.DeclStore = (*KB)(nil)
var _ tech.Axioms = (*KB)(nil)
var _ tech.AxiomStore = (*KB)(nil)

func NewKB() *KB {
	return &KB{
		AllAnnotationPropertyDecls: make(map[string]*declarations.AnnotationPropertyDecl),
		AllClassDecls:              make(map[string]*declarations.ClassDecl),
		AllDataPropertyDecls:       make(map[string]*declarations.DataPropertyDecl),
		AllNamedIndividualDecls:    make(map[string]*declarations.NamedIndividualDecl),
		AllObjectPropertyDecls:     make(map[string]*declarations.ObjectPropertyDecl),
	}
}

func (s *KB) GetAnnotationPropertyDecl(ident tech.IRI) (decl *declarations.AnnotationPropertyDecl, ok bool) {
	decl, ok = s.AllAnnotationPropertyDecls[ident.String()]
	return
}

func (s *KB) GetClassDecl(ident tech.IRI) (decl *declarations.ClassDecl, ok bool) {
	decl, ok = s.AllClassDecls[ident.String()]
	return
}

func (s *KB) GetDataPropertyDecl(ident tech.IRI) (decl *declarations.DataPropertyDecl, ok bool) {
	decl, ok = s.AllDataPropertyDecls[ident.String()]
	return
}

func (s *KB) GetDatatypeDecl(ident tech.IRI) (decl *declarations.DatatypeDecl, ok bool) {
	decl, ok = s.AllDatatypeDecls[ident.String()]
	return
}

func (s *KB) GetNamedIndividualDecl(ident tech.IRI) (decl *declarations.NamedIndividualDecl, ok bool) {
	decl, ok = s.AllNamedIndividualDecls[ident.String()]
	return
}

func (s *KB) GetObjectPropertyDecl(ident tech.IRI) (decl *declarations.ObjectPropertyDecl, ok bool) {
	decl, ok = s.AllObjectPropertyDecls[ident.String()]
	return
}

func (s *KB) AllAnnotationAssertions() []annotations.AnnotationAssertion {
	return s.allAnnotationAssertions
}

func (s *KB) AllAsymmetricObjectProperties() []meta.ObjectPropertyExpression {
	return s.allAsymmetricObjectProperties
}

func (s *KB) AllClassAssertions() []axioms.ClassAssertion {
	return s.allClassAssertions
}

func (s *KB) AllDataPropertyAssertions() []axioms.DataPropertyAssertion {
	return s.allDataPropertyAssertions
}

func (s *KB) AllFunctionalDataProperties() []meta.DataProperty {
	return s.allFunctionalDataProperties
}

func (s *KB) AllFunctionalObjectProperties() []meta.ObjectPropertyExpression {
	return s.allFunctionalObjectProperties
}

func (s *KB) AllInverseFunctionalObjectProperties() []meta.ObjectPropertyExpression {
	return s.allInverseFunctionalObjectProperties
}

func (s *KB) AllInverseObjectProperties() []axioms.InverseObjectProperties {
	return s.allInverseObjectProperties
}

func (s *KB) AllIrreflexiveObjectProperties() []meta.ObjectPropertyExpression {
	return s.allIrreflexiveObjectProperties
}

func (s *KB) AllDataPropertyDomains() []axioms.DataPropertyDomain {
	return s.allDataPropertyDomains
}

func (s *KB) AllDataPropertyRanges() []axioms.DataPropertyRange {
	return s.allDataPropertyRanges
}

func (s *KB) AllDisjointClasses() []axioms.DisjointClasses {
	return s.allDisjointClasses
}

func (s *KB) AllDifferentIndividuals() []axioms.DifferentIndividuals {
	return s.allDifferentIndividuals
}

func (s *KB) AllEquivalentClasses() []axioms.EquivalentClasses {
	return s.allEquivalentClasses
}

func (s *KB) AllObjectPropertyDomains() []axioms.ObjectPropertyDomain {
	return s.allObjectPropertyDomains
}

func (s *KB) AllObjectPropertyRanges() []axioms.ObjectPropertyRange {
	return s.allObjectPropertyRanges
}

func (s *KB) AllReflexiveObjectProperties() []meta.ObjectPropertyExpression {
	return s.allReflexiveObjectProperties
}

func (s *KB) AllSubClassOfs() []axioms.SubClassOf {
	return s.allSubClassOfs
}

func (s *KB) AllSubDataPropertyOfs() []axioms.SubDataPropertyOf {
	return s.allSubDataPropertyOfs
}

func (s *KB) AllSubObjectPropertyOfs() []axioms.SubObjectPropertyOf {
	return s.allSubObjectPropertyOfs
}

func (s *KB) AllSymmetricObjectProperties() []meta.ObjectPropertyExpression {
	return s.allSymmetricObjectProperties
}

func (s *KB) AllTransitiveObjectProperties() []meta.ObjectPropertyExpression {
	return s.allTransitiveObjectProperties
}

func (s *KB) StoreAnnotationPropertyDecl(ident tech.IRI, decl *declarations.AnnotationPropertyDecl) {
	s.AllAnnotationPropertyDecls[ident.String()] = decl
}

func (s *KB) StoreClassDecl(ident tech.IRI, decl *declarations.ClassDecl) {
	s.AllClassDecls[ident.String()] = decl
}

func (s *KB) StoreDataPropertyDecl(ident tech.IRI, decl *declarations.DataPropertyDecl) {
	s.AllDataPropertyDecls[ident.String()] = decl
}

func (s *KB) StoreDatatypeDecl(ident tech.IRI, decl *declarations.DatatypeDecl) {
	s.AllDatatypeDecls[ident.String()] = decl
}

func (s *KB) StoreNamedIndividualDecl(ident tech.IRI, decl *declarations.NamedIndividualDecl) {
	s.AllNamedIndividualDecls[ident.String()] = decl
}

func (s *KB) StoreObjectPropertyDecl(ident tech.IRI, decl *declarations.ObjectPropertyDecl) {
	s.AllObjectPropertyDecls[ident.String()] = decl
}

func (s *KB) StoreAnnotationAssertion(a annotations.AnnotationAssertion) {
	s.allAnnotationAssertions = append(s.allAnnotationAssertions, a)
}
func (s *KB) StoreAsymmetricObjectProperty(a meta.ObjectPropertyExpression) {
	s.allAsymmetricObjectProperties = append(s.allAsymmetricObjectProperties, a)
}

func (s *KB) StoreClassAssertion(a axioms.ClassAssertion) {
	s.allClassAssertions = append(s.allClassAssertions, a)
}

func (s *KB) StoreDataPropertyAssertion(a axioms.DataPropertyAssertion) {
	s.allDataPropertyAssertions = append(s.allDataPropertyAssertions, a)
}

func (s *KB) StoreFunctionalDataProperty(a meta.DataProperty) {
	s.allFunctionalDataProperties = append(s.allFunctionalDataProperties, a)
}

func (s *KB) StoreFunctionalObjectProperty(a meta.ObjectPropertyExpression) {
	s.allFunctionalObjectProperties = append(s.allFunctionalObjectProperties, a)
}

func (s *KB) StoreInverseFunctionalObjectProperty(a meta.ObjectPropertyExpression) {
	s.allInverseFunctionalObjectProperties = append(s.allInverseFunctionalObjectProperties, a)
}

func (s *KB) StoreInverseObjectProperties(a axioms.InverseObjectProperties) {
	s.allInverseObjectProperties = append(s.allInverseObjectProperties, a)
}

func (s *KB) StoreIrreflexiveObjectProperty(a meta.ObjectPropertyExpression) {
	s.allIrreflexiveObjectProperties = append(s.allIrreflexiveObjectProperties, a)
}

func (s *KB) StoreDataPropertyDomain(a axioms.DataPropertyDomain) {
	s.allDataPropertyDomains = append(s.allDataPropertyDomains, a)
}

func (s *KB) StoreDataPropertyRange(a axioms.DataPropertyRange) {
	s.allDataPropertyRanges = append(s.allDataPropertyRanges, a)
}

func (s *KB) StoreDisjointClasses(a axioms.DisjointClasses) {
	s.allDisjointClasses = append(s.allDisjointClasses, a)
}

func (s *KB) StoreDifferentIndividuals(a axioms.DifferentIndividuals) {
	s.allDifferentIndividuals = append(s.allDifferentIndividuals, a)
}

func (s *KB) StoreEquivalentClasses(a axioms.EquivalentClasses) {
	s.allEquivalentClasses = append(s.allEquivalentClasses, a)
}

func (s *KB) StoreObjectPropertyDomain(a axioms.ObjectPropertyDomain) {
	s.allObjectPropertyDomains = append(s.allObjectPropertyDomains, a)
}

func (s *KB) StoreObjectPropertyRange(a axioms.ObjectPropertyRange) {
	s.allObjectPropertyRanges = append(s.allObjectPropertyRanges, a)
}

func (s *KB) StoreReflexiveObjectProperty(a meta.ObjectPropertyExpression) {
	s.allReflexiveObjectProperties = append(s.allReflexiveObjectProperties, a)
}

func (s *KB) StoreSubClassOf(a axioms.SubClassOf) {
	s.allSubClassOfs = append(s.allSubClassOfs, a)
}

func (s *KB) StoreSubDataPropertyOf(a axioms.SubDataPropertyOf) {
	s.allSubDataPropertyOfs = append(s.allSubDataPropertyOfs, a)
}

func (s *KB) StoreSubObjectPropertyOf(a axioms.SubObjectPropertyOf) {
	s.allSubObjectPropertyOfs = append(s.allSubObjectPropertyOfs, a)
}

func (s *KB) StoreSymmetricObjectProperty(a meta.ObjectPropertyExpression) {
	s.allSymmetricObjectProperties = append(s.allSymmetricObjectProperties, a)
}

func (s *KB) StoreTransitiveObjectProperty(a meta.ObjectPropertyExpression) {
	s.allTransitiveObjectProperties = append(s.allTransitiveObjectProperties, a)
}
