package api

import "unsafe"

// IADsComputerOperationsVtbl represents the component object model virtual
// function table for the IADsComputerOperations interface.
type IADsComputerOperationsVtbl struct {
	IADsVtbl
	Status   uintptr
	Shutdown uintptr
}

// IADsComputerOperations represents the component object model interface for
// active directory computer operations.
type IADsComputerOperations struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// computer operations.
func (v *IADsComputerOperations) VTable() *IADsComputerOperationsVtbl {
	return (*IADsComputerOperationsVtbl)(unsafe.Pointer(v.RawVTable))
}
