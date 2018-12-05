package parsefuncs

import (
	"reifenberg.de/gofp/owlfunctional/individual"
	"reifenberg.de/gofp/owlfunctional/literal"
	"reifenberg.de/gofp/owlfunctional/meta"
	"reifenberg.de/gofp/owlfunctional/parser"
	"reifenberg.de/gofp/parsehelper"
	"reifenberg.de/gofp/tech"
)

// parseNPC parses the triple (n,P,[C]) and consumes both braces.
// C is optional. If found, isQualified is true.
func parseNPC(p *parser.Parser, decls tech.Declarations, prefixes tech.Prefixes) (n int, P meta.ObjectPropertyExpression, C meta.ClassExpression, isQualified bool, err error) {
	if err = p.ConsumeTokens(parser.B1); err != nil {
		return
	}

	n, err = parsehelper.ParseNonNegativeInteger(p)
	if err != nil {
		return
	}

	P, err = ParseObjectPropertyExpression(p, decls, prefixes)
	if err != nil {
		return
	}

	// expect ) or C)
	tok, _, _ := p.ScanIgnoreWSAndComment()
	p.Unscan()
	switch tok {
	case parser.B2:
		// unqualified
	default:
		C, err = ParseClassExpression(p, decls, prefixes)
		if err != nil {
			return
		}
	}
	err = p.ConsumeTokens(parser.B2)

	return
}

// ParseNRD parses the triple (n,R,[D]) and consumes both braces.
// D is optional. If found, isQualified is true.
func parseNRD(p *parser.Parser, decls tech.Declarations, prefixes tech.Prefixes) (n int, R meta.DataProperty, D meta.DataRange, isQualified bool, err error) {
	if err = p.ConsumeTokens(parser.B1); err != nil {
		return
	}

	n, err = parsehelper.ParseNonNegativeInteger(p)
	if err != nil {
		return
	}

	R, err = ParseDataProperty(p, decls, prefixes)
	if err != nil {
		return
	}

	// expect ) or D)
	tok, _, pos := p.ScanIgnoreWSAndComment()
	p.Unscan()
	switch tok {
	case parser.B2:
		// unqualified
	default:
		D, err = ParseDataRange(p, decls, prefixes)
		if err != nil {
			err = pos.Errorf("parsing D in DataExactCardinality:%v", err)
			return
		}
		isQualified = true
	}
	err = p.ConsumeTokens(parser.B2)

	return
}

// ParseRD parses the pair (R,D) and consumes both braces.
func ParseRD(p *parser.Parser, decls tech.Declarations, prefixes tech.Prefixes) (R meta.DataProperty, D meta.DataRange, err error) {
	if err = p.ConsumeTokens(parser.B1); err != nil {
		return
	}

	R, err = ParseDataProperty(p, decls, prefixes)
	if err != nil {
		return
	}

	D, err = ParseDataRange(p, decls, prefixes)
	if err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}

	return
}

// ParsePC parses the pair (P,C) and consumes both braces.
func ParsePC(p *parser.Parser, decls tech.Declarations, prefixes tech.Prefixes) (P meta.ObjectPropertyExpression, C meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.B1); err != nil {
		return
	}

	P, err = ParseObjectPropertyExpression(p, decls, prefixes)
	if err != nil {
		return
	}

	C, err = ParseClassExpression(p, decls, prefixes)
	if err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}

	return
}

// ParsePa parses the pair (P,a) and consumes both braces.
func ParsePa(p *parser.Parser, decls tech.Declarations, prefixes tech.Prefixes) (P meta.ObjectPropertyExpression, A individual.Individual, err error) {
	if err = p.ConsumeTokens(parser.B1); err != nil {
		return
	}

	P, err = ParseObjectPropertyExpression(p, decls, prefixes)
	if err != nil {
		return
	}

	A, err = ParseIndividual(p, decls, prefixes)
	if err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}

	return
}

type ARGTYPE int

const (
	ArgtypeAnonymousIndividual ARGTYPE = iota
	ArgtypeIRI
	ArgtypeLiteral
)

// Parset reads IRI or literal or anonymous individual, which is shortened as "t" in the OWL spec
func Parset(p *parser.Parser, decls tech.Declarations, prefixes tech.Prefixes) (expr string, argtype ARGTYPE, err error) {
	var tok parser.Token
	tok, _, _ = p.ScanIgnoreWSAndComment()
	p.Unscan()

	if literal.MaybeOWLLiteral(tok) {
		var l literal.OWLLiteral
		l, err = ParseOWLLiteral(p, prefixes)
		if err == nil {
			expr = l.LiteralString()
			argtype = ArgtypeLiteral
		}
		return
	}

	expr, err = parsehelper.ParseUnprefixedIRI(p)
	if err == nil {
		argtype = ArgtypeIRI
		return
	}
	p.Unscan()

	// read prefixed name at last, because unparsing prefixed names is not possible
	// since the parser currently can unparse one token only.
	pos := p.Pos()
	var tech, IRI ident
	ident, err = parsehelper.ParseAndResolveIdentifier(p, prefixes)
	if err == nil {
		expr = IRI.String()
		if s1 != "_" { hier weiter: was ist mit AnonymousIndividual, wenn wir jetzt ParseAndResolveIdentifier nehmen?
			mache Übersicht, wo wird ParsePrefixedName genutzt, und wo sind welche Typen möglich:
			IRI, :name, prefix:name, name, _  mache passende Funktionen in Parsehelper.
			Notiere auch überall, welcher Typ im Aufrufer gebraucht wird: hier ist es expr string plus Typinfo.
			argtype = ArgtypeAnonymousIndividual
		} else {
			argtype = ArgtypeIRI
		}
		return
	}

	pos.EnrichErrorMsg(err, "expected IRI,anonymous individual or literal")
	return
}
