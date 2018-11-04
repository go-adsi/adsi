// +build !windows

package api

import ole "github.com/go-ole/go-ole"

// HighPart retrieves the upper 32 bits of the 64 bit value.
func (v *IADsLargeInteger) HighPart() (upper int32, err error) {
	return 0, ole.NewError(ole.E_NOTIMPL)
}

// LowPart retrieves the lower 32 bits of the 64 bit value.
func (v *IADsLargeInteger) LowPart() (lower int32, err error) {
	return 0, ole.NewError(ole.E_NOTIMPL)
}

// Value retrieves the 64 bit value.
func (v *IADsLargeInteger) Value() (value int64, err error) {
	return 0, ole.NewError(ole.E_NOTIMPL)
}
