package api

import "unsafe"

// IADsFileShareVtbl represents the component object model virtual
// function table for the IADsFileShare interface.
type IADsFileShareVtbl struct {
	IADsVtbl
	CurrentUserCount uintptr
	Description      uintptr
	SetDescription   uintptr
	HostComputer     uintptr
	SetHostComputer  uintptr
	Path             uintptr
	SetPath          uintptr
	MaxUserCount     uintptr
	SetMaxUserCount  uintptr
}

// IADsFileShare represents the component object model interface for
// active directory file shares.
type IADsFileShare struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// file share.
func (v *IADsFileShare) VTable() *IADsFileShareVtbl {
	return (*IADsFileShareVtbl)(unsafe.Pointer(v.RawVTable))
}
