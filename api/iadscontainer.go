package api

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

// IADsContainerVtbl represents the component object model virtual
// function table for the IADsContainer interface.
type IADsContainerVtbl struct {
	ole.IDispatchVtbl
	Count     uintptr
	NewEnum   uintptr
	Filter    uintptr
	SetFilter uintptr
	Hints     uintptr
	SetHints  uintptr
	GetObject uintptr
	Create    uintptr
	Delete    uintptr
	CopyHere  uintptr
	MoveHere  uintptr
}

// IADsContainer represents the component object model interface for
// active directory containers.
type IADsContainer struct {
	ole.IDispatch
}

// VTable returns the component object model virtual function table for the
// container.
func (v *IADsContainer) VTable() *IADsContainerVtbl {
	return (*IADsContainerVtbl)(unsafe.Pointer(v.RawVTable))
}
