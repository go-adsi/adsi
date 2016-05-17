package api

import "unsafe"

// IADsGroupVtbl represents the component object model virtual
// function table for the IADsGroup interface.
type IADsGroupVtbl struct {
	IADsVtbl
	Description    uintptr
	SetDescription uintptr
	Members        uintptr
	IsMember       uintptr
	Add            uintptr
	Remove         uintptr
}

// IADsGroup represents the component object model interface for
// active directory groups.
type IADsGroup struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// group.
func (v *IADsGroup) VTable() *IADsGroupVtbl {
	return (*IADsGroupVtbl)(unsafe.Pointer(v.RawVTable))
}
