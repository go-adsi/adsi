package api

import "unsafe"

// IADsServiceVtbl represents the component object model virtual
// function table for the IADsService interface.
type IADsServiceVtbl struct {
	IADsVtbl
	HostComputer          uintptr
	SetHostComputer       uintptr
	DisplayName           uintptr
	SetDisplayName        uintptr
	Version               uintptr
	SetVersion            uintptr
	ServiceType           uintptr
	SetServiceType        uintptr
	StartType             uintptr
	SetStartType          uintptr
	Path                  uintptr
	SetPath               uintptr
	StartupParameters     uintptr
	SetStartupParameters  uintptr
	ErrorControl          uintptr
	SetErrorControl       uintptr
	LoadOrderGroup        uintptr
	SetLoadOrderGroup     uintptr
	ServiceAccountName    uintptr
	SetServiceAccountName uintptr
	ServiceAccountPath    uintptr
	SetServiceAccountPath uintptr
	Dependencies          uintptr
	SetDependencies       uintptr
}

// IADsService represents the component object model interface for
// active directory services.
type IADsService struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// service.
func (v *IADsService) VTable() *IADsServiceVtbl {
	return (*IADsServiceVtbl)(unsafe.Pointer(v.RawVTable))
}
