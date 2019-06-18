package parsefuncs

import (
	"fmt"

	"github.com/shful/gofp/owlfunctional/annotations"
	"github.com/shful/gofp/owlfunctional/meta"
	"github.com/shful/gofp/owlfunctional/parser"
	"github.com/shful/gofp/store"
	"github.com/shful/gofp/tech"
)

// ParseAnnotation parses a single Annotation(...) expression, including braces.
func ParseAnnotation(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (expr *annotations.Annotation, err error) {

	if err = p.ConsumeTokens(parser.Annotation, parser.B1); err != nil {
		return
	}
	pos := p.Pos()

	var A meta.AnnotationProperty
	A, err = ParseA(p, decls, prefixes)
	if err != nil {
		err = pos.EnrichErrorMsg(err, "reading 1st param in Annotation")
		return
	}

	var t string
	t, _, err = Parset(p, decls, prefixes)
	if err != nil {
		err = pos.EnrichErrorMsg(err, "reading 2nd param in Annotation")
		return
	}

	if err = p.ConsumeTokens(parser.B2); err != nil {
		return
	}

	expr = annotations.NewAnnotation(A, t)
	return
}

// ParseAnnotations parses 0..n Annotation(...) expressions, as long as such expressions are found.
// OWL allows these as variadic first arguments of each Axiom, including Declarations.
func ParseAnnotations(p *parser.Parser, decls store.Decls, prefixes tech.Prefixes) (exprs []meta.Annotation, err error) {
	for i := 1; err == nil; i++ {
		// test if one or more Annotation expressions are found
		tok, _, _ := p.ScanIgnoreWSAndComment()
		p.Unscan()
		switch tok {
		case parser.Annotation:
			pos := p.Pos()
			var expr *annotations.Annotation
			expr, err = ParseAnnotation(p, decls, prefixes)
			if err != nil {
				err = pos.EnrichErrorMsg(err, fmt.Sprintf("parsing %d. Annotation", i))
				return
			}
			exprs = append(exprs, expr)
		default:
			// no (more) Annotation
			return
		}
	}
	return
}
