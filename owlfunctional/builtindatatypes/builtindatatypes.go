package builtindatatypes

import "reifenberg.de/gofp/owlfunctional/parser"

// BuiltinDatatypes are the IRIs predefined in OWL.
// The mapped value is the parsed token for the IRI.
// For example, owl:rational and owl:read result in the same token because they are equally parsed into Golangs floating type.
var BuiltinDatatypes map[string]parser.Token = map[string]parser.Token{ //todo: replace prefixes by IRI
	"http://www.w3.org/2002/07/owl#rational":                  parser.FLOATLIT,
	"http://www.w3.org/2002/07/owl#real":                      parser.FLOATLIT,
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#PlainLiteral": parser.STRINGLIT,
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#XMLLiteral":   parser.STRINGLIT,
	"http://www.w3.org/2000/01/rdf-schema#Literal":            parser.STRINGLIT,
	"http://www.w3.org/2001/XMLSchema#anyURI":                 parser.STRINGLIT,
	"http://www.w3.org/2001/XMLSchema#base64Binary":           parser.STRINGLIT,
	"http://www.w3.org/2001/XMLSchema#boolean":                parser.STRINGLIT,
	"http://www.w3.org/2001/XMLSchema#byte":                   parser.INTLIT,
	"http://www.w3.org/2001/XMLSchema#dateTime":               parser.STRINGLIT,
	"http://www.w3.org/2001/XMLSchema#dateTimeStamp":          parser.STRINGLIT,
	"http://www.w3.org/2001/XMLSchema#decimal":                parser.FLOATLIT,
	"http://www.w3.org/2001/XMLSchema#double":                 parser.STRINGLIT,
	"http://www.w3.org/2001/XMLSchema#hexBinary":              parser.STRINGLIT,
	"http://www.w3.org/2001/XMLSchema#int":                    parser.INTLIT,
	"http://www.w3.org/2001/XMLSchema#integer":                parser.INTLIT,
	"http://www.w3.org/2001/XMLSchema#long":                   parser.INTLIT,
	"http://www.w3.org/2001/XMLSchema#Name":                   parser.STRINGLIT,
	"http://www.w3.org/2001/XMLSchema#negativeInteger":        parser.INTLIT,
	"http://www.w3.org/2001/XMLSchema#NMTOKEN":                parser.STRINGLIT,
	"http://www.w3.org/2001/XMLSchema#nonNegativeInteger":     parser.INTLIT,
	"http://www.w3.org/2001/XMLSchema#nonPositiveInteger":     parser.INTLIT,
	"http://www.w3.org/2001/XMLSchema#normalizedString":       parser.STRINGLIT,
	"http://www.w3.org/2001/XMLSchema#positiveInteger":        parser.INTLIT,
	"http://www.w3.org/2001/XMLSchema#short":                  parser.INTLIT,
	"http://www.w3.org/2001/XMLSchema#string":                 parser.STRINGLIT,
	"http://www.w3.org/2001/XMLSchema#token":                  parser.STRINGLIT,
	"http://www.w3.org/2001/XMLSchema#unsignedByte":           parser.INTLIT,
	"http://www.w3.org/2001/XMLSchema#unsignedInt":            parser.INTLIT,
	"http://www.w3.org/2001/XMLSchema#unsignedLong":           parser.INTLIT,
	"http://www.w3.org/2001/XMLSchema#unsignedShort":          parser.INTLIT,
}

func BuiltinDatatypeExists(iri string) bool {
	_, ok := BuiltinDatatypes[iri]
	return ok
}
