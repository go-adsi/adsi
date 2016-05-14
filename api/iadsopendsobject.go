package api

import (
	"errors"
	"unsafe"

	"github.com/go-ole/go-ole"
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
func NewIADsOpenDSObject(server string) (ds *IADsOpenDSObject, err error) {
	// FIXME: Right now the caller is expected to wrap this call between
	//        CoInitialize and CoUninitialize, but maybe this function should call
	//        them itself?

	var bserver *int16
	if len(server) > 0 {
		bserver = ole.SysAllocStringLen(server)
		if bserver == nil {
			return nil, ole.NewError(ole.E_OUTOFMEMORY)
		}
		defer ole.SysFreeString(bserver)
	}

	serverInfo := &CoServerInfo{
		Name: bserver,
	}

	var context uint
	if server == "" {
		context = ole.CLSCTX_SERVER
	} else {
		context = ole.CLSCTX_REMOTE_SERVER
	}

	results := make([]MultiQI, 0, 1)
	results = append(results, MultiQI{IID: IID_IADsOpenDSObject})

	// TODO: Figure out how to handle varying namespace providers
	err = CreateInstanceEx(CLSID_LDAPNamespace, context, serverInfo, results)
	//err = CreateInstanceEx(CLSID_WinNTNamespace, context, serverInfo, results)
	if err != nil {
		return nil, err
	}

	ds = (*IADsOpenDSObject)(unsafe.Pointer(results[0].Interface))
	if ds == nil {
		err = errors.New("Unable to create IADsOpenDSObject instance.")
	}
	return
}
