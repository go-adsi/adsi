package api

import (
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/google/uuid"
	"github.com/scjalliance/comutil"
	"github.com/go-adsi/adsi/v2/comiid"
)

// IADsOpenDSObjectVtbl represents the component object model virtual
// function table for the IADsOpenDSObject interface.
type IADsOpenDSObjectVtbl struct {
	ole.IDispatchVtbl
	OpenDSObject uintptr
}

// IADsOpenDSObject represents the component object model interface for
// directory services.
type IADsOpenDSObject struct {
	ole.IDispatch
}

// VTable returns the component object model virtual function table for the
// directory service.
func (v *IADsOpenDSObject) VTable() *IADsOpenDSObjectVtbl {
	return (*IADsOpenDSObjectVtbl)(unsafe.Pointer(v.RawVTable))
}

// NewIADsOpenDSObject returns a new instance of the IADsOpenDSObject
// component object model interface.
func NewIADsOpenDSObject(server string, clsid uuid.UUID) (ds *IADsOpenDSObject, err error) {
	p, err := comutil.CreateRemoteObject(server, clsid, comiid.IADsOpenDSObject)
	return (*IADsOpenDSObject)(unsafe.Pointer(p)), err
}
