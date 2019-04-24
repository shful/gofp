package storedefaults

import (
	"fmt"

	"github.com/shful/gofp/owlfunctional/declarations"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/store"
)

type DeclStore struct {
	// Decls result each in a set[IRI string]:
	// Currently, we require explicit declaration before usage. However, OWL does not require that:
	// Although declarations are not always required, they can be used to catch obvious errors in ontologies.(https://www.w3.org/2007/OWL/wiki/Syntax#Declaration_Consistency)
	annotationPropertyDecls map[string]*declarations.AnnotationPropertyDecl
	classDecls              map[string]*declarations.ClassDecl
	dataPropertyDecls       map[string]*declarations.DataPropertyDecl
	datatypeDecls           map[string]*declarations.DatatypeDecl
	namedIndividualDecls    map[string]*declarations.NamedIndividualDecl
	objectPropertyDecls     map[string]*declarations.ObjectPropertyDecl

	// ExplicitDecls = true means, any declaration must be stored explicitly before it can be requested with a Get method.
	// ExplicitDecls = false is the standard OWL behaviour where a declaration is created implicitly when we request it.
	ExplicitDecls bool
}

var _ store.Decls = (*DeclStore)(nil)
var _ store.DeclStore = (*DeclStore)(nil)

func NewDeclStore() *DeclStore {
	return &DeclStore{
		annotationPropertyDecls: map[string]*declarations.AnnotationPropertyDecl{},
		classDecls:              map[string]*declarations.ClassDecl{},
		dataPropertyDecls:       map[string]*declarations.DataPropertyDecl{},
		datatypeDecls:           map[string]*declarations.DatatypeDecl{},
		namedIndividualDecls:    map[string]*declarations.NamedIndividualDecl{},
		objectPropertyDecls:     map[string]*declarations.ObjectPropertyDecl{},
		ExplicitDecls:           true,
	}
}

// === Get - methods that return a single decl by key ========

func (s *DeclStore) AnnotationPropertyDecl(ident string) (decl interface{}, ok bool) {
	decl, ok = s.annotationPropertyDecls[ident]
	return
}

func (s *DeclStore) ClassDecl(ident string) (decl meta.ClassExpression, ok bool) {
	decl, ok = s.classDecls[ident]
	if !ok && !s.ExplicitDecls {
		s.StoreClassDecl(ident)
		decl, ok = s.classDecls[ident]
	}
	return
}

func (s *DeclStore) DataPropertyDecl(ident string) (decl meta.DataProperty, ok bool) {
	decl, ok = s.dataPropertyDecls[ident]
	if !ok && !s.ExplicitDecls {
		s.StoreDataPropertyDecl(ident)
		decl, ok = s.dataPropertyDecls[ident]
	}
	return
}

func (s *DeclStore) DatatypeDecl(ident string) (decl meta.DataRange, ok bool) {
	decl, ok = s.datatypeDecls[ident]
	if !ok && !s.ExplicitDecls {
		s.StoreDatatypeDecl(ident)
		decl, ok = s.datatypeDecls[ident]
	}
	return
}

func (s *DeclStore) NamedIndividualDecl(ident string) (decl interface{}, ok bool) {
	decl, ok = s.namedIndividualDecls[ident]
	if !ok && !s.ExplicitDecls {
		s.StoreNamedIndividualDecl(ident)
		decl, ok = s.namedIndividualDecls[ident]
	}
	return
}

func (s *DeclStore) ObjectPropertyDecl(ident string) (decl meta.ObjectPropertyExpression, ok bool) {
	decl, ok = s.objectPropertyDecls[ident]
	if !ok && !s.ExplicitDecls {
		s.StoreObjectPropertyDecl(ident)
		decl, ok = s.objectPropertyDecls[ident]
	}
	return
}

// === End Get - methods ========

