// +build windows

package api

import (
	"syscall"
	"unsafe"
)

// HighPart retrieves the upper 32 bits of the 64 bit value.
func (v *IADsLargeInteger) HighPart() (upper int32, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().HighPart),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&upper)),
		0)
	if hr != 0 {
		return 0, convertHresultToError(hr)
	}
	return
}

// LowPart retrieves the lower 32 bits of the 64 bit value.
func (v *IADsLargeInteger) LowPart() (lower int32, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().LowPart),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&lower)),
		0)
	if hr != 0 {
		return 0, convertHresultToError(hr)
	}
	return
}

// Value retrieves the 64 bit value.
func (v *IADsLargeInteger) Value() (value int64, err error) {
	var upper, lower int32
	upper, err = v.HighPart()
	if err != nil {
		return
	}
	lower, err = v.LowPart()
	if err != nil {
		return
	}
	return (int64(uint32(upper)) << 32) | int64(uint32(lower)), nil
}
