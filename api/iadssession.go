package api

import "unsafe"

// IADsSessionVtbl represents the component object model virtual
// function table for the IADsSession interface.
type IADsSessionVtbl struct {
	IADsVtbl
	User         uintptr
	UserPath     uintptr
	Computer     uintptr
	ComputerPath uintptr
	ConnectTime  uintptr
	IdleTime     uintptr
}

// IADsSession represents the component object model interface for
// active directory sessions.
type IADsSession struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// session.
func (v *IADsSession) VTable() *IADsSessionVtbl {
	return (*IADsSessionVtbl)(unsafe.Pointer(v.RawVTable))
}
