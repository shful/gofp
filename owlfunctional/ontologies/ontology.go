package ontologies

import (
	"fmt"

	"github.com/shful/gofp/owlfunctional/annotations"
	"github.com/shful/gofp/owlfunctional/axioms"
	"github.com/shful/gofp/owlfunctional/declarations"
	"github.com/shful/gofp/owlfunctional/individual"
	"github.com/shful/gofp/owlfunctional/literal"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/owlfunctional/parsefuncs"
	"github.com/shful/gofp/owlfunctional/parser"
	"github.com/shful/gofp/parsehelper"
	"github.com/shful/gofp/tech"
)

type Ontology struct {
	IRI        string
	VERSIONIRI string
	Prefixes   map[string]string

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

var _ tech.Declarations = (*Ontology)(nil)
var _ tech.Prefixes = (*Ontology)(nil)

func NewOntology(prefixes map[string]string) (res *Ontology) {
	res = &Ontology{Prefixes: prefixes}
	res.AllAnnotationPropertyDecls = make(map[string]*declarations.AnnotationPropertyDecl)
	res.AllClassDecls = make(map[string]*declarations.ClassDecl)
	res.AllDataPropertyDecls = make(map[string]*declarations.DataPropertyDecl)
	res.AllNamedIndividualDecls = make(map[string]*declarations.NamedIndividualDecl)
	res.AllObjectPropertyDecls = make(map[string]*declarations.ObjectPropertyDecl)
	return
}

func (s *Ontology) GetAnnotationPropertyDecl(ident tech.IRI) (decl *declarations.AnnotationPropertyDecl, ok bool) {
	decl, ok = s.AllAnnotationPropertyDecls[ident.String()]
	return
}
func (s *Ontology) GetClassDecl(ident tech.IRI) (decl *declarations.ClassDecl, ok bool) {
	decl, ok = s.AllClassDecls[ident.String()]
	return
}

func (s *Ontology) GetDataPropertyDecl(ident tech.IRI) (decl *declarations.DataPropertyDecl, ok bool) {
	decl, ok = s.AllDataPropertyDecls[ident.String()]
	return
}

func (s *Ontology) GetDatatypeDecl(ident tech.IRI) (decl *declarations.DatatypeDecl, ok bool) {
	decl, ok = s.AllDatatypeDecls[ident.String()]
	return
}

func (s *Ontology) GetNamedIndividualDecl(ident tech.IRI) (decl *declarations.NamedIndividualDecl, ok bool) {
	decl, ok = s.AllNamedIndividualDecls[ident.String()]
	return
}

func (s *Ontology) GetObjectPropertyDecl(ident tech.IRI) (decl *declarations.ObjectPropertyDecl, ok bool) {
	decl, ok = s.AllObjectPropertyDecls[ident.String()]
	return
}

// Parse consumes "Ontology(...)" with both enclosing braces.
func (s *Ontology) Parse(p *parser.Parser) (err error) {
	var initialPBal = p.PBal()
	var pos parser.ParserPosition
	if err = p.ConsumeTokens(parser.Ontology, parser.B1); err != nil {
		return pos.EnrichErrorMsg(err, "Parsing Ontology element:%v")
	}

	// IRI as name is possible
	tok, lit, pos := p.ScanIgnoreWSAndComment()
	if tok == parser.IRI {
		// there's an IRI as ontology name
		s.IRI = lit
		tok, lit, _ := p.ScanIgnoreWSAndComment()
		if tok == parser.IRI {
			// there's another IRI as ontology version
			s.VERSIONIRI = lit
		} else {
			p.Unscan()
		}
	} else {
		p.Unscan()
	}

	for p.PBal() > initialPBal {
		tok, lit, pos := p.ScanIgnoreWSAndComment()
		switch tok {
		case parser.B2:
			// must be the end of the Ontology() expression
			if p.PBal() < initialPBal {
				panic(pos.Errorf("internal: %v<%v", p.PBal(), initialPBal))
			}
			return
		}
		p.Unscan()

		switch tok {
		case parser.AnnotationAssertion:
			err = s.parseAnnotationAssertion(p)
		case parser.AsymmetricObjectProperty:
			err = s.parseAsymmetricObjectProperty(p)
		case parser.ClassAssertion:
			err = s.parseClassAssertion(p)
		case parser.DataPropertyAssertion:
			err = s.parseDataPropertyAssertion(p)
		case parser.Declaration:
			err = s.parseDeclaration(p)
		case parser.DataPropertyDomain:
			err = s.parseDataPropertyDomain(p)
		case parser.DataPropertyRange:
			err = s.parseDataPropertyRange(p)
		case parser.DifferentIndividuals:
			err = s.parseDifferentIndividuals(p)
		case parser.DisjointClasses:
			err = s.parseDisjointClasses(p)
		case parser.EquivalentClasses:
			err = s.parseEquivalentClasses(p)
		case parser.FunctionalDataProperty:
			err = s.parseFunctionalDataProperty(p)
		case parser.FunctionalObjectProperty:
			err = s.parseFunctionalObjectProperty(p)
		case parser.InverseFunctionalObjectProperty:
			err = s.parseInverseFunctionalObjectProperty(p)
		case parser.InverseObjectProperties:
			err = s.parseInverseObjectProperties(p)
		case parser.IrreflexiveObjectProperty:
			err = s.parseIrreflexiveObjectProperty(p)
		case parser.ObjectPropertyDomain:
			err = s.parseObjectPropertyDomain(p)
		case parser.ObjectPropertyRange:
			err = s.parseObjectPropertyRange(p)
		case parser.ReflexiveObjectProperty:
			err = s.parseReflexiveObjectProperty(p)
		case parser.SubClassOf:
			err = s.parseSubClassOf(p)
		case parser.SubDataPropertyOf:
			err = s.parseSubDataPropertyOf(p)
		case parser.SubObjectPropertyOf:
			err = s.parseSubObjectPropertyOf(p)
		case parser.SymmetricObjectProperty:
			s.parseSymmetricObjectProperty(p)
		case parser.TransitiveObjectProperty:
			err = s.parseTransitiveObjectProperty(p)
		default:
			err = pos.Errorf(`unexpected ontology token %v ("%v")`, parser.Tokenname(tok), lit)
		}

		if err != nil {
			return
		}
	}

	return
}

// parseAnnotationAssertion
// - should not parse individuals into strings but maintain these individuals and reference them
func (s *Ontology) parseAnnotationAssertion(p *parser.Parser) (err error) {

	if err = p.ConsumeTokens(parser.AnnotationAssertion, parser.B1); err != nil {
		return
	}
	pos := p.Pos()
	var ident *tech.IRI

	ident, err = parsehelper.ParseAndResolveIRI(p, s)
	if err != nil {
		err = pos.EnrichErrorMsg(err, "reading 1st param in AnnotationAssertion")
		return
	}
	var s_ string
	s_, _, err = parsefuncs.Parses(p, s, s)
	if err != nil {
		err = pos.EnrichErrorMsg(err, "reading 2nd param in AnnotationAssertion")
		return
	}
	var t string

	t, _, err = parsefuncs.Parset(p, s, s)
	if err != nil {
		err = pos.EnrichErrorMsg(err, "reading 3rd param in AnnotationAssertion")
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	s.AllAnnotationAssertions = append(s.AllAnnotationAssertions, annotations.AnnotationAssertion{
		A: ident.String(),
		S: s_,
		T: t,
	})
	return
}

func (s *Ontology) parseAsymmetricObjectProperty(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.AsymmetricObjectProperty); err != nil {
		return
	}
	var P meta.ObjectPropertyExpression
	P, err = s.parseP(p)
	if err != nil {
		return
	}
	s.AllAsymmetricObjectProperties = append(s.AllAsymmetricObjectProperties, P)
	return
}

func (s *Ontology) parseClassAssertion(p *parser.Parser) (err error) {

	if err = p.ConsumeTokens(parser.ClassAssertion, parser.B1); err != nil {
		return
	}
	var C meta.ClassExpression
	C, err = parsefuncs.ParseClassExpression(p, s, s)
	if err != nil {
		return
	}
	var a individual.Individual
	a, err = parsefuncs.ParseIndividual(p, s, s)
	if err != nil {
		return
	}
	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	s.AllClassAssertions = append(s.AllClassAssertions, axioms.ClassAssertion{C: C, A: a})
	return
}

func (s *Ontology) parseDataPropertyAssertion(p *parser.Parser) (err error) {

	if err = p.ConsumeTokens(parser.DataPropertyAssertion, parser.B1); err != nil {
		return
	}
	pos := p.Pos()
	var R meta.DataProperty
	R, err = parsefuncs.ParseDataProperty(p, s, s)
	if err != nil {
		err = pos.EnrichErrorMsg(err, "1st param in DataPropertyAssertion")
		return
	}
	var a individual.Individual
	a, err = parsefuncs.ParseIndividual(p, s, s)
	if err != nil {
		err = pos.EnrichErrorMsg(err, "2nd param in DataPropertyAssertion")
		return
	}
	var v literal.OWLLiteral
	v, err = parsefuncs.ParseOWLLiteral(p, s)
	if err != nil {
		err = pos.EnrichErrorMsg(err, "3rd param in DataPropertyAssertion")
		return
	}
	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	s.AllDataPropertyAssertions = append(s.AllDataPropertyAssertions, axioms.DataPropertyAssertion{R: R, A: a, V: v})
	return
}

func (s *Ontology) parseDeclaration(p *parser.Parser) (err error) {
	var ident *tech.IRI

	if err = p.ConsumeTokens(parser.Declaration, parser.B1); err != nil {
		return
	}
	tok, _, _ := p.ScanIgnoreWSAndComment()
	switch tok {
	case parser.AnnotationProperty:
		if ident, err = s.parseBracedIRI(p); err != nil {
			return
		}
		s.AllAnnotationPropertyDecls[ident.String()] = &declarations.AnnotationPropertyDecl{Declaration: declarations.Declaration{IRI: ident.String()}}
	case parser.Class:
		if ident, err = s.parseBracedIRI(p); err != nil {
			return
		}
		s.AllClassDecls[ident.String()] = &declarations.ClassDecl{Declaration: declarations.Declaration{IRI: ident.String()}}
	case parser.DataProperty:
		if ident, err = s.parseBracedIRI(p); err != nil {
			return
		}
		s.AllDataPropertyDecls[ident.String()] = &declarations.DataPropertyDecl{Declaration: declarations.Declaration{IRI: ident.String()}}
	case parser.Datatype:
		if ident, err = s.parseBracedIRI(p); err != nil {
			return
		}
		s.AllDatatypeDecls[ident.String()] = &declarations.DatatypeDecl{Declaration: declarations.Declaration{IRI: ident.String()}}
	case parser.NamedIndividual:
		if ident, err = s.parseBracedIRI(p); err != nil {
			return
		}
		s.AllNamedIndividualDecls[ident.String()] = &declarations.NamedIndividualDecl{Declaration: declarations.Declaration{IRI: ident.String()}}
	case parser.ObjectProperty:
		if ident, err = s.parseBracedIRI(p); err != nil {
			return
		}
		s.AllObjectPropertyDecls[ident.String()] = &declarations.ObjectPropertyDecl{Declaration: declarations.Declaration{IRI: ident.String()}}
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	return
}

func (s *Ontology) parseDifferentIndividuals(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.DifferentIndividuals, parser.B1); err != nil {
		return
	}

	var as []individual.Individual
	as, err = parsefuncs.ParseIndividualsUntilB2(p, s, s)
	if err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}

	s.AllDifferentIndividuals = append(s.AllDifferentIndividuals, axioms.DifferentIndividuals{As: as})

	return
}

