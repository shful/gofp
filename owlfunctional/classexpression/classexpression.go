package classexpression

import (
	"reifenberg.de/gofp/owlfunctional/individual"
	"reifenberg.de/gofp/owlfunctional/literal"
	"reifenberg.de/gofp/owlfunctional/meta"
)

type OWLThing struct {
}

var _ meta.ClassExpression = (*OWLThing)(nil)

func (s *OWLThing) IsNamedClass() bool {
	return false
}

type OWLNothing struct {
}

var _ meta.ClassExpression = (*OWLNothing)(nil)

func (s *OWLNothing) IsNamedClass() bool {
	return false
}

type DataAllValuesFrom struct {
	R meta.DataProperty
	D meta.DataRange
}

var _ meta.ClassExpression = (*DataAllValuesFrom)(nil)

func (s *DataAllValuesFrom) IsNamedClass() bool {
	return false
}

type DataExactCardinality struct {
	N int
	R meta.DataProperty
}

var _ meta.ClassExpression = (*DataExactCardinality)(nil)

func (s *DataExactCardinality) IsNamedClass() bool {
	return false
}

type DataHasValue struct {
	R meta.DataProperty
	V literal.OWLLiteral
}

var _ meta.ClassExpression = (*DataHasValue)(nil)

func (s *DataHasValue) IsNamedClass() bool {
	return false
}

type DataMaxCardinality struct {
	N int
	R meta.DataProperty
}

var _ meta.ClassExpression = (*DataMaxCardinality)(nil)

func (s *DataMaxCardinality) IsNamedClass() bool {
	return false
}

type DataMinCardinality struct {
	N int
	R meta.DataProperty
}

var _ meta.ClassExpression = (*DataMinCardinality)(nil)

func (s *DataMinCardinality) IsNamedClass() bool {
	return false
}

type DataQualifiedExactCardinality struct {
	N int
	R meta.DataProperty
	D meta.DataRange
}

var _ meta.ClassExpression = (*DataQualifiedExactCardinality)(nil)

func (s *DataQualifiedExactCardinality) IsNamedClass() bool {
	return false
}

type DataQualifiedMaxCardinality struct {
	N int
	R meta.DataProperty
	D meta.DataRange
}

var _ meta.ClassExpression = (*DataQualifiedMaxCardinality)(nil)

func (s *DataQualifiedMaxCardinality) IsNamedClass() bool {
	return false
}

type DataQualifiedMinCardinality struct {
	N int
	R meta.DataProperty
	D meta.DataRange
}

var _ meta.ClassExpression = (*DataQualifiedMinCardinality)(nil)

func (s *DataQualifiedMinCardinality) IsNamedClass() bool {
	return false
}

type DataSomeValuesFrom struct {
	R meta.DataProperty
	D meta.DataRange
}

var _ meta.ClassExpression = (*DataSomeValuesFrom)(nil)

func (s *DataSomeValuesFrom) IsNamedClass() bool {
	return false
}

type ObjectAllValuesFrom struct {
	P meta.ObjectPropertyExpression
	C meta.ClassExpression
}

var _ meta.ClassExpression = (*ObjectAllValuesFrom)(nil)

func (s *ObjectAllValuesFrom) IsNamedClass() bool {
	return false
}

type ObjectComplementOf struct {
	C meta.ClassExpression
}

var _ meta.ClassExpression = (*ObjectComplementOf)(nil)

func (s *ObjectComplementOf) IsNamedClass() bool {
	return false
}

type ObjectExactCardinality struct {
	N int
	P meta.ObjectPropertyExpression
}

var _ meta.ClassExpression = (*ObjectExactCardinality)(nil)

func (s *ObjectExactCardinality) IsNamedClass() bool {
	return false
}

type ObjectHasSelf struct {
	P meta.ObjectPropertyExpression
}

var _ meta.ClassExpression = (*ObjectHasSelf)(nil)

func (s *ObjectHasSelf) IsNamedClass() bool {
	return false
}

type ObjectIntersectionOf struct {
	Cs []meta.ClassExpression
}

var _ meta.ClassExpression = (*ObjectIntersectionOf)(nil)

func (s *ObjectIntersectionOf) IsNamedClass() bool {
	return false
}

type ObjectMaxCardinality struct {
	N int
	P meta.ObjectPropertyExpression
}

var _ meta.ClassExpression = (*ObjectMaxCardinality)(nil)

func (s *ObjectMaxCardinality) IsNamedClass() bool {
	return false
}

type ObjectMinCardinality struct {
	N int
	P meta.ObjectPropertyExpression
}

var _ meta.ClassExpression = (*ObjectMinCardinality)(nil)

func (s *ObjectMinCardinality) IsNamedClass() bool {
	return false
}

type ObjectHasValue struct {
	P meta.ObjectPropertyExpression
	A individual.Individual
}

var _ meta.ClassExpression = (*ObjectHasValue)(nil)

func (s *ObjectHasValue) IsNamedClass() bool {
	return false
}

type ObjectSomeValuesFrom struct {
	P meta.ObjectPropertyExpression
	C meta.ClassExpression
}

var _ meta.ClassExpression = (*ObjectSomeValuesFrom)(nil)

func (s *ObjectSomeValuesFrom) IsNamedClass() bool {
	return false
}

type ObjectQualifiedExactCardinality struct {
	N int
	P meta.ObjectPropertyExpression
	C meta.ClassExpression
}

var _ meta.ClassExpression = (*ObjectQualifiedExactCardinality)(nil)

func (s *ObjectQualifiedExactCardinality) IsNamedClass() bool {
	return false
}

type ObjectQualifiedMaxCardinality struct {
	N int
	P meta.ObjectPropertyExpression
	C meta.ClassExpression
}

var _ meta.ClassExpression = (*ObjectQualifiedMaxCardinality)(nil)

func (s *ObjectQualifiedMaxCardinality) IsNamedClass() bool {
	return false
}

type ObjectQualifiedMinCardinality struct {
	N int
	P meta.ObjectPropertyExpression
	C meta.ClassExpression
}

var _ meta.ClassExpression = (*ObjectQualifiedMinCardinality)(nil)

func (s *ObjectQualifiedMinCardinality) IsNamedClass() bool {
	return false
}

type ObjectUnionOf struct {
	Cs []meta.ClassExpression
}

var _ meta.ClassExpression = (*ObjectUnionOf)(nil)

func (s *ObjectUnionOf) IsNamedClass() bool {
	return false
}

type ObjectOneOf struct {
	As []individual.Individual
}

var _ meta.ClassExpression = (*ObjectOneOf)(nil)

func (s *ObjectOneOf) IsNamedClass() bool {
	return false
}
