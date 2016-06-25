// +build !windows

package api

import "github.com/go-ole/go-ole"

// ComputerID retrieves the globally unique identifier of the computer.
func (v *IADsComputer) ComputerID() (id string, err error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}

// Site retrieves the site of the computer.
func (v *IADsComputer) Site() (site string, err error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}

// OperatingSystem retrieves the operating system of the computer.
func (v *IADsComputer) OperatingSystem() (os string, err error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}
