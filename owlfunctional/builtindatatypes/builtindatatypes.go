package builtindatatypes

import "reifenberg.de/gofp/owlfunctional/parser"

// BuiltinDatatypes are the IRIs predefined in OWL.
// The mapped value is the parsed token for the IRI.
// For example, owl:rational and owl:read result in the same token because they are equally parsed into Golangs floating type.
var BuiltinDatatypes map[string]parser.Token = map[string]parser.Token{
	"owl:rational":           parser.FLOATLIT,
	"owl:real":               parser.FLOATLIT,
	"rdf:PlainLiteral":       parser.STRINGLIT,
	"rdf:XMLLiteral":         parser.STRINGLIT,
	"rdfs:Literal":           parser.STRINGLIT,
	"xsd:anyURI":             parser.STRINGLIT,
	"xsd:base64Binary":       parser.STRINGLIT,
	"xsd:boolean":            parser.STRINGLIT,
	"xsd:byte":               parser.INTLIT,
	"xsd:dateTime":           parser.STRINGLIT,
	"xsd:dateTimeStamp":      parser.STRINGLIT,
	"xsd:decimal":            parser.FLOATLIT,
	"xsd:double":             parser.STRINGLIT,
	"xsd:hexBinary":          parser.STRINGLIT,
	"xsd:int":                parser.INTLIT,
	"xsd:integer":            parser.INTLIT,
	"xsd:long":               parser.INTLIT,
	"xsd:Name":               parser.STRINGLIT,
	"xsd:negativeInteger":    parser.INTLIT,
	"xsd:NMTOKEN":            parser.STRINGLIT,
	"xsd:nonNegativeInteger": parser.INTLIT,
	"xsd:nonPositiveInteger": parser.INTLIT,
	"xsd:normalizedString":   parser.STRINGLIT,
	"xsd:positiveInteger":    parser.INTLIT,
	"xsd:short":              parser.INTLIT,
	"xsd:string":             parser.STRINGLIT,
	"xsd:token":              parser.STRINGLIT,
	"xsd:unsignedByte":       parser.INTLIT,
	"xsd:unsignedInt":        parser.INTLIT,
	"xsd:unsignedLong":       parser.INTLIT,
	"xsd:unsignedShort":      parser.INTLIT,
}

func BuiltinDatatypeExists(iri string) bool {
	_, ok := BuiltinDatatypes[iri]
	return ok
}
