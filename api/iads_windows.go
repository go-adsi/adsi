// +build windows

package api

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

var (
	ErrUnsupportedType = errors.New("unsupported type")
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

// Put sets the values of an attribute in the ADSI attribute cache. The value
// must be commited with SetInfo to be made persistent.
func (v *IADs) Put(name string, val interface{}) error {
	bname := ole.SysAllocStringLen(name)
	if bname == nil {
		return ole.NewError(ole.E_OUTOFMEMORY)
	}
	defer ole.SysFreeString(bname)
	var prop ole.VARIANT
	switch vt := val.(type) {
	case string:
		prop = ole.NewVariant(ole.VT_BSTR, int64(uintptr(unsafe.Pointer(ole.SysAllocStringLen(vt)))))
	case int:
		prop = ole.NewVariant(ole.VT_I4, int64(vt))
	default:
		return ErrUnsupportedType
	}

	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Put),
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(bname)),
		uintptr(unsafe.Pointer(&prop)))
	if hr != 0 {
		defer ole.VariantClear(&prop)
		return convertHresultToError(hr)
	}
	return nil
}

// SetInfo saves the cached property values of the ADSI object to the underlying directory store.
func (v *IADs) SetInfo() error {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().SetInfo),
		1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
	if hr != 0 {
		return convertHresultToError(hr)
	}
	return nil
}
