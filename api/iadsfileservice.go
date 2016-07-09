package api

import "unsafe"

// IADsFileServiceVtbl represents the component object model virtual
// function table for the IADsFileService interface.
type IADsFileServiceVtbl struct {
	IADsServiceVtbl
	Description     uintptr
	SetDescription  uintptr
	MaxUserCount    uintptr
	SetMaxUserCount uintptr
}

// IADsFileService represents the component object model interface for
// active directory file services.
type IADsFileService struct {
	IADsService
}

// VTable returns the component object model virtual function table for the
// file service.
func (v *IADsFileService) VTable() *IADsFileServiceVtbl {
	return (*IADsFileServiceVtbl)(unsafe.Pointer(v.RawVTable))
}
