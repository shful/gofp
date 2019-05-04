package meta

type AnnotationProperty interface {
}

// ClassExpression is one of: a named class,
// Thing or Nothing, a Boolean Connective, an Enumeration,
// or a Property Restriction.
type ClassExpression interface {
	// IsNamedClass is a marker method. true for user declared (primitive) classes only.
	IsNamedClass() bool
}

type DataProperty interface {
	IsNamedDataProperty() bool
}

// DataRange is s
type DataRange interface {
	IsNamedDatatype() bool
}

// NamedDatatype is shortened as DN
type NamedDatatype interface {
	DataRange
}

// ObjectPropertyExpression is one of: a named property,
// the top or bottom property, or the inverse of a named property.
type ObjectPropertyExpression interface {
	// IsNamedObjectProperty is a marker method. true for user declared properties only.
	IsNamedObjectProperty() bool
}
