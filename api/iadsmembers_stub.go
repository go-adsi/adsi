// +build !windows

package api

import "github.com/go-ole/go-ole"

// NewEnum retrieves an enumerator interface that provides access to the objects
// within the membership.
//
// See https://msdn.microsoft.com/library/aa706042
func (v *IADsMembers) NewEnum() (enum *ole.IUnknown, err error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

// Filter retrieves the filter for the membership.
func (v *IADsMembers) Filter() (variant *ole.VARIANT, err error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

// SetFilter sets the filter for the membership.
func (v *IADsMembers) SetFilter(variant *ole.VARIANT) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
