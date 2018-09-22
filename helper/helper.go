package helper

import (
	"reifenberg.de/gofp/owlfunctional/declarations"
	"reifenberg.de/gofp/owlfunctional/parser"
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

func (s *DeclarationsImpl) GetAnnotationPropertyDecl(prefix, name string) (decl *declarations.AnnotationPropertyDecl, ok bool) {
	decl, ok = s.annotationPropertyDecls[parser.FmtPrefixedName(prefix, name)]
	return
}
func (s *DeclarationsImpl) GetClassDecl(prefix, name string) (decl *declarations.ClassDecl, ok bool) {
	decl, ok = s.classDecls[parser.FmtPrefixedName(prefix, name)]
	return
}

func (s *DeclarationsImpl) GetDataPropertyDecl(prefix, name string) (decl *declarations.DataPropertyDecl, ok bool) {
	decl, ok = s.dataPropertyDecls[parser.FmtPrefixedName(prefix, name)]
	return
}

func (s *DeclarationsImpl) GetDatatypeDecl(prefix, name string) (decl *declarations.DatatypeDecl, ok bool) {
	decl, ok = s.datatypeDecls[parser.FmtPrefixedName(prefix, name)]
	return
}

func (s *DeclarationsImpl) GetNamedIndividualDecl(prefix, name string) (decl *declarations.NamedIndividualDecl, ok bool) {
	decl, ok = s.namedIndividualDecls[parser.FmtPrefixedName(prefix, name)]
	return
}

func (s *DeclarationsImpl) GetObjectPropertyDecl(prefix, name string) (decl *declarations.ObjectPropertyDecl, ok bool) {
	decl, ok = s.objectPropertyDecls[parser.FmtPrefixedName(prefix, name)]
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
