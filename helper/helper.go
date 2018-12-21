package helper

import (
	"reifenberg.de/gofp/owlfunctional/declarations"
	"reifenberg.de/gofp/tech"
)

type DeclarationsImpl struct {
	annotationPropertyDecls map[string]*declarations.AnnotationPropertyDecl
	classDecls              map[string]*declarations.ClassDecl
	dataPropertyDecls       map[string]*declarations.DataPropertyDecl
	datatypeDecls           map[string]*declarations.DatatypeDecl
	namedIndividualDecls    map[string]*declarations.NamedIndividualDecl
	objectPropertyDecls     map[string]*declarations.ObjectPropertyDecl
}

func NewDeclarationsImpl() *DeclarationsImpl {
	return &DeclarationsImpl{
		map[string]*declarations.AnnotationPropertyDecl{},
		map[string]*declarations.ClassDecl{},
		map[string]*declarations.DataPropertyDecl{},
		map[string]*declarations.DatatypeDecl{},
		map[string]*declarations.NamedIndividualDecl{},
		map[string]*declarations.ObjectPropertyDecl{},
	}
}

func (s *DeclarationsImpl) GetAnnotationPropertyDecl(ident tech.IRI) (decl *declarations.AnnotationPropertyDecl, ok bool) {
	decl, ok = s.annotationPropertyDecls[ident.String()]
	return
}
func (s *DeclarationsImpl) GetClassDecl(ident tech.IRI) (decl *declarations.ClassDecl, ok bool) {
	decl, ok = s.classDecls[ident.String()]
	return
}

func (s *DeclarationsImpl) GetDataPropertyDecl(ident tech.IRI) (decl *declarations.DataPropertyDecl, ok bool) {
	decl, ok = s.dataPropertyDecls[ident.String()]
	return
}

func (s *DeclarationsImpl) GetDatatypeDecl(ident tech.IRI) (decl *declarations.DatatypeDecl, ok bool) {
	decl, ok = s.datatypeDecls[ident.String()]
	return
}

func (s *DeclarationsImpl) GetNamedIndividualDecl(ident tech.IRI) (decl *declarations.NamedIndividualDecl, ok bool) {
	decl, ok = s.namedIndividualDecls[ident.String()]
	return
}

func (s *DeclarationsImpl) GetObjectPropertyDecl(ident tech.IRI) (decl *declarations.ObjectPropertyDecl, ok bool) {
	decl, ok = s.objectPropertyDecls[ident.String()]
	return
}

func (s *DeclarationsImpl) AddClassDeclForTest(prefixedName string, value *declarations.ClassDecl) {
	s.classDecls[prefixedName] = value
}

func (s *DeclarationsImpl) AddDataPropertyDeclForTest(prefixedName string, value *declarations.DataPropertyDecl) {
	s.dataPropertyDecls[prefixedName] = value
}

func (s *DeclarationsImpl) AddObjectPropertyDeclForTest(prefixedName string, value *declarations.ObjectPropertyDecl) {
	s.objectPropertyDecls[prefixedName] = value
}