func (s *Ontology) parseBracedIRI(p *parser.Parser) (ident *tech.IRI, err error) {
	if err = p.ConsumeTokens(parser.B1); err != nil {
		return
	}

	if ident, err = parsehelper.ParseAndResolveIRI(p, s); err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}

	return
}

func (s *Ontology) parseDataPropertyDomain(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.DataPropertyDomain, parser.B1); err != nil {
		return
	}
	var R meta.DataProperty
	R, err = parsefuncs.ParseDataProperty(p, s, s)
	if err != nil {
		return
	}
	var C meta.ClassExpression
	C, err = parsefuncs.ParseClassExpression(p, s, s)
	if err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	s.AllDataPropertyDomains = append(s.AllDataPropertyDomains, axioms.DataPropertyDomain{R: R, C: C})
	return
}

func (s *Ontology) parseDataPropertyRange(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.DataPropertyRange, parser.B1); err != nil {
		return
	}
	var R meta.DataProperty
	R, err = parsefuncs.ParseDataProperty(p, s, s)
	if err != nil {
		return
	}
	var D meta.DataRange
	D, err = parsefuncs.ParseDataRange(p, s, s)
	if err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	s.AllDataPropertyRanges = append(s.AllDataPropertyRanges, axioms.DataPropertyRange{R: R, D: D})
	return
}

