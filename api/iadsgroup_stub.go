// +build !windows

package api

import "github.com/go-ole/go-ole"

// Description retrieves the description of the group.
func (v *IADsGroup) Description() (desc string, err error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}

// Members retrieves an IADsMembers interface that provides access to the
// membership of the group.
func (v *IADsGroup) Members() (members *IADsMembers, err error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}
