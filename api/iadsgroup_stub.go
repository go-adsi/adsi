// +build !windows

package api

import "github.com/go-ole/go-ole"

// Add adds an ADSI object to an existing group.
func (v *IADsGroup) Add(member string) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

// Description retrieves the description of the group.
func (v *IADsGroup) Description() (desc string, err error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}

// Members retrieves an IADsMembers interface that provides access to the
// membership of the group.
func (v *IADsGroup) Members() (members *IADsMembers, err error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

// Remove removes the specified user object from this group. The operation
// does not remove the group object itself even when there is no member remaining in the group.
func (v *IADsGroup) Remove(member string) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
