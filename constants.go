package adsi

import "errors"

var (
	// ErrClosed is returned from calls to a service or interface in the event
	// that the Close() function has already been called.
	ErrClosed = errors.New("Interface is closing or already closed.")

	// ErrNonDispatchVariant is returned when an attempt is made to cast an
	// ole.VARIANT to an ole.IDispatch type, but the VARIANT is of some other
	// This might happen, for example, when an iterator is interrogating the
	// members of an IEnumVARIANT in an attempt to convert them into an expected
	// type.
	ErrNonDispatchVariant = errors.New("Object iterator unexpectedly yielded non-dispatch variant.")
)
