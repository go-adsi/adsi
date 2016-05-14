// +build !windows

package api

import "github.com/go-ole/go-ole"

// OpenDSObject retrieves a directory services object for the given path.
//
// See: https://msdn.microsoft.com/library/aa706065
func (v *IADsOpenDSObject) OpenDSObject(path, user, password string, flags uint32) (obj *ole.IDispatch, err error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}
