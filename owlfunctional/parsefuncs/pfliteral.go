package parsefuncs

import (
	"fmt"
	"strconv"

	"github.com/shful/gofp/owlfunctional/builtindatatypes"
	"github.com/shful/gofp/owlfunctional/literal"
	"github.com/shful/gofp/owlfunctional/parser"
	"github.com/shful/gofp/parsehelper"
	"github.com/shful/gofp/tech"
)

func ParseOWLLiteral(p *parser.Parser, prefixes tech.Prefixes) (l literal.OWLLiteral, err error) {
	var tok parser.Token
	var lit string
	var pos parser.ParserPosition
	var langtag string
	var datatypeIRI *tech.IRI

	tok, lit, pos = p.ScanIgnoreWSAndComment()

	switch tok {
	case parser.OWLTrue, parser.OWLFalse:
		langtag = ""
		datatypeIRI = tech.MustNewFragmentedIRI(builtindatatypes.PRE_XSD, "boolean")
	case parser.STRINGLIT:
		langtag, err = parseSuffixLangtag(p)
		if err != nil {
			return
		}
		fallthrough
	case parser.INTLIT, parser.FLOATLIT:
		datatypeIRI, err = parseSuffixLiteraltype(p, prefixes)
		if err != nil {
			return
		}
		if datatypeIRI == nil { // literal had no ^^
			switch tok {
			case parser.INTLIT:
				datatypeIRI = tech.MustNewFragmentedIRI(builtindatatypes.PRE_XSD, "integer")
			case parser.FLOATLIT:
				datatypeIRI = tech.MustNewFragmentedIRI(builtindatatypes.PRE_XSD, "decimal")
			case parser.STRINGLIT:
				datatypeIRI = tech.MustNewFragmentedIRI(builtindatatypes.PRE_XSD, "string")
			}
		} else { // explicit literal type given with ^^
			// numbers can be quoted like "123" or "0.01".
			// The lexer syntactically decides for string token.
			// Correct token type if explicit number type is given, and value fits:
			if tok == parser.STRINGLIT {
				if mustTok, ok := builtindatatypes.BuiltinDatatypes[datatypeIRI.String()]; ok {
					switch mustTok {
					case parser.INTLIT:
						if _, err = strconv.Atoi(lit); err == nil {
							tok = mustTok
						}
					case parser.FLOATLIT:
						if _, err = strconv.ParseFloat(lit, 64); err == nil {
							tok = mustTok
						}
					}
				}
			}
			err = literaltypeMismatch(tok, datatypeIRI.String())
		}
		if err != nil {
			err = pos.EnrichErrorMsg(err, "parsing literal")
			return
		}
	default:
		err = pos.Errorf("unexpected %v when parsing literal", parser.DescribeToklit(tok, lit))
		return
	}

	l = literal.OWLLiteral{Value: lit, LangTag: langtag, Literaltype: datatypeIRI.String()}
	return
}

// parseSuffixLangtag returns "en", if "@en" is found.
// Empty string if not @... is found. Error if @<syntactically-invalid-langtag> is found.
func parseSuffixLangtag(p *parser.Parser) (langtag string, err error) {
	var tok parser.Token

	tok, _, _ = p.ScanIgnoreWSAndComment()
	if tok != parser.AT {
		p.Unscan()
		return
	}
	var pos parser.ParserPosition
	tok, langtag, pos = p.ScanIgnoreWSAndComment()
	if tok != parser.IDENT {
		p.Unscan()
		err = pos.Errorf("expected langtag, not %v", parser.DescribeToklit(tok, langtag))
	}
	return
}

// parseSuffixLiteraltype expects a prefixed name starting with ^^, e.g. "^^xsd:integer"
// nil if not ^^... is found. Error if ^^<syntactically-invalid-literaltype> is found.
func parseSuffixLiteraltype(p *parser.Parser, prefixes tech.Prefixes) (ident *tech.IRI, err error) {
	var tok parser.Token
	var pos parser.ParserPosition

	tok, _, pos = p.ScanIgnoreWSAndComment()
	if tok != parser.DOUBLECIRCUM {
		p.Unscan()
		return
	}

	var prefix, name string
	prefix, name, err = parsehelper.ParsePrefixedName(p)
	if err != nil {
		p.Unscan()
		return
	}

	var ok bool
	var resolved string
	resolved, ok = prefixes.ResolvePrefix(prefix)
	if !ok {
		err = pos.Errorf("unknown prefix (%v) in literal type", prefix)
		return
	}

	ident, err = tech.NewIRIFromString(resolved + name)

	if err != nil {
		err = pos.Errorf("prefixed name (%v:%v) resolved to invalid IRI (%v)", prefix, name, resolved+name)
		return
	}
	return
}

// literaltypeMismatch returns an error if, for example, an INTLIT token comes with ^^xsd:string suffix.
func literaltypeMismatch(tok parser.Token, literaltype string) error {
	var mustTok parser.Token
	var ok bool

	if mustTok, ok = builtindatatypes.BuiltinDatatypes[literaltype]; ok {
		if tok != mustTok {
			return fmt.Errorf("literal type mismatch with value (%v)", literaltype)
		}
	}
	// no mismatch check for custom literaltype
	return nil
}
