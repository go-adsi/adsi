//go:build !windows
// +build !windows

package api

import (
	"github.com/go-ole/go-ole"
)

// NewIADsNameTranslate returns an IADsNameTranslate that manages the given COM interface.
func NewIADsNameTranslate(server string) (*IADsNameTranslate, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

// Get retrieves the name of a directory object in the specified format.
// The distinguished name must have been set in the appropriate format by the Set function.
func (v *IADsNameTranslate) Get(formatType uint32) (adsPath string, err error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}

// Init initializes a name translate object by binding to a specified directory server, domain,
// or global catalog, using the credentials of the current user.
func (v *IADsNameTranslate) Init(adsPath string, initType uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

// Set directs the directory service to set up a specified object for name translation.
func (v *IADsNameTranslate) Set(adsPath string, setType uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
