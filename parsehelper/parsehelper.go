package parsehelper

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/shful/gofp/owlfunctional/parser"
	"github.com/shful/gofp/tech"
)

func ParseAndResolveIRI(p *parser.Parser, prefixes tech.Prefixes) (ident *tech.IRI, err error) {
	var head, name string
	pos := p.Pos()

	tok, lit, pos := p.ScanIgnoreWSAndComment()
	p.Unscan()
	switch tok {
	case parser.IRI:
		// IRI means, no prefix to resolve:
		head, name, err = ParseIRIWithFragment(p)
		if err != nil {
			return
		}
		ident, err = tech.NewIRIFromString(head + name)
	case parser.IDENT:
		fallthrough
	case parser.COLON:
		// IDENT and COLON both require resolving a prefix:
		var prefix string
		prefix, name, err = ParsePrefixedName(p)
		if err != nil {
			return
		}

		var ok bool
		head, ok = prefixes.ResolvePrefix(prefix)
		if !ok {
			err = pos.Errorf("unknown prefix %v", prefix)
			return
		}

		ident, err = tech.NewIRIFromString(head + name)
		if err != nil {
			err = pos.Errorf("prefixed name (%v:%v) resolved to invalid IRI (%v)", prefix, name, head+name)
			return
		}
	default:
		err = pos.Errorf("unexpected \"%v\" - need IRI, or prefixed name, or _ for anonymous individual.", lit)
	}

	return
}

func ParsePrefixedName(p *parser.Parser) (prefix, name string, err error) {

	tok, lit, pos := p.ScanIgnoreWSAndComment()
	if tok == parser.IDENT {
		// prefix:classname
		lit += prolongIDENT(p)

		prefix = lit
		if err = p.ConsumeTokens(parser.COLON); err != nil {
			return
		}
	} else if tok == parser.COLON {
		// :classname
		prefix = ""
	} else {
		err = pos.Errorf("unexpected \"%v\" - need prefixed name", lit)
		return
	}

	tok, name, pos = p.ScanIgnoreWSAndComment()
	if tok != parser.IDENT {
		err = pos.Errorf("unexpected \"%v\" - need identifier in prefixed name", lit)
	}
	name += prolongIDENT(p)

	return
}

// prolongIDENT continues to parse a parser.IDENT which eventuall was parsed partially only,
// and returns the remaining which needs to be added to the already parsed IDENT.
// The parser must be directly behind the IDENT to eventually prolong.
// That is needed when an IDENT contains a character which is a token for itself. For example, the hyphen "-"
// is allowed inside an IDENT and must be added, along with eventually following IDENT or hyphen tokens.
func prolongIDENT(p *parser.Parser) (suffix string) {
	for {
		tok, lit, _ := p.Scan()
		switch tok {
		case parser.MINUS:
			suffix += lit
		case parser.IDENT:
			suffix += lit
		case parser.INTLIT:
			suffix += lit
		default:
			p.Unscan()
			return
		}
	}
}

// ParseUnprefixedIRI parses an IRI which is not shortened with a prefix. Instead, it must look like "<.*>"
func ParseUnprefixedIRI(p *parser.Parser) (iri string, err error) {
	pos := p.Pos()
	tok, lit, pos := p.ScanIgnoreWSAndComment()
	if tok == parser.IRI {
		if !(strings.HasPrefix(lit, "<") && strings.HasSuffix(lit, ">")) {
			err = pos.Errorf("expected IRI, but missing < and > on the ends (found:%v)", lit)
		} else {
			iri = lit[1 : len(lit)-1]
		}
	} else {
		err = pos.Errorf("expected IRI, but found:%v", parser.DescribeToklit(tok, lit))
	}
	return
}

// ParseIRIWithFragment parses an IRI which must be surrounded with "<" ">". The surrounding <> are not returned.
// If there's a fragment, head is everything until and including the #, and fragment is the remaining.
// With no fragment, the head is the full IRI content and fragment is empty.
// The IRI must not be empty. "<>" results in an error.
func ParseIRIWithFragment(p *parser.Parser) (head, fragment string, err error) {
	pos := p.Pos()
	tok, iri, pos := p.ScanIgnoreWSAndComment()
	if tok == parser.IRI {
		if !(strings.HasPrefix(iri, "<") && strings.HasSuffix(iri, ">")) {
			err = pos.Errorf("expected IRI, but missing < and > on the ends (found:%v)", iri)
			return
		}
		if len(iri) == 2 {
			err = pos.Errorf("empty IRI between <>")
			return
		}
		var u *url.URL
		u, err = url.Parse(iri[1 : len(iri)-1])
		if err != nil {
			return
		}
		fragment = u.Fragment
		head = iri[1 : len(iri)-1-len(fragment)] // everything until, and including, the fragments "#"
	} else {
		err = pos.Errorf("expected IRI, but found:%v", parser.DescribeToklit(tok, iri))
	}
	return
}

//todo support number formats, especially int with sign
func ParseNonNegativeInteger(p *parser.Parser) (res int, err error) {
	tok, lit, pos := p.ScanIgnoreWSAndComment()
	if tok != parser.INTLIT {
		err = pos.Errorf("int literal needed, found %v", lit)
		return
	}
	res, err = strconv.Atoi(lit)
	if res < 0 {
		err = pos.Errorf("nonnegative integer needed, found %v", lit)
		return
	}
	return
}
