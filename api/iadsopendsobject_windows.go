// +build windows

package api

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// OpenDSObject retrieves a directory services object for the given path.
//
// See: https://msdn.microsoft.com/library/aa706065
func (v *IADsOpenDSObject) OpenDSObject(path, user, password string, flags uint32) (obj *ole.IDispatch, err error) {
	// FIXME: Right now the caller is expected to wrap this call between
	//        CoInitialize and CoUninitialize, but maybe this function should call
	//        them itself?

	var bpath, buser, bpassword *int16

	if len(path) > 0 {
		bpath = ole.SysAllocStringLen(path)
		if bpath == nil {
			return nil, ole.NewError(ole.E_OUTOFMEMORY)
		}
		defer ole.SysFreeString(bpath)
	}

	if len(user) > 0 {
		buser = ole.SysAllocStringLen(user)
		if buser == nil {
			return nil, ole.NewError(ole.E_OUTOFMEMORY)
		}
		defer ole.SysFreeString(buser)
	}

	if len(password) > 0 {
		bpassword = ole.SysAllocStringLen(password)
		if bpassword == nil {
			return nil, ole.NewError(ole.E_OUTOFMEMORY)
		}
		defer ole.SysFreeString(bpassword)
	}

	hr, _, _ := syscall.Syscall6(
		uintptr(v.VTable().OpenDSObject),
		6,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(bpath)),
		uintptr(unsafe.Pointer(buser)),
		uintptr(unsafe.Pointer(bpassword)),
		uintptr(flags),
		uintptr(unsafe.Pointer(&obj)))
	if hr != 0 {
		return nil, convertHresultToError(hr)
	}
	return
}
