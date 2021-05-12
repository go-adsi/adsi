package adsi

import (
	"errors"

	"github.com/go-adsi/adsi/api"
)

var (
	// ErrClosed is returned from calls to a service or interface in the event
	// that the Close() function has already been called.
	ErrClosed = errors.New("interface is closing or already closed")

	// ErrNonDispatchVariant is returned when an attempt is made to cast an
	// ole.VARIANT to an ole.IDispatch type, but the VARIANT is of some other
	// This might happen, for example, when an iterator is interrogating the
	// members of an IEnumVARIANT in an attempt to convert them into an expected
	// type.
	ErrNonDispatchVariant = errors.New("object iterator unexpectedly yielded non-dispatch variant")

	// ErrInvalidGUID is returned when a given value cannot be interpreted as
	// a globally unique identifier.
	ErrInvalidGUID = errors.New("invalid GUID")

	// ErrNonArrayAttribute is returned when a given attribute cannot be converted
	// to a safe array.
	ErrNonArrayAttribute = errors.New("attribute is not an array")

	// ErrMultiDimArrayAttribute is returned when an attribute contains more than
	// one dimension in its array of values.
	ErrMultiDimArrayAttribute = errors.New("attribute contains a multi-dimensional array of values")

	// ErrNonVariantArrayAttribute is returned when the array members of a given
	// attribute are not variants.
	ErrNonVariantArrayAttribute = errors.New("attribute contains non-variant array members")
)

const (
	defaultFlags = api.ADS_READONLY_SERVER | api.ADS_SECURE_AUTHENTICATION | api.ADS_USE_SEALING
)
