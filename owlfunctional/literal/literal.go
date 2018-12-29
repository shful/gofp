package literal

import (
	"fmt"

	"github.com/shful/gofp/owlfunctional/parser"
)

type OWLLiteral struct {
	// Value comes without type info and without language tag.
	Value string

	// Literaltype can be one of the builtin datatypes like xsd:integer, or a custom IRI.
	Literaltype string

	// langTag can be set on strings. The literal "foo"@en results in langTag="en".
	// If not given, langTag is empty. For non-string types, it is also empty.
	LangTag string
}

// LiteralString reconstructs the literal as written in OWL functional.
func (s *OWLLiteral) LiteralString() string {
	res := fmt.Sprintf(`"%v"`, s.Value)
	if s.LangTag != "" {
		res = fmt.Sprintf("%v@%v", res, s.LangTag)
	}
	if s.Literaltype != "" {
		res = fmt.Sprintf("%v^^%v", res, s.Literaltype)
	}
	return res
}

// MaybeOWLLiteral is true when this token can be a valid literal expression.
func MaybeOWLLiteral(tok parser.Token) bool {
	switch tok {
	case parser.OWLTrue, parser.OWLFalse, parser.STRINGLIT, parser.INTLIT, parser.FLOATLIT:
		return true
	}
	return false
}
