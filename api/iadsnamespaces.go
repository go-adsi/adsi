package api

import "unsafe"

// IADsNamespacesVtbl represents the component object model virtual
// function table for the IADsNamespaces interface.
type IADsNamespacesVtbl struct {
	IADsVtbl
	GetDefaultContainer uintptr
	SetDefaultContainer uintptr
}

// IADsNamespaces represents the component object model interface for
// the active directory namespace registry.
type IADsNamespaces struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// namespace registry.
func (v *IADsNamespaces) VTable() *IADsNamespacesVtbl {
	return (*IADsNamespacesVtbl)(unsafe.Pointer(v.RawVTable))
}
