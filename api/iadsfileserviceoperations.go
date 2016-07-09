package api

import "unsafe"

// IADsFileServiceOperationsVtbl represents the component object model virtual
// function table for the IADsFileServiceOperations interface.
type IADsFileServiceOperationsVtbl struct {
	IADsServiceOperationsVtbl
	Sessions  uintptr
	Resources uintptr
}

// IADsFileServiceOperations represents the component object model interface for
// active directory file service operations.
type IADsFileServiceOperations struct {
	IADsServiceOperations
}

// VTable returns the component object model virtual function table for the
// file service operations.
func (v *IADsFileServiceOperations) VTable() *IADsFileServiceOperationsVtbl {
	return (*IADsFileServiceOperationsVtbl)(unsafe.Pointer(v.RawVTable))
}
