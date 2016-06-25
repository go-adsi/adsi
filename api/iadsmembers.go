package api

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

// IADsMembersVtbl represents the component object model virtual
// function table for the IADsMembers interface.
type IADsMembersVtbl struct {
	ole.IDispatchVtbl
	Count     uintptr
	NewEnum   uintptr
	Filter    uintptr
	SetFilter uintptr
}

// IADsMembers represents the component object model interface for group
// membership.
type IADsMembers struct {
	ole.IDispatch
}

// VTable returns the component object model virtual function table for the
// group membership.
func (v *IADsMembers) VTable() *IADsMembersVtbl {
	return (*IADsMembersVtbl)(unsafe.Pointer(v.RawVTable))
}
