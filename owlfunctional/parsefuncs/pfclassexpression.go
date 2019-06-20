package parsefuncs

import (
	"github.com/shful/gofp/owlfunctional/builtindatatypes"
	"github.com/shful/gofp/owlfunctional/classexpression"
	"github.com/shful/gofp/owlfunctional/individual"
	"github.com/shful/gofp/owlfunctional/literal"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/owlfunctional/parser"
	"github.com/shful/gofp/parsehelper"
	"github.com/shful/gofp/store"
	"github.com/shful/gofp/tech"
)

func ParseClassExpression(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	tok, lit, pos := p.ScanIgnoreWSAndComment()
	p.Unscan()
	switch tok {
	// Boolean Conectives and Enumeration of Individuals
	case parser.ObjectComplementOf:
		expr, err = parseObjectComplementOf(p, decls, prefixes)
	case parser.ObjectIntersectionOf:
		expr, err = parseObjectIntersectionOf(p, decls, prefixes)
	case parser.ObjectOneOf:
		expr, err = parseObjectOneOf(p, decls, prefixes)
	case parser.ObjectUnionOf:
		expr, err = parseObjectUnionOf(p, decls, prefixes)
	// Object Property Restrictions
	case parser.ObjectAllValuesFrom:
		expr, err = parseObjectAllValuesFrom(p, decls, prefixes)
	case parser.ObjectExactCardinality:
		expr, err = parseObjectExactCardinality(p, decls, prefixes)
	case parser.ObjectHasValue:
		expr, err = parseObjectHasValue(p, decls, prefixes)
	case parser.ObjectHasSelf:
		expr, err = parseObjectHasSelf(p, decls, prefixes)
	case parser.ObjectMaxCardinality:
		expr, err = parseObjectMaxCardinality(p, decls, prefixes)
	case parser.ObjectMinCardinality:
		expr, err = parseObjectMinCardinality(p, decls, prefixes)
	case parser.ObjectSomeValuesFrom:
		expr, err = parseObjectSomeValuesFrom(p, decls, prefixes)
	// Data Property Restrictions
	case parser.DataAllValuesFrom:
		expr, err = parseDataAllValuesFrom(p, decls, prefixes)
	case parser.DataSomeValuesFrom:
		expr, err = parseDataSomeValuesFrom(p, decls, prefixes)
	case parser.DataHasValue:
		expr, err = parseDataHasValue(p, decls, prefixes)
	case parser.DataExactCardinality:
		expr, err = parseDataExactCardinality(p, decls, prefixes)
	case parser.DataMaxCardinality:
		expr, err = parseDataMaxCardinality(p, decls, prefixes)
	case parser.DataMinCardinality:
		expr, err = parseDataMinCardinality(p, decls, prefixes)
	case parser.IRI, parser.COLON, parser.IDENT:
		// must be simply CN
		var ident *tech.IRI
		ident, err = parsehelper.ParseAndResolveIRI(p, prefixes)

		if err != nil {
			err = pos.ErrorfUnexpectedToken(tok, lit, "IRI as Class Expression found but parse failed:"+err.Error())
			return
		}

		if builtindatatypes.IsOWL(*ident) {
			//must be one of the predefined OWL classes
			switch ident.Fragment {
			case "Thing":
				expr = &classexpression.OWLThing{}
			case "Nothing":
				expr = &classexpression.OWLNothing{}
			default:
				err = pos.Errorf(`unexpected OWL name "%v"`, ident.Fragment)
			}
			return
		} else {

		}

		var ok bool
		expr, ok = decls.ClassDecl(ident.String())
		if !ok {
			err = pos.Errorf("Unknown ref to %v. Expected class expression.", ident)
		}
	default:
		err = pos.Errorf("Expected class expression (found:%v which seems something different)", lit)
	}

	return
}

func parseDataAllValuesFrom(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.B1); err != nil {
		return
	}

	var R meta.DataProperty
	var D meta.DataRange
	R, D, err = ParseRD(p, decls, prefixes)
	if err != nil {
		return
	}
	expr = &classexpression.DataAllValuesFrom{R: R, D: D}
	return
}

