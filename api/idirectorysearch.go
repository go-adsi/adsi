package api

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

// IDirectorySearchVtbl represents the component object model virtual
// function table for the IDirectorySearch interface.
type IDirectorySearchVtbl struct {
	ole.IDispatchVtbl
	SetSearchPreferences uintptr
	ExecuteSearch        uintptr
	AbandonSearch        uintptr
	GetFirstRow          uintptr
	GetNextRow           uintptr
	GetPreviousRow       uintptr
	GetNextColumnName    uintptr
	GetColumn            uintptr
	FreeColumn           uintptr
	CloseSearchHandle    uintptr
}

// IDirectorySearch represents the component object model interface for
// conducting directory searches.
type IDirectorySearch struct {
	ole.IDispatch
}

// VTable returns the component object model virtual function table for the
// property value.
func (v *IDirectorySearch) VTable() *IDirectorySearchVtbl {
	return (*IDirectorySearchVtbl)(unsafe.Pointer(v.RawVTable))
}
