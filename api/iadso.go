package api

import "unsafe"

// IADsOVtbl represents the component object model virtual
// function table for the IADsO interface.
type IADsOVtbl struct {
	IADsVtbl
	Description        uintptr
	SetDescription     uintptr
	LocalityName       uintptr
	SetLocalityName    uintptr
	PostalAddress      uintptr
	SetPostalAddress   uintptr
	TelephoneNumber    uintptr
	SetTelephoneNumber uintptr
	FaxNumber          uintptr
	SetFaxNumber       uintptr
	SeeAlso            uintptr
	SetSeeAlso         uintptr
}

// IADsO represents the component object model interface for
// active directory organizations.
type IADsO struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// organization.
func (v *IADsO) VTable() *IADsOVtbl {
	return (*IADsOVtbl)(unsafe.Pointer(v.RawVTable))
}