func parseDataExactCardinality(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.DataExactCardinality, parser.B1); err != nil {
		return
	}
	var R meta.DataProperty
	var D meta.DataRange
	var n int
	var isQualified bool
	n, R, D, isQualified, err = parseNRD(p, decls, prefixes)
	if err != nil {
		return
	}
	if isQualified {
		expr = &classexpression.DataQualifiedExactCardinality{N: n, R: R, D: D}
	} else {
		expr = &classexpression.DataExactCardinality{N: n, R: R}

	}
	return
}

func parseDataMaxCardinality(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.DataMaxCardinality, parser.B1); err != nil {
		return
	}
	var R meta.DataProperty
	var D meta.DataRange
	var n int
	var isQualified bool
	n, R, D, isQualified, err = parseNRD(p, decls, prefixes)
	if err != nil {
		return
	}
	if isQualified {
		expr = &classexpression.DataQualifiedMaxCardinality{N: n, R: R, D: D}
	} else {
		expr = &classexpression.DataMaxCardinality{N: n, R: R}

	}
	return
}

func parseDataMinCardinality(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.DataMinCardinality, parser.B1); err != nil {
		return
	}

	var R meta.DataProperty
	var D meta.DataRange
	var n int
	var isQualified bool
	n, R, D, isQualified, err = parseNRD(p, decls, prefixes)
	if err != nil {
		return
	}
	if isQualified {
		expr = &classexpression.DataQualifiedMinCardinality{N: n, R: R, D: D}
	} else {
		expr = &classexpression.DataMinCardinality{N: n, R: R}

	}
	return
}

func parseDataHasValue(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	var R meta.DataProperty
	var v literal.OWLLiteral
	pos := p.Pos()
	if err = p.ConsumeTokens(parser.DataHasValue, parser.B1); err != nil {
		err = pos.EnrichErrorMsg(err, "parsing DataHasvalue")
		return
	}

	R, err = ParseDataProperty(p, decls, prefixes)
	if err != nil {
		err = pos.EnrichErrorMsg(err, "parsing R in DataHasvalue")
		return
	}
	v, err = ParseOWLLiteral(p, prefixes)
	if err != nil {
		err = pos.EnrichErrorMsg(err, "parsing v in DataHasvalue")
		return
	}
	if err = p.ConsumeTokens(parser.B2); err != nil {
		err = pos.EnrichErrorMsg(err, "parsing DataHasvalue")
		return
	}
	expr = &classexpression.DataHasValue{R: R, V: v}
	return
}

func parseDataSomeValuesFrom(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.DataSomeValuesFrom, parser.B1); err != nil {
		return
	}

	var R meta.DataProperty
	var D meta.DataRange
	R, D, err = ParseRD(p, decls, prefixes)
	if err != nil {
		return
	}
	expr = &classexpression.DataSomeValuesFrom{R: R, D: D}
	return
}

func parseObjectAllValuesFrom(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.ObjectAllValuesFrom, parser.B1); err != nil {
		return
	}
	var P meta.ObjectPropertyExpression
	var C meta.ClassExpression
	P, C, err = ParsePC(p, decls, prefixes)
	if err != nil {
		return
	}
	expr = &classexpression.ObjectAllValuesFrom{P: P, C: C}
	return
}

func parseObjectSomeValuesFrom(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.ObjectSomeValuesFrom, parser.B1); err != nil {
		return
	}
	var P meta.ObjectPropertyExpression
	var C meta.ClassExpression
	P, C, err = ParsePC(p, decls, prefixes)
	if err != nil {
		return
	}
	expr = &classexpression.ObjectSomeValuesFrom{P: P, C: C}
	return
}

func parseObjectExactCardinality(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.ObjectExactCardinality, parser.B1); err != nil {
		return
	}
	n, P, C, isQualified, err := parseNPC(p, decls, prefixes)
	if err != nil {
		return
	}
	if isQualified {
		expr = &classexpression.ObjectQualifiedExactCardinality{N: n, P: P, C: C}
	} else {
		expr = &classexpression.ObjectExactCardinality{N: n, P: P}
	}

	return
}

