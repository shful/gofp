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
	AllAnnotationAssertions              []annotations.AnnotationAssertion
	AllAsymmetricObjectProperties        []meta.ObjectPropertyExpression
	AllClassAssertions                   []axioms.ClassAssertion
	AllDataPropertyAssertions            []axioms.DataPropertyAssertion
	AllFunctionalDataProperties          []meta.DataProperty
	AllFunctionalObjectProperties        []meta.ObjectPropertyExpression
	AllInverseFunctionalObjectProperties []meta.ObjectPropertyExpression
	AllInverseObjectProperties           []axioms.InverseObjectProperties
	AllIrreflexiveObjectProperties       []meta.ObjectPropertyExpression
	AllDataPropertyDomains               []axioms.DataPropertyDomain
	AllDataPropertyRanges                []axioms.DataPropertyRange
	AllDisjointClasses                   []axioms.DisjointClasses
	AllDifferentIndividuals              []axioms.DifferentIndividuals
	AllEquivalentClasses                 []axioms.EquivalentClasses
	AllObjectPropertyDomains             []axioms.ObjectPropertyDomain
	AllObjectPropertyRanges              []axioms.ObjectPropertyRange
	AllReflexiveObjectProperties         []meta.ObjectPropertyExpression
	AllSubClassOfs                       []axioms.SubClassOf
	AllSubDataPropertyOfs                []axioms.SubDataPropertyOf
	AllSubObjectPropertyOfs              []axioms.SubObjectPropertyOf
	AllSymmetricObjectProperties         []meta.ObjectPropertyExpression
	AllTransitiveObjectProperties        []meta.ObjectPropertyExpression
}

var _ tech.Declarations = (*KB)(nil)
var _ tech.AxiomStore = (*KB)(nil)
var _ tech.DeclStore = (*KB)(nil)

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
	s.AllAnnotationAssertions = append(s.AllAnnotationAssertions, a)
}
func (s *KB) StoreAsymmetricObjectProperty(a meta.ObjectPropertyExpression) {
	s.AllAsymmetricObjectProperties = append(s.AllAsymmetricObjectProperties, a)
}

func (s *KB) StoreClassAssertion(a axioms.ClassAssertion) {
	s.AllClassAssertions = append(s.AllClassAssertions, a)
}

func (s *KB) StoreDataPropertyAssertion(a axioms.DataPropertyAssertion) {
	s.AllDataPropertyAssertions = append(s.AllDataPropertyAssertions, a)
}

func (s *KB) StoreFunctionalDataProperty(a meta.DataProperty) {
	s.AllFunctionalDataProperties = append(s.AllFunctionalDataProperties, a)
}

func (s *KB) StoreFunctionalObjectProperty(a meta.ObjectPropertyExpression) {
	s.AllFunctionalObjectProperties = append(s.AllFunctionalObjectProperties, a)
}

func (s *KB) StoreInverseFunctionalObjectProperty(a meta.ObjectPropertyExpression) {
	s.AllInverseFunctionalObjectProperties = append(s.AllInverseFunctionalObjectProperties, a)
}

func (s *KB) StoreInverseObjectProperties(a axioms.InverseObjectProperties) {
	s.AllInverseObjectProperties = append(s.AllInverseObjectProperties, a)
}

func (s *KB) StoreIrreflexiveObjectProperty(a meta.ObjectPropertyExpression) {
	s.AllIrreflexiveObjectProperties = append(s.AllIrreflexiveObjectProperties, a)
}

func (s *KB) StoreDataPropertyDomain(a axioms.DataPropertyDomain) {
	s.AllDataPropertyDomains = append(s.AllDataPropertyDomains, a)
}

func (s *KB) StoreDataPropertyRange(a axioms.DataPropertyRange) {
	s.AllDataPropertyRanges = append(s.AllDataPropertyRanges, a)
}

func (s *KB) StoreDisjointClasses(a axioms.DisjointClasses) {
	s.AllDisjointClasses = append(s.AllDisjointClasses, a)
}

func (s *KB) StoreDifferentIndividuals(a axioms.DifferentIndividuals) {
	s.AllDifferentIndividuals = append(s.AllDifferentIndividuals, a)
}

func (s *KB) StoreEquivalentClasses(a axioms.EquivalentClasses) {
	s.AllEquivalentClasses = append(s.AllEquivalentClasses, a)
}

func (s *KB) StoreObjectPropertyDomain(a axioms.ObjectPropertyDomain) {
	s.AllObjectPropertyDomains = append(s.AllObjectPropertyDomains, a)
}

func (s *KB) StoreObjectPropertyRange(a axioms.ObjectPropertyRange) {
	s.AllObjectPropertyRanges = append(s.AllObjectPropertyRanges, a)
}

func (s *KB) StoreReflexiveObjectProperty(a meta.ObjectPropertyExpression) {
	s.AllReflexiveObjectProperties = append(s.AllReflexiveObjectProperties, a)
}

func (s *KB) StoreSubClassOf(a axioms.SubClassOf) {
	s.AllSubClassOfs = append(s.AllSubClassOfs, a)
}

func (s *KB) StoreSubDataPropertyOf(a axioms.SubDataPropertyOf) {
	s.AllSubDataPropertyOfs = append(s.AllSubDataPropertyOfs, a)
}

func (s *KB) StoreSubObjectPropertyOf(a axioms.SubObjectPropertyOf) {
	s.AllSubObjectPropertyOfs = append(s.AllSubObjectPropertyOfs, a)
}

func (s *KB) StoreSymmetricObjectProperty(a meta.ObjectPropertyExpression) {
	s.AllSymmetricObjectProperties = append(s.AllSymmetricObjectProperties, a)
}

func (s *KB) StoreTransitiveObjectProperty(a meta.ObjectPropertyExpression) {
	s.AllTransitiveObjectProperties = append(s.AllTransitiveObjectProperties, a)
}
