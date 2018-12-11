package parsefuncs

import (
	"reifenberg.de/gofp/owlfunctional/meta"
	"reifenberg.de/gofp/owlfunctional/parser"
	"reifenberg.de/gofp/owlfunctional/properties"
	"reifenberg.de/gofp/parsehelper"
	"reifenberg.de/gofp/tech"
)

func ParseObjectPropertyExpression(p *parser.Parser, decls tech.Declarations, prefixes tech.Prefixes) (expr meta.ObjectPropertyExpression, err error) {
	tok, lit, pos := p.ScanIgnoreWSAndComment()
	switch tok {
	case parser.ObjectInverseOf:
		p.Unscan()
		expr, err = parseObjectInverseOf(p, prefixes)
	default:
		p.Unscan()
		// must be PN
		var ident *tech.IRI
		ident, err = parsehelper.ParseAndResolveIRI(p, prefixes)
		if err != nil {
			err = pos.ErrorfUnexpectedToken(tok, lit, "Object Property Name")
			return
		}

		if ident.IsOWL() {
			// must be one of the predefined OWL property names
			switch ident.Name {
			case "topObjectProperty":
				expr = &properties.OWLTopObjectProperty{}
			case "bottomObjectProperty":
				expr = &properties.OWLBottomObjectProperty{}
			default:
				err = pos.Errorf(`unexpected OWL property "%v"`, ident.Name)
			}
			return
		}

		var ok bool
		expr, ok = decls.GetObjectPropertyDecl(*ident)
		if !ok {
			err = pos.Errorf("Unknown ref to %v. Expected object property name.", ident)
		}
	}
	return
}

func parseObjectInverseOf(p *parser.Parser, prefixes tech.Prefixes) (expr meta.ObjectPropertyExpression, err error) {
	if err = p.ConsumeTokens(parser.ObjectInverseOf, parser.B1); err != nil {
		return
	}
	pos := p.Pos()
	var prefix, name string
	prefix, name, err = parsehelper.ParsePrefixedName(p)
	if err != nil {
		return
	}
	if !prefixes.IsPrefixKnown(prefix) {
		err = pos.Errorf("unknown prefix in ObjectInverseOf (%v)", prefix)
		return
	}
	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}
	expr = &properties.ObjectInverseOf{PN: parser.FmtPrefixedName(prefix, name)}
	return
}
