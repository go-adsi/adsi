//go:build !windows
// +build !windows

package api

import "github.com/go-ole/go-ole"

// AccountDisabled retrieves the disablement status of a user account.
func (v *IADsUser) AccountDisabled() (disabled bool, err error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

// SetAccountDisabled sets an account as disabled.
func (v *IADsUser) SetAccountDisabled(disabled bool) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

// FullName returns the user's FullName property.
func (v *IADsUser) FullName() (name string, err error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}
