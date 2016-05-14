package api

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

// IADsVtbl represents the component object model virtual
// function table for the IADs interface.
type IADsVtbl struct {
	ole.IDispatchVtbl
	Name      uintptr
	Class     uintptr
	GUID      uintptr
	AdsPath   uintptr
	Parent    uintptr
	Schema    uintptr
	GetInfo   uintptr
	SetInfo   uintptr
	Get       uintptr
	Put       uintptr
	GetEx     uintptr
	PutEx     uintptr
	GetInfoEx uintptr
}

// IADs represents the component object model interface for
// active directory objects.
type IADs struct {
	ole.IDispatch
}

// VTable returns the component object model virtual function table for the
// object.
func (v *IADs) VTable() *IADsVtbl {
	return (*IADsVtbl)(unsafe.Pointer(v.RawVTable))
}
