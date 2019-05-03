package annotations

type Annotation struct {
	// Value comes without type info and without language tag.
	A string

	// langTag can be set on strings. The literal "foo"@en results in langTag="en".
	// If not given, langTag is empty. For non-string types, it is also empty.
	T string
}

type AnnotationAssertion struct {
	// Value comes without type info and without language tag.
	A string

	// Literaltype can be one of the builtin datatypes like xsd:integer, or a custom IRI.
	S string

	// langTag can be set on strings. The literal "foo"@en results in langTag="en".
	// If not given, langTag is empty. For non-string types, it is also empty.
	T string
}
