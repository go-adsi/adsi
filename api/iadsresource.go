package api

import "unsafe"

// IADsResourceVtbl represents the component object model virtual
// function table for the IADsResource interface.
type IADsResourceVtbl struct {
	IADsVtbl
	User      uintptr
	UserPath  uintptr
	Path      uintptr
	LockCount uintptr
}

// IADsResource represents the component object model interface for
// active directory resources.
type IADsResource struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// resource.
func (v *IADsResource) VTable() *IADsResourceVtbl {
	return (*IADsResourceVtbl)(unsafe.Pointer(v.RawVTable))
}
