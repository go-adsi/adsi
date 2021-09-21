//go:build windows
// +build windows

package api

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// AccountDisabled retrieves the disablement status of a user account.
func (v *IADsUser) AccountDisabled() (disabled bool, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().AccountDisabled),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&disabled)),
		0)
	if hr != 0 {
		return false, convertHresultToError(hr)
	}
	return
}

// SetAccountDisabled sets an account as disabled.
func (v *IADsUser) SetAccountDisabled(disabled bool) (err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().SetAccountDisabled),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&disabled)),
		0)
	if hr != 0 {
		return convertHresultToError(hr)
	}
	return
}

// FullName returns the user's FullName property.
func (v *IADsUser) FullName() (name string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().FullName),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr != 0 {
		return "", convertHresultToError(hr)
	}
	name = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	return
}
