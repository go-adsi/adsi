// +build windows

package api

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// Name retrieves the name of the object.
func (v *IADs) Name() (name string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Name),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr == 0 {
		name = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	} else {
		return "", convertHresultToError(hr)
	}
	return
}

// Class retrieves the class of the object.
func (v *IADs) Class() (class string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Class),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr == 0 {
		class = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	} else {
		return "", convertHresultToError(hr)
	}
	return
}

// GUID retrieves the GUID of the object as a string.
func (v *IADs) GUID() (guid string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().GUID),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr == 0 {
		guid = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	} else {
		return "", convertHresultToError(hr)
	}
	return
}

// AdsPath retrieves the fully qualified path of the object.
func (v *IADs) AdsPath() (path string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().AdsPath),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr == 0 {
		path = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	} else {
		return "", convertHresultToError(hr)
	}
	return
}

// Parent retrieves the fully qualified path of the object's parent.
func (v *IADs) Parent() (path string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Parent),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr == 0 {
		path = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	} else {
		return "", convertHresultToError(hr)
	}
	return
}

// Schema retrieves the fully qualified path of the object's schema class
// object.
func (v *IADs) Schema() (path string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Schema),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr == 0 {
		path = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	} else {
		return "", convertHresultToError(hr)
	}
	return
}

// Get retrieves a property with the given name. If the property holds a single
// item, a VARIANT for that item is returned. If the property holds multiple
// items, a VARIANT array is returned containing the items, with each value
// being a VARIANT itself.
func (v *IADs) Get(name string) (prop *ole.VARIANT, err error) {
	bname := ole.SysAllocStringLen(name)
	if bname == nil {
		return nil, ole.NewError(ole.E_OUTOFMEMORY)
	}
	defer ole.SysFreeString(bname)
	prop = new(ole.VARIANT)
	ole.VariantInit(prop)
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Get),
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(bname)),
		uintptr(unsafe.Pointer(prop)))
	if hr != 0 {
		defer ole.VariantClear(prop)
		return nil, convertHresultToError(hr)
	}
	return
}

// GetEx retrieves a property with the given name. The property is returned as
// a VARIANT array type, with each value being a VARIANT itself. Unlike the
// Get function, if the property holds a single item, it is returned as a
// VARIANT array with one member.
func (v *IADs) GetEx(name string) (prop *ole.VARIANT, err error) {
	bname := ole.SysAllocStringLen(name)
	if bname == nil {
		return nil, ole.NewError(ole.E_OUTOFMEMORY)
	}
	defer ole.SysFreeString(bname)
	prop = new(ole.VARIANT)
	ole.VariantInit(prop)
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().GetEx),
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(bname)),
		uintptr(unsafe.Pointer(prop)))
	if hr != 0 {
		defer prop.Clear()
		return nil, convertHresultToError(hr)
	}
	return
}

// GetInfoEx loads the given set of property names into the cache. The given
// variant must be a safe array of null-terminated unicode strings.
func (v *IADs) GetInfoEx(variant *ole.VARIANT) (err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().GetInfoEx),
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(variant)),
		0) // This is a reserved value: it must be included and must be zero
	if hr != 0 {
		return convertHresultToError(hr)
	}
	return nil
}
