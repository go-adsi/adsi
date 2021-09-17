// +build windows

package api

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// Add adds an ADSI object to an existing group.
func (v *IADsGroup) Add(member string) (err error) {
	m := ole.SysAllocStringLen(member)
	if m == nil {
		return ole.NewError(ole.E_OUTOFMEMORY)
	}
	defer ole.SysFreeString(m)
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Add),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(m)),
		0)
	if hr != 0 {
		return convertHresultToError(hr)
	}
	return
}

// Description retrieves the description of the group.
func (v *IADsGroup) Description() (desc string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Description),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr == 0 {
		desc = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	} else {
		return "", convertHresultToError(hr)
	}
	return
}

// Members retrieves an IADsMembers interface that provides access to the
// membership of the group.
func (v *IADsGroup) Members() (members *IADsMembers, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Members),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&members)),
		0)
	if hr != 0 {
		return nil, convertHresultToError(hr)
	}
	return
}

// Remove removes the specified user object from this group. The operation
// does not remove the group object itself even when there is no member remaining in the group.
func (v *IADsGroup) Remove(member string) (err error) {
	m := ole.SysAllocStringLen(member)
	if m == nil {
		return ole.NewError(ole.E_OUTOFMEMORY)
	}
	defer ole.SysFreeString(m)
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Remove),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(m)),
		0)
	if hr != 0 {
		return convertHresultToError(hr)
	}
	return
}
