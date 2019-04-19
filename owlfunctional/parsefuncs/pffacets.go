package parsefuncs

import (
	"fmt"

	"github.com/shful/gofp/owlfunctional/builtindatatypes"
	"github.com/shful/gofp/owlfunctional/facets"
	"github.com/shful/gofp/owlfunctional/literal"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/owlfunctional/parser"
	"github.com/shful/gofp/parsehelper"
	"github.com/shful/gofp/store"
	"github.com/shful/gofp/tech"
)

func ParseDataRange(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.DataRange, err error) {
	tok, _, _ := p.ScanIgnoreWSAndComment()

	p.Unscan()

	switch tok {
	case parser.DataComplementOf:
		panic("not implemented")
	case parser.DataIntersectionOf:
		panic("not implemented")
	case parser.DataUnionOf:
		panic("not implemented")
	case parser.DataOneOf:
		panic("not implemented")
	case parser.DatatypeRestriction:
		parseDatatypeRestriction(p, decls, prefixes)
	default:
		// must be literal, i.e. named datatype (DN)
		expr, err = parseNamedDatatype(p, decls, prefixes)
	}
	return
}

func parseDatatypeRestriction(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr facets.DatatypeRestriction, err error) {
	if err = p.ConsumeTokens(parser.DatatypeRestriction, parser.B1); err != nil {
		return
	}

	// The DN
	var DN meta.NamedDatatype
	DN, err = parseNamedDatatype(p, decls, prefixes)
	if err != nil {
		return
	}

	// The f,v - pairs
	var fvPairs []*facets.FVPair
	fvPairs, err = parseFVPairsUntilB2(p, prefixes)
	if err != nil {
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}

	expr = facets.DatatypeRestriction{DN: DN, FVPairs: fvPairs}
	return
}

// parseFVPairsUntilB2 consumes unlimited (f,v) until, but excluding, a closing B2 )
func parseFVPairsUntilB2(p *parser.Parser, prefixes tech.Prefixes) (fvPairs []*facets.FVPair, err error) {
	var fvPair *facets.FVPair
	for i := 0; ; i++ {
		pos := p.Pos()
		fvPair, err = parseFVPairUntilB2(p, prefixes)
		if err != nil {
			err = pos.EnrichErrorMsg(err, fmt.Sprintf("parsing %v (facet,literal) pair", tech.ZeroBasedPosWord(i)))
			return
		}
		if fvPair == nil {
			// closing brace reached
			return
		}
		fvPairs = append(fvPairs, fvPair)
	}
	return
}

func parseFacet(p *parser.Parser) (facet facets.Facet, err error) {
	var prefix, name string
	pos := p.Pos()
	prefix, name, err = parsehelper.ParsePrefixedName(p)
	if err != nil {
		err = pos.EnrichErrorMsg(err, "parsing facet:%v")
		return
	}
	switch x := parser.FmtPrefixedName(prefix, name); x {
	case "xsd:maxExclusive":
		facet = facets.Xsd_maxExclusive
	case "xsd:maxInclusive":
		facet = facets.Xsd_maxInclusive
	case "xsd:minExclusive":
		facet = facets.Xsd_minExclusive
	case "xsd:minInclusive":
		facet = facets.Xsd_minInclusive
	default:
		err = pos.Errorf("expected known facet, found %v.", x)
	}
	return
}

// parseFVPairUntilB2 returns nil when closing brace ) is reached. The brace is not consumed.
func parseFVPairUntilB2(p *parser.Parser, prefixes tech.Prefixes) (fvPair *facets.FVPair, err error) {
	var f facets.Facet
	tok, _, _ := p.ScanIgnoreWSAndComment()
	p.Unscan()

	if tok == parser.B2 {
		return
	}
	f, err = parseFacet(p)
	if err != nil {
		return
	}
	var v literal.OWLLiteral
	v, err = ParseOWLLiteral(p, prefixes)
	if err != nil {
		return
	}
	fvPair = &facets.FVPair{F: f, V: v}
	return
}

// Datatypes are entities that refer to sets of data values.
// Thus, datatypes are analogous to classes, the main difference being that
// the former contain data values such as strings and numbers, rather than individuals.
// Datatypes are a kind of data range, which allows them to be used in restrictions.
// As explained in Section 7, each data range is associated with an arity; for datatypes, the arity is always one.
// The built-in datatype rdfs:Literal denotes any set of data values that contains the union of the value spaces of all datatypes.
func parseNamedDatatype(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr meta.NamedDatatype, err error) {
	var ident *tech.IRI

	pos := p.Pos()
	ident, err = parsehelper.ParseAndResolveIRI(p, prefixes)
	if err != nil {
		err = pos.EnrichErrorMsg(err, "parsing named datatype")
		return
	}

	// Builtin Datatype IRI is allowed:
	if builtindatatypes.BuiltinDatatypeExists(ident.String()) {
		expr = &facets.BuiltinDatatype{facets.NamedDatatypeImpl{DatatypeIRI: ident.String()}}
		return
	}

	// Declared Datatype IRI is allowed:
	if _, ok := decls.GetDatatypeDecl(*ident); ok {
		expr = &facets.CustomNamedDatatype{facets.NamedDatatypeImpl{DatatypeIRI: ident.String()}}
		return
	}

	err = pos.Errorf("unknown datatype literal (%v)", ident)
	return
}
