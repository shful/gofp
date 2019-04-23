package storedefaults

import (
	"fmt"

	"github.com/shful/gofp/owlfunctional/declarations"
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
}

var _ store.Decls = (*DeclStore)(nil)
var _ store.DeclStore = (*DeclStore)(nil)

func NewDeclStore() *DeclStore {
	return &DeclStore{
		map[string]*declarations.AnnotationPropertyDecl{},
		map[string]*declarations.ClassDecl{},
		map[string]*declarations.DataPropertyDecl{},
		map[string]*declarations.DatatypeDecl{},
		map[string]*declarations.NamedIndividualDecl{},
		map[string]*declarations.ObjectPropertyDecl{},
	}
}

func (s *DeclStore) AnnotationPropertyDecl(ident string) (decl *declarations.AnnotationPropertyDecl, ok bool) {
	decl, ok = s.annotationPropertyDecls[ident]
	return
}
func (s *DeclStore) ClassDecl(ident string) (decl *declarations.ClassDecl, ok bool) {
	decl, ok = s.classDecls[ident]
	return
}

func (s *DeclStore) DataPropertyDecl(ident string) (decl *declarations.DataPropertyDecl, ok bool) {
	decl, ok = s.dataPropertyDecls[ident]
	if !ok {
		for k, val := range s.dataPropertyDecls {
			fmt.Println("  have only:", k, "->", val)
		}
	}
	return
}

func (s *DeclStore) DatatypeDecl(ident string) (decl *declarations.DatatypeDecl, ok bool) {
	decl, ok = s.datatypeDecls[ident]
	return
}

func (s *DeclStore) NamedIndividualDecl(ident string) (decl *declarations.NamedIndividualDecl, ok bool) {
	decl, ok = s.namedIndividualDecls[ident]
	return
}

func (s *DeclStore) ObjectPropertyDecl(ident string) (decl *declarations.ObjectPropertyDecl, ok bool) {
	decl, ok = s.objectPropertyDecls[ident]
	return
}

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