func (s *Ontology) parseDisjointClasses(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.DisjointClasses, parser.B1); err != nil {
		return
	}

	var Cs []meta.ClassExpression
	pos := p.Pos()
	Cs, err = parsefuncs.ParseClassExpressionsUntilB2(p, s, s)
	if err != nil {
		return
	}
	if len(Cs) < 2 { //todo: is there a minimum ?
		err = pos.Errorf("nt enough (%d) in DisjointClasses, expected >=2", len(Cs))
		return
	}
	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	s.AllDisjointClasses = append(s.AllDisjointClasses, axioms.DisjointClasses{Cs})
	return
}

func (s *Ontology) parseEquivalentClasses(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.EquivalentClasses, parser.B1); err != nil {
		return
	}

	var Cs []meta.ClassExpression
	Cs, err = parsefuncs.ParseClassExpressionsUntilB2(p, s, s)
	if err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}

	s.AllEquivalentClasses = append(s.AllEquivalentClasses, axioms.EquivalentClasses{Cs})
	return
}

func (s *Ontology) parseFunctionalDataProperty(p *parser.Parser) (err error) {
	var R meta.DataProperty

	if err = p.ConsumeTokens(parser.FunctionalDataProperty, parser.B1); err != nil {
		return
	}

	if R, err = parsefuncs.ParseDataProperty(p, s, s); err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}

	s.AllFunctionalDataProperties = append(s.AllFunctionalDataProperties, R)
	return
}

