package api

import "unsafe"

// IADsClassVtbl represents the component object model virtual
// function table for the IADsClass interface.
type IADsClassVtbl struct {
	IADsVtbl
	PrimaryInterface       uintptr
	CLSID                  uintptr
	SetCLSID               uintptr
	OID                    uintptr
	SetOID                 uintptr
	Abstract               uintptr
	SetAbstract            uintptr
	Auxilary               uintptr
	SetAuxilary            uintptr
	MandatoryProperties    uintptr
	SetMandatoryProperties uintptr
	OptionalProperties     uintptr
	SetOptionalProperties  uintptr
	NamingProperties       uintptr
	SetNamingProperties    uintptr
	DerivedFrom            uintptr
	SetDerivedFrom         uintptr
	AuxDerivedFrom         uintptr
	SetAuxDerivedFrom      uintptr
	PossibleSuperiors      uintptr
	SetPossibleSuperiors   uintptr
	Containment            uintptr
	SetContainment         uintptr
	Container              uintptr
	SetContainer           uintptr
	HelpFileName           uintptr
	SetHelpFileName        uintptr
	HelpFileContext        uintptr
	SetHelpFileContext     uintptr
	Qualifiers             uintptr
}

// IADsClass represents the component object model interface for
// active directory classes.
type IADsClass struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// class.
func (v *IADsClass) VTable() *IADsClassVtbl {
	return (*IADsClassVtbl)(unsafe.Pointer(v.RawVTable))
}
