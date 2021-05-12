// +build windows

package api

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/google/uuid"
	"github.com/scjalliance/comutil"
	"github.com/go-adsi/adsi/comiid"
)

// NewIADsContainer returns a new instance of the IADsContainer
// component object model interface.
func NewIADsContainer(server string, clsid uuid.UUID) (*IADsContainer, error) {
	p, err := comutil.CreateRemoteObject(server, clsid, comiid.IADsContainer)
	return (*IADsContainer)(unsafe.Pointer(p)), err
}

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

// Filter retrieves the filter for the container.
func (v *IADsContainer) Filter() (variant *ole.VARIANT, err error) {
	variant = new(ole.VARIANT)
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Filter),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(variant)),
		0)
	if hr != 0 {
		return nil, convertHresultToError(hr)
	}
	return
}

// SetFilter sets the filter for the container.
func (v *IADsContainer) SetFilter(variant *ole.VARIANT) (err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().SetFilter),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(variant)),
		0)
	if hr != 0 {
		return convertHresultToError(hr)
	}
	return nil
}

// GetObject returns a descendant object with the given class and relative
// name.
//
// If a class is not provided then the first item matching the relative name
// will be returned regardless of its class.
func (v *IADsContainer) GetObject(class, name string) (obj *ole.IDispatch, err error) {
	var bclass, bname *int16

	if len(class) > 0 {
		bclass = ole.SysAllocStringLen(class)
		if bclass == nil {
			return nil, ole.NewError(ole.E_OUTOFMEMORY)
		}
		defer ole.SysFreeString(bclass)
	}

	bname = ole.SysAllocStringLen(name)
	if bname == nil {
		return nil, ole.NewError(ole.E_OUTOFMEMORY)
	}
	defer ole.SysFreeString(bname)

	hr, _, _ := syscall.Syscall6(
		uintptr(v.VTable().GetObject),
		4,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(bclass)),
		uintptr(unsafe.Pointer(bname)),
		uintptr(unsafe.Pointer(&obj)),
		0,
		0)
	if hr != 0 {
		return nil, convertHresultToError(hr)
	}
	return
}