func (s *Ontology) parseFunctionalObjectProperty(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.FunctionalObjectProperty); err != nil {
		return
	}
	var P meta.ObjectPropertyExpression
	P, err = s.parseP(p)
	if err != nil {
		return
	}
	s.AllFunctionalObjectProperties = append(s.AllFunctionalObjectProperties, P)
	return
}

func (s *Ontology) parseInverseFunctionalObjectProperty(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.InverseFunctionalObjectProperty); err != nil {
		return
	}
	var P meta.ObjectPropertyExpression
	P, err = s.parseP(p)
	if err != nil {
		return
	}
	s.AllInverseFunctionalObjectProperties = append(s.AllInverseFunctionalObjectProperties, P)
	return
}

func (s *Ontology) parseInverseObjectProperties(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.InverseObjectProperties, parser.B1); err != nil {
		return
	}

	var P1, P2 meta.ObjectPropertyExpression
	if P1, err = parsefuncs.ParseObjectPropertyExpression(p, s, s); err != nil {
		return
	}
	if P2, err = parsefuncs.ParseObjectPropertyExpression(p, s, s); err != nil {
		return
	}

	s.AllInverseObjectProperties = append(s.AllInverseObjectProperties, axioms.InverseObjectProperties{P1, P2})
	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	return
}

func (s *Ontology) parseIrreflexiveObjectProperty(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.IrreflexiveObjectProperty); err != nil {
		return
	}
	var P meta.ObjectPropertyExpression
	P, err = s.parseP(p)
	if err != nil {
		return
	}
	s.AllIrreflexiveObjectProperties = append(s.AllIrreflexiveObjectProperties, P)
	return
}

func (s *Ontology) parseObjectPropertyDomain(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.ObjectPropertyDomain); err != nil {
		return
	}
	var P meta.ObjectPropertyExpression
	var C meta.ClassExpression
	P, C, err = parsefuncs.ParsePC(p, s, s)
	if err != nil {
		return
	}
	s.AllObjectPropertyDomains = append(s.AllObjectPropertyDomains, axioms.ObjectPropertyDomain{P: P, C: C})
	return
}

func (s *Ontology) parseObjectPropertyRange(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.ObjectPropertyRange); err != nil {
		return
	}
	var P meta.ObjectPropertyExpression
	var C meta.ClassExpression
	P, C, err = parsefuncs.ParsePC(p, s, s)
	if err != nil {
		return
	}
	s.AllObjectPropertyRanges = append(s.AllObjectPropertyRanges, axioms.ObjectPropertyRange{P: P, C: C})
	return
}

