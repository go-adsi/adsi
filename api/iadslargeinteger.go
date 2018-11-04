package api

import (
	"unsafe"

	ole "github.com/go-ole/go-ole"
)

// IADsLargeIntegerVtbl represents the component object model virtual
// function table for the IADsLargeInteger interface.
type IADsLargeIntegerVtbl struct {
	ole.IDispatchVtbl
	HighPart    uintptr
	SetHighPart uintptr
	LowPart     uintptr
	SetLowPart  uintptr
}

// IADsLargeInteger represents the component object model interface for
// large integers.
type IADsLargeInteger struct {
	ole.IDispatch
}

// VTable returns the component object model virtual function table for the
// large integer.
func (v *IADsLargeInteger) VTable() *IADsLargeIntegerVtbl {
	return (*IADsLargeIntegerVtbl)(unsafe.Pointer(v.RawVTable))
}
