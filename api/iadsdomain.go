package api

import "unsafe"

// IADsDomainVtbl represents the component object model virtual
// function table for the IADsDomain interface.
type IADsDomainVtbl struct {
	IADsVtbl
	IsWorkgroup                   uintptr
	MinPasswordLength             uintptr
	SetMinPasswordLength          uintptr
	MinPasswordAge                uintptr
	SetMinPasswordAge             uintptr
	MaxPasswordAge                uintptr
	SetMaxPasswordAge             uintptr
	MaxBadPasswordsAllowed        uintptr
	SetMaxBadPasswordsAllowed     uintptr
	PasswordHistoryLength         uintptr
	SetPasswordHistoryLength      uintptr
	PasswordAttributes            uintptr
	SetPasswordAttributes         uintptr
	AutoUnlockInterval            uintptr
	SetAutoUnlockInterval         uintptr
	LockoutObservationInterval    uintptr
	SetLockoutObservationInterval uintptr
}

// IADsDomain represents the component object model interface for
// active directory domains.
type IADsDomain struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// domain.
func (v *IADsDomain) VTable() *IADsDomainVtbl {
	return (*IADsDomainVtbl)(unsafe.Pointer(v.RawVTable))
}
