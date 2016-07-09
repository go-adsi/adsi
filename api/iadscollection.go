package api

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

// IADsCollectionVtbl represents the component object model virtual
// function table for the IADsCollection interface.
type IADsCollectionVtbl struct {
	ole.IDispatchVtbl
	NewEnum uintptr
	Add     uintptr
	Remove  uintptr
	Object  uintptr
}

// IADsCollection represents the component object model interface for
// active directory collections.
type IADsCollection struct {
	ole.IDispatch
}

// VTable returns the component object model virtual function table for the
// collection.
func (v *IADsCollection) VTable() *IADsCollectionVtbl {
	return (*IADsCollectionVtbl)(unsafe.Pointer(v.RawVTable))
}
