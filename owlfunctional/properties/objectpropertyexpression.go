package properties

import (
	"reifenberg.de/gofp/owlfunctional/meta"
)

type OWLTopObjectProperty struct {
}

var _ meta.ObjectPropertyExpression = (*OWLTopObjectProperty)(nil)

func (*OWLTopObjectProperty) IsNamedObjectProperty() bool {
	return false
}

type OWLBottomObjectProperty struct {
}

var _ meta.ObjectPropertyExpression = (*OWLBottomObjectProperty)(nil)

func (*OWLBottomObjectProperty) IsNamedObjectProperty() bool {
	return false
}

type ObjectInverseOf struct {
	PN string
}

var _ meta.ObjectPropertyExpression = (*ObjectInverseOf)(nil)

func (*ObjectInverseOf) IsNamedObjectProperty() bool {
	return false
}