func (s *Ontology) parseReflexiveObjectProperty(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.ReflexiveObjectProperty); err != nil {
		return
	}
	var P meta.ObjectPropertyExpression
	P, err = s.parseP(p)
	if err != nil {
		return
	}
	s.AllReflexiveObjectProperties = append(s.AllReflexiveObjectProperties, P)
	return
}

func (s *Ontology) parseSubClassOf(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.SubClassOf, parser.B1); err != nil {
		return
	}

	var Cs []meta.ClassExpression
	pos := p.Pos()
	Cs, err = parsefuncs.ParseClassExpressionsUntilB2(p, s, s)
	if err != nil {
		return
	}
	if len(Cs) != 2 {
		err = pos.Errorf("wrong param count (%d) in SubClassOf, expected 2", len(Cs))
		return
	}
	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	s.AllSubClassOfs = append(s.AllSubClassOfs, axioms.SubClassOf{C1: Cs[0], C2: Cs[1]})
	return
}

func (s *Ontology) parseSubDataPropertyOf(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.SubDataPropertyOf, parser.B1); err != nil {
		return
	}

	var P1, P2 meta.DataProperty
	if P1, err = parsefuncs.ParseDataProperty(p, s, s); err != nil {
		return
	}
	if P2, err = parsefuncs.ParseDataProperty(p, s, s); err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	s.AllSubDataPropertyOfs = append(s.AllSubDataPropertyOfs, axioms.SubDataPropertyOf{P1, P2})

	return
}

func (s *Ontology) parseSubObjectPropertyOf(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.SubObjectPropertyOf, parser.B1); err != nil {
		return
	}

	var P1, P2 meta.ObjectPropertyExpression
	if P1, err = parsefuncs.ParseObjectPropertyExpression(p, s, s); err != nil {
		return
	}
	if P2, err = parsefuncs.ParseObjectPropertyExpression(p, s, s); err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	s.AllSubObjectPropertyOfs = append(s.AllSubObjectPropertyOfs, axioms.SubObjectPropertyOf{P1, P2})

	return
}

func (s *Ontology) parseSymmetricObjectProperty(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.SymmetricObjectProperty); err != nil {
		return
	}
	var P meta.ObjectPropertyExpression
	P, err = s.parseP(p)
	if err != nil {
		return
	}
	s.AllSymmetricObjectProperties = append(s.AllSymmetricObjectProperties, P)
	return
}

func (s *Ontology) parseTransitiveObjectProperty(p *parser.Parser) (err error) {
	if err = p.ConsumeTokens(parser.TransitiveObjectProperty); err != nil {
		return
	}
	var P meta.ObjectPropertyExpression
	P, err = s.parseP(p)
	if err != nil {
		return
	}
	s.AllTransitiveObjectProperties = append(s.AllTransitiveObjectProperties, P)
	return
}

func (s *Ontology) parseP(p *parser.Parser) (P meta.ObjectPropertyExpression, err error) {
	if err = p.ConsumeTokens(parser.B1); err != nil {
		return
	}

	if P, err = parsefuncs.ParseObjectPropertyExpression(p, s, s); err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	return
}

func (s *Ontology) ClassDeclExists(ident string) bool {
	_, ok := s.AllClassDecls[ident]
	return ok
}

func (s *Ontology) DataPropertyDeclExists(ident string) bool {
	_, ok := s.AllDataPropertyDecls[ident]
	return ok
}

func (s *Ontology) NamedIndividualDeclExists(ident string) bool {
	_, ok := s.AllNamedIndividualDecls[ident]
	return ok
}

func (s *Ontology) ObjectPropertyDeclExists(ident string) bool {
	_, ok := s.AllObjectPropertyDecls[ident]
	return ok
}

func (s *Ontology) ResolvePrefix(prefix string) (res string, ok bool) {
	res, ok = s.Prefixes[prefix]
	return
}

func (s *Ontology) About() string {
	return fmt.Sprintf("%v with %d annotations, %d classes, %d object properties, %d data properties, %d named individuals, %d datatypes.",
		s.IRI,
		len(s.AllAnnotationPropertyDecls),
		len(s.AllClassDecls),
		len(s.AllObjectPropertyDecls),
		len(s.AllDataPropertyDecls),
		len(s.AllNamedIndividualDecls),
		len(s.AllDatatypeDecls),
	)
}
