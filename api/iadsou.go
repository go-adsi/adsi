package api

import "unsafe"

// IADsOUVtbl represents the component object model virtual
// function table for the IADsOU interface.
type IADsOUVtbl struct {
	IADsVtbl
	Description         uintptr
	SetDescription      uintptr
	LocalityName        uintptr
	SetLocalityName     uintptr
	PostalAddress       uintptr
	SetPostalAddress    uintptr
	TelephoneNumber     uintptr
	SetTelephoneNumber  uintptr
	FaxNumber           uintptr
	SetFaxNumber        uintptr
	SeeAlso             uintptr
	SetSeeAlso          uintptr
	BusinessCategory    uintptr
	SetBusinessCategory uintptr
}

// IADsOU represents the component object model interface for
// active directory organizational units.
type IADsOU struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// organizational unit.
func (v *IADsOU) VTable() *IADsOUVtbl {
	return (*IADsOUVtbl)(unsafe.Pointer(v.RawVTable))
}
