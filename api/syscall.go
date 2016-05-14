package api

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// See: https://msdn.microsoft.com/en-us/library/ms680701
// See: https://msdn.microsoft.com/en-us/library/ms687322
// See: https://msdn.microsoft.com/en-us/library/ms693716

var (
	modole32, _ = syscall.LoadDLL("ole32.dll")
)

var (
	procCoCreateInstanceEx, _ = modole32.FindProc("CoCreateInstanceEx")
)

// CreateInstanceEx supports remote creation of several interfaces within one
// class.
func CreateInstanceEx(clsid *ole.GUID, context uint, serverInfo *CoServerInfo, results []MultiQI) (err error) {
	var _p0 *MultiQI
	if len(results) > 0 {
		_p0 = &results[0]
	}
	hr, _, _ := procCoCreateInstanceEx.Call(
		uintptr(unsafe.Pointer(clsid)),
		0,
		uintptr(context),
		uintptr(unsafe.Pointer(serverInfo)),
		uintptr(len(results)),
		uintptr(unsafe.Pointer(_p0)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
