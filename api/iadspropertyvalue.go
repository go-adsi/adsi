package api

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

// IADsPropertyValueVtbl represents the component object model virtual
// function table for the IADsPropertyValue interface.
type IADsPropertyValueVtbl struct {
	ole.IDispatchVtbl
	Clear                 uintptr
	Type                  uintptr
	SetType               uintptr
	DistinguishedName     uintptr
	SetDistinguishedName  uintptr
	CaseExactString       uintptr
	SetCaseExactString    uintptr
	IgnoreCaseString      uintptr
	SetIgnoreCaseString   uintptr
	PrintableString       uintptr
	SetPrintableString    uintptr
	NumericString         uintptr
	SetNumericString      uintptr
	Boolean               uintptr
	SetBoolean            uintptr
	Integer               uintptr
	SetInteger            uintptr
	OctetString           uintptr
	SetOctetString        uintptr
	SecurityDescriptor    uintptr
	SetSecurityDescriptor uintptr
	LargeInteger          uintptr
	SetLargeInteger       uintptr
	UTCTime               uintptr
	SetUTCTime            uintptr
}

// IADsPropertyValue represents the component object model interface for
// property values.
type IADsPropertyValue struct {
	ole.IDispatch
}

// VTable returns the component object model virtual function table for the
// property value.
func (v *IADsPropertyValue) VTable() *IADsPropertyValueVtbl {
	return (*IADsPropertyValueVtbl)(unsafe.Pointer(v.RawVTable))
}
