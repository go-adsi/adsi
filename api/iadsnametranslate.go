package api

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IADsNameTranslateVtbl struct {
	ole.IDispatchVtbl
	ChaseReferral uintptr
	Init          uintptr
	InitEx        uintptr
	Set           uintptr
	Get           uintptr
	SetEx         uintptr
	GetEx         uintptr
}

type IADsNameTranslate struct {
	ole.IDispatch
}

func (v *IADsNameTranslate) VTable() *IADsNameTranslateVtbl {
	return (*IADsNameTranslateVtbl)(unsafe.Pointer(v.RawVTable))
}
