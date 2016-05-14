// +build !windows

package api

import "github.com/go-ole/go-ole"

// NewEnum retrieves an enumerator interface that provides access to the objects
// within the container.
//
// See https://msdn.microsoft.com/library/aa705990
func (v *IADsContainer) NewEnum() (enum *ole.IEnumVARIANT, err error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}
