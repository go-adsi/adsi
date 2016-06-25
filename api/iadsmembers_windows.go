// +build windows

package api

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// NewEnum retrieves an enumerator interface that provides access to the objects
// within the membership.
//
// See https://msdn.microsoft.com/library/aa706042
func (v *IADsMembers) NewEnum() (enum *ole.IUnknown, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().NewEnum),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&enum)),
		0)
	if hr != 0 {
		return nil, convertHresultToError(hr)
	}
	return
}

// Filter retrieves the filter for the membership.
func (v *IADsMembers) Filter() (variant *ole.VARIANT, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Filter),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&variant)),
		0)
	if hr != 0 {
		return nil, convertHresultToError(hr)
	}
	return
}

// SetFilter sets the filter for the membership.
func (v *IADsMembers) SetFilter(variant *ole.VARIANT) (err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().SetFilter),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(variant)),
		0)
	if hr != 0 {
		return convertHresultToError(hr)
	}
	return nil
}
