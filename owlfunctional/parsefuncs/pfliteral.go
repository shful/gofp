package parsefuncs

import (
	"fmt"
	"strconv"

	"reifenberg.de/gofp/owlfunctional/builtindatatypes"
	"reifenberg.de/gofp/owlfunctional/literal"
	"reifenberg.de/gofp/owlfunctional/parser"
	"reifenberg.de/gofp/parsehelper"
	"reifenberg.de/gofp/tech"
)

func ParseOWLLiteral(p *parser.Parser, prefixes tech.Prefixes) (l literal.OWLLiteral, err error) {
	var tok parser.Token
	var lit string
	var pos parser.ParserPosition
	var langtag, literaltype string

	tok, lit, pos = p.ScanIgnoreWSAndComment()

	switch tok {
	case parser.OWLTrue, parser.OWLFalse:
		l = literal.OWLLiteral{Value: lit, LangTag: "", Literaltype: "xsd:boolean"}
		return
	case parser.STRINGLIT:
		langtag, err = parseSuffixLangtag(p)
		if err != nil {
			return
		}
		fallthrough
	case parser.INTLIT, parser.FLOATLIT:
		literaltype, err = parseSuffixLiteraltype(p, prefixes)
		if err != nil {
			return
		}
		if literaltype == "" {
			switch tok {
			case parser.INTLIT:
				literaltype = "xsd:integer"
			case parser.FLOATLIT:
				literaltype = "xsd:decimal"
			case parser.STRINGLIT:
				literaltype = "xsd:string"
			}
		} else {
			// numbers can be quoted like "123" or "0.01".
			// The lexer syntactically decides for string token.
			// Correct token type if explicit number type is given, and value fits:
			if tok == parser.STRINGLIT {
				if mustTok, ok := builtindatatypes.BuiltinDatatypes[literaltype]; ok {
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
			err = literaltypeMismatch(tok, literaltype)
		}
		if err != nil {
			err = pos.EnrichErrorMsg(err, "parsing literal")
			return
		}
		l = literal.OWLLiteral{Value: lit, LangTag: langtag, Literaltype: literaltype}
		return
	default:
		err = pos.Errorf("unexpected %v when parsing literal", parser.DescribeToklit(tok, lit))
	}
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

// parseSuffixLiteraltype returns "xsd:integer", if "^^xsd:integer" is found.
// Empty string if not ^^... is found. Error if ^^<syntactically-invalid-literaltype> is found.
func parseSuffixLiteraltype(p *parser.Parser, prefixes tech.Prefixes) (literaltype string, err error) {
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
	if !prefixes.IsPrefixKnown(prefix) {
		p.Unscan()
		err = pos.Errorf("unknown prefix %v in literal type", prefix)
		return
	}
	literaltype = parser.FmtPrefixedName(prefix, name)
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
