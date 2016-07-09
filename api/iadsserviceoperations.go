package api

import "unsafe"

// IADsServiceOperationsVtbl represents the component object model virtual
// function table for the IADsServiceOperations interface.
type IADsServiceOperationsVtbl struct {
	IADsVtbl
	Status      uintptr
	Start       uintptr
	Stop        uintptr
	Pause       uintptr
	Continue    uintptr
	SetPassword uintptr
}

// IADsServiceOperations represents the component object model interface for
// active directory service operations.
type IADsServiceOperations struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// service operations.
func (v *IADsServiceOperations) VTable() *IADsServiceOperationsVtbl {
	return (*IADsServiceOperationsVtbl)(unsafe.Pointer(v.RawVTable))
}
