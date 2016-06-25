// +build windows

package api

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// ComputerID retrieves the globally unique identifier of the computer.
func (v *IADsComputer) ComputerID() (id string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().ComputerID),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr == 0 {
		id = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	} else {
		return "", convertHresultToError(hr)
	}
	return
}

// Site retrieves the site of the computer.
func (v *IADsComputer) Site() (site string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Site),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr == 0 {
		site = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	} else {
		return "", convertHresultToError(hr)
	}
	return
}

// OperatingSystem retrieves the operating system of the computer.
func (v *IADsComputer) OperatingSystem() (os string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().OperatingSystem),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr == 0 {
		os = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	} else {
		return "", convertHresultToError(hr)
	}
	return
}
