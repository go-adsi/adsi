//go:build !windows
// +build !windows

package api

import (
	"github.com/go-ole/go-ole"
	"github.com/google/uuid"
)

// NewIADsContainer returns a new instance of the IADsContainer
// component object model interface.
func NewIADsContainer(server string, clsid uuid.UUID) (*IADsContainer, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

// NewEnum retrieves an enumerator interface that provides access to the objects
// within the container.
//
// See https://msdn.microsoft.com/library/aa705990
func (v *IADsContainer) NewEnum() (enum *ole.IUnknown, err error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

// Filter retrieves the filter for the container.
func (v *IADsContainer) Filter() (variant *ole.VARIANT, err error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

// GetObject returns a descendant object with the given class and relative
// name.
func (v *IADsContainer) GetObject(class, name string) (obj *ole.IDispatch, err error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

// SetFilter sets the filter for the container.
func (v *IADsContainer) SetFilter(variant *ole.VARIANT) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