func parseObjectHasValue(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.ObjectHasValue, parser.B1); err != nil {
		return
	}

	var P meta.ObjectPropertyExpression
	var a individual.Individual
	P, a, err = ParsePa(p, decls, prefixes)
	if err != nil {
		return
	}
	expr = &classexpression.ObjectHasValue{P: P, A: a}

	return
}

func parseObjectHasSelf(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.ObjectHasSelf, parser.B1); err != nil {
		return
	}
	var P meta.ObjectPropertyExpression

	P, err = ParseObjectPropertyExpression(p, decls, prefixes)
	if err != nil {
		return
	}

	expr = &classexpression.ObjectHasSelf{P: P}
	return
}

func parseObjectMaxCardinality(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.ObjectMaxCardinality, parser.B1); err != nil {
		return
	}
	n, P, C, qualified, err := parseNPC(p, decls, prefixes)
	if err != nil {
		return
	}
	if qualified {
		expr = &classexpression.ObjectQualifiedMaxCardinality{N: n, P: P, C: C}
	} else {
		expr = &classexpression.ObjectMaxCardinality{N: n, P: P}
	}
	return
}

func parseObjectMinCardinality(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.ObjectMinCardinality, parser.B1); err != nil {
		return
	}

	n, P, C, qualified, err := parseNPC(p, decls, prefixes)
	if err != nil {
		return
	}
	if qualified {
		expr = &classexpression.ObjectQualifiedMinCardinality{N: n, P: P, C: C}
	} else {
		expr = &classexpression.ObjectMinCardinality{N: n, P: P}
	}
	return
}

func parseObjectComplementOf(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.ObjectComplementOf, parser.B1); err != nil {
		return
	}

	var Cs []meta.ClassExpression
	pos := p.Pos()
	Cs, err = ParseClassExpressionsUntilB2(p, decls, prefixes)
	if err != nil {
		return
	}
	if len(Cs) != 1 {
		err = pos.Errorf("wrong param count (%d) in ObjectComplementOf, expected 1", len(Cs))
		return
	}
	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	expr = &classexpression.ObjectComplementOf{C: Cs[0]}
	return
}

func parseObjectIntersectionOf(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.ObjectIntersectionOf, parser.B1); err != nil {
		return
	}

	var Cs []meta.ClassExpression
	pos := p.Pos()
	Cs, err = ParseClassExpressionsUntilB2(p, decls, prefixes)
	if err != nil {
		return
	}
	if len(Cs) < 2 { //todo allow 1 or even 0==Nothing?
		err = pos.Errorf("not enough params (%d) in ObjectIntersectionOf", len(Cs))
		return
	}
	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	expr = &classexpression.ObjectIntersectionOf{Cs}
	return
}

func parseObjectOneOf(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.ObjectOneOf, parser.B1); err != nil {
		return
	}

	var as []individual.Individual
	as, err = ParseIndividualsUntilB2(p, decls, prefixes)
	if err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	expr = &classexpression.ObjectOneOf{As: as}

	return
}

func parseObjectUnionOf(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.ClassExpression, err error) {
	if err = p.ConsumeTokens(parser.ObjectUnionOf, parser.B1); err != nil {
		return
	}

	var Cs []meta.ClassExpression
	pos := p.Pos()
	Cs, err = ParseClassExpressionsUntilB2(p, decls, prefixes)
	if err != nil {
		return
	}
	if len(Cs) < 2 {
		err = pos.Errorf("not enough params (%d) in ObjectUnionOf", len(Cs))
		return
	}
	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	expr = &classexpression.ObjectUnionOf{Cs}
	return
}

// ParseClassExpressionsUntilB2 parses all ClassExpression until ")" is found
// The closing ")" is not consumed.
func ParseClassExpressionsUntilB2(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (Cs []meta.ClassExpression, err error) {

	var tok parser.Token
	var C meta.ClassExpression

	for {
		tok, _, _ = p.ScanIgnoreWSAndComment()
		p.Unscan()
		if tok == parser.B2 {
			break
		}

		C, err = ParseClassExpression(p, decls, prefixes)
		if err != nil {
			return
		}
		Cs = append(Cs, C)
	}

	return
}
