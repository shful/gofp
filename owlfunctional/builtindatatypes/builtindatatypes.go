package builtindatatypes

import "github.com/shful/gofp/owlfunctional/parser"

const (
	PRE_OWL  = "http://www.w3.org/2002/07/owl#"
	PRE_RDF  = "http://www.w3.org/1999/02/22-rdf-syntax-ns#"
	PRE_RDFS = "http://www.w3.org/2000/01/rdf-schema#"
	PRE_XML  = "http://www.w3.org/XML/1998/namespace"
	PRE_XSD  = "http://www.w3.org/2001/XMLSchema#"
)

// BuiltinDatatypes are the IRIs predefined in OWL.
// The mapped value is the parsed token for the IRI.
// For example, owl:rational and owl:read result in the same token because they are equally parsed into Golangs floating type.
var BuiltinDatatypes map[string]parser.Token = map[string]parser.Token{
	PRE_OWL + "rational":           parser.FLOATLIT,
	PRE_OWL + "real":               parser.FLOATLIT,
	PRE_RDF + "#PlainLiteral":      parser.STRINGLIT,
	PRE_RDF + "#XMLLiteral":        parser.STRINGLIT,
	PRE_RDFS + "#Literal":          parser.STRINGLIT,
	PRE_XSD + "anyURI":             parser.STRINGLIT,
	PRE_XSD + "base64Binary":       parser.STRINGLIT,
	PRE_XSD + "boolean":            parser.STRINGLIT,
	PRE_XSD + "byte":               parser.INTLIT,
	PRE_XSD + "dateTime":           parser.STRINGLIT,
	PRE_XSD + "dateTimeStamp":      parser.STRINGLIT,
	PRE_XSD + "decimal":            parser.FLOATLIT,
	PRE_XSD + "double":             parser.STRINGLIT,
	PRE_XSD + "hexBinary":          parser.STRINGLIT,
	PRE_XSD + "int":                parser.INTLIT,
	PRE_XSD + "integer":            parser.INTLIT,
	PRE_XSD + "long":               parser.INTLIT,
	PRE_XSD + "Name":               parser.STRINGLIT,
	PRE_XSD + "negativeInteger":    parser.INTLIT,
	PRE_XSD + "NMTOKEN":            parser.STRINGLIT,
	PRE_XSD + "nonNegativeInteger": parser.INTLIT,
	PRE_XSD + "nonPositiveInteger": parser.INTLIT,
	PRE_XSD + "normalizedString":   parser.STRINGLIT,
	PRE_XSD + "positiveInteger":    parser.INTLIT,
	PRE_XSD + "short":              parser.INTLIT,
	PRE_XSD + "string":             parser.STRINGLIT,
	PRE_XSD + "token":              parser.STRINGLIT,
	PRE_XSD + "unsignedByte":       parser.INTLIT,
	PRE_XSD + "unsignedInt":        parser.INTLIT,
	PRE_XSD + "unsignedLong":       parser.INTLIT,
	PRE_XSD + "unsignedShort":      parser.INTLIT,
}

func BuiltinDatatypeExists(iri string) bool {
	_, ok := BuiltinDatatypes[iri]
	return ok
}
