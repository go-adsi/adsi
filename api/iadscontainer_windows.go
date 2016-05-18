// +build windows

package api

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// NewEnum retrieves an enumerator interface that provides access to the objects
// within the container.
//
// See https://msdn.microsoft.com/library/aa705990
func (v *IADsContainer) NewEnum() (enum *ole.IUnknown, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().NewEnum),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&enum)),
		0)
	if hr != 0 {
		return nil, convertHresultToError(hr)
	}
	return
}