// === All* - methods that return slices ========
func (s *DeclStore) AllAnnotationPropertyDecls() []*declarations.AnnotationPropertyDecl {
	res := make([]*declarations.AnnotationPropertyDecl, 0, len(s.annotationPropertyDecls))
	for _, v := range s.annotationPropertyDecls {
		res = append(res, v)
	}
	return res
}
func (s *DeclStore) AllClassDecls() []*declarations.ClassDecl {
	res := make([]*declarations.ClassDecl, 0, len(s.classDecls))
	for _, v := range s.classDecls {
		res = append(res, v)
	}
	return res
}

func (s *DeclStore) AllDataPropertyDecls() []*declarations.DataPropertyDecl {
	res := make([]*declarations.DataPropertyDecl, 0, len(s.dataPropertyDecls))
	for _, v := range s.dataPropertyDecls {
		res = append(res, v)
	}
	return res
}

func (s *DeclStore) AllDatatypeDecls() []*declarations.DatatypeDecl {
	res := make([]*declarations.DatatypeDecl, 0, len(s.datatypeDecls))
	for _, v := range s.datatypeDecls {
		res = append(res, v)
	}
	return res
}

func (s *DeclStore) AllNamedIndividualDecls() []*declarations.NamedIndividualDecl {
	res := make([]*declarations.NamedIndividualDecl, 0, len(s.namedIndividualDecls))
	for _, v := range s.namedIndividualDecls {
		res = append(res, v)
	}
	return res
}

func (s *DeclStore) AllObjectPropertyDecls() []*declarations.ObjectPropertyDecl {
	res := make([]*declarations.ObjectPropertyDecl, 0, len(s.objectPropertyDecls))
	for _, v := range s.objectPropertyDecls {
		res = append(res, v)
	}
	return res
}

// === end All - methods =======

// === Store - methods =======

func (s *DeclStore) StoreAnnotationPropertyDecl(iri string) {
	s.annotationPropertyDecls[iri] = &declarations.AnnotationPropertyDecl{Declaration: declarations.Declaration{IRI: iri}}
}

func (s *DeclStore) StoreClassDecl(iri string) {
	s.classDecls[iri] = &declarations.ClassDecl{Declaration: declarations.Declaration{IRI: iri}}
}

func (s *DeclStore) StoreDataPropertyDecl(iri string) {
	s.dataPropertyDecls[iri] = &declarations.DataPropertyDecl{Declaration: declarations.Declaration{IRI: iri}}
}

func (s *DeclStore) StoreDatatypeDecl(iri string) {
	s.datatypeDecls[iri] = &declarations.DatatypeDecl{Declaration: declarations.Declaration{IRI: iri}}
}

func (s *DeclStore) StoreNamedIndividualDecl(iri string) {
	s.namedIndividualDecls[iri] = &declarations.NamedIndividualDecl{Declaration: declarations.Declaration{IRI: iri}}
}

func (s *DeclStore) StoreObjectPropertyDecl(iri string) {
	s.objectPropertyDecls[iri] = &declarations.ObjectPropertyDecl{Declaration: declarations.Declaration{IRI: iri}}
}

// === end Store - methods =======

// === Exists - methods, intended for tests only =======

func (s *DeclStore) ClassDeclExists(ident string) bool {
	_, ok := s.classDecls[ident]
	return ok
}

func (s *DeclStore) DataPropertyDeclExists(ident string) bool {
	_, ok := s.dataPropertyDecls[ident]
	return ok
}

func (s *DeclStore) NamedIndividualDeclExists(ident string) bool {
	_, ok := s.namedIndividualDecls[ident]
	return ok
}

func (s *DeclStore) ObjectPropertyDeclExists(ident string) bool {
	_, ok := s.objectPropertyDecls[ident]
	return ok
}

// === End Exists - methods =======

func (s *DeclStore) String() string {
	return fmt.Sprintf("%d annotations, %d classes, %d object properties, %d data properties, %d named individuals, %d datatypes",
		len(s.annotationPropertyDecls),
		len(s.classDecls),
		len(s.objectPropertyDecls),
		len(s.dataPropertyDecls),
		len(s.namedIndividualDecls),
		len(s.datatypeDecls),
	)
}
