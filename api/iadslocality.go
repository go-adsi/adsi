package api

import "unsafe"

// IADsLocalityVtbl represents the component object model virtual
// function table for the IADsLocality interface.
type IADsLocalityVtbl struct {
	IADsVtbl
	Description      uintptr
	SetDescription   uintptr
	LocalityName     uintptr
	SetLocalityName  uintptr
	PostalAddress    uintptr
	SetPostalAddress uintptr
	SeeAlso          uintptr
	SetSeeAlso       uintptr
}

// IADsLocality represents the component object model interface for
// active directory localities.
type IADsLocality struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// locality.
func (v *IADsLocality) VTable() *IADsLocalityVtbl {
	return (*IADsLocalityVtbl)(unsafe.Pointer(v.RawVTable))
}
