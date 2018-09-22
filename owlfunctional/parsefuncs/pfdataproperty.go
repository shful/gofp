package parsefuncs

import (
	"reifenberg.de/gofp/owlfunctional/meta"
	"reifenberg.de/gofp/owlfunctional/parser"
	"reifenberg.de/gofp/owlfunctional/properties"
	"reifenberg.de/gofp/parsehelper"
	"reifenberg.de/gofp/tech"
)

func ParseDataProperty(p *parser.Parser, decls tech.Declarations, prefixes tech.Prefixes) (expr meta.DataProperty, err error) {

	pos := p.Pos()
	// must be R
	var prefix, name string
	prefix, name, err = parsehelper.ParsePrefixedName(p)

	if err != nil {
		return
	}

	if prefixes.IsOWL(prefix) {
		// must be one of the predefined OWL property names
		switch name {
		case "topDataProperty":
			expr = &properties.OWLTopDataProperty{}
		case "bottomDataProperty":
			expr = &properties.OWLBottomDataProperty{}
		default:
			err = pos.Errorf(`unexpected OWL property "%v"`, name)
		}
		return
	}

	if !prefixes.IsPrefixKnown(prefix) {
		err = pos.Errorf("Unknown prefix when parsing DataProperty (%v)", prefix)
	}

	var ok bool
	expr, ok = decls.GetDataPropertyDecl(prefix, name)
	if !ok {
		err = pos.Errorf("Unknown ref to %v:%v", prefix, name)
	}
	return
}
