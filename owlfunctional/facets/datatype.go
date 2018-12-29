package facets

import (
	"github.com/shful/gofp/owlfunctional/meta"
)

type NamedDatatypeImpl struct {
	DatatypeIRI string
}

var _ meta.NamedDatatype = (*NamedDatatypeImpl)(nil)

func (s *NamedDatatypeImpl) IsNamedDatatype() bool {
	return true
}

type BuiltinDatatype struct {
	NamedDatatypeImpl
}

// CustomNamedDatatype is any NamedDatatype which is not a BuiltinDatatype.
type CustomNamedDatatype struct {
	NamedDatatypeImpl
}

type DatatypeRestriction struct {
	DN      meta.NamedDatatype
	FVPairs []*FVPair
}
