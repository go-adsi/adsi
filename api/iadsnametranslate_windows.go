//go:build windows
// +build windows

package api

import (
	"syscall"
	"unsafe"

	"github.com/go-adsi/adsi/comclsid"
	"github.com/go-adsi/adsi/comiid"
	"github.com/go-ole/go-ole"
	"github.com/scjalliance/comutil"
)

// NewIADsNameTranslate returns an IADsNameTranslate that manages the given COM interface.
func NewIADsNameTranslate(server string) (*IADsNameTranslate, error) {
	p, err := comutil.CreateRemoteObject(server, comclsid.NameTranslate, comiid.IADsNameTranslate)
	return (*IADsNameTranslate)(unsafe.Pointer(p)), err
}

// Get retrieves the name of a directory object in the specified format.
// The distinguished name must have been set in the appropriate format by the Set function.
func (v *IADsNameTranslate) Get(formatType uint32) (adsPath string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Get),
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(formatType),
		uintptr(unsafe.Pointer(&bstr)))
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr == 0 {
		adsPath = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	} else {
		return "", convertHresultToError(hr)
	}
	return
}

// Init initializes a name translate object by binding to a specified directory server, domain,
// or global catalog, using the credentials of the current user.
func (v *IADsNameTranslate) Init(adsPath string, initType uint32) (err error) {
	bname := ole.SysAllocStringLen(adsPath)
	if bname == nil {
		return ole.NewError(ole.E_OUTOFMEMORY)
	}
	defer ole.SysFreeString(bname)
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Init),
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(initType),
		uintptr(unsafe.Pointer(bname)))
	if hr != 0 {
		return convertHresultToError(hr)
	}
	return
}

// Set directs the directory service to set up a specified object for name translation.
func (v *IADsNameTranslate) Set(adsPath string, setType uint32) (err error) {
	bname := ole.SysAllocStringLen(adsPath)
	if bname == nil {
		return ole.NewError(ole.E_OUTOFMEMORY)
	}
	defer ole.SysFreeString(bname)
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Set),
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(setType),
		uintptr(unsafe.Pointer(bname)))
	if hr != 0 {
		return convertHresultToError(hr)
	}
	return
}
