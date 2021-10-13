//go:build !windows
// +build !windows

package api

import (
	"github.com/go-ole/go-ole"
	"github.com/google/uuid"
)

// Name retrieves the name of the object.
func (v *IADs) Name() (name string, err error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}

// Class retrieves the class of the object.
func (v *IADs) Class() (class string, err error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}

// GUID retrieves the GUID of the object as a string.
func (v *IADs) GUID() (guid uuid.UUID, err error) {
	return uuid.New(), ole.NewError(ole.E_NOTIMPL)
}

// AdsPath retrieves the fully qualified path of the object.
func (v *IADs) AdsPath() (path string, err error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}

// Parent retrieves the fully qualified path of the object's parent.
func (v *IADs) Parent() (path string, err error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}

// Schema retrieves the fully qualified path of the object's schema class
// object.
func (v *IADs) Schema() (path string, err error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}

// Get retrieves a property with the given name. If the property holds a single
// item, a VARIANT for that item is returned. If the property holds multiple
// items, a VARIANT array is returned containing the items, with each value
// being a VARIANT itself.
func (v *IADs) Get(name string) (prop *ole.VARIANT, err error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

// GetEx retrieves a property with the given name. The property is returned as
// a VARIANT array type, with each value being a VARIANT itself. Unlike the
// Get function, if the property holds a single item, it is returned as a
// VARIANT array with one member.
func (v *IADs) GetEx(name string) (prop *ole.VARIANT, err error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

// PutInt sets the values of an int attribute in the ADSI attribute
// cache. The value must be commited with SetInfo to be made persistent.
func (v *IADs) PutInt(name string, val int) error {
	return ole.NewError(ole.E_NOTIMPL)
}

// PutString sets the values of a string attribute in the ADSI attribute
// cache. The value must be commited with SetInfo to be made persistent.
func (v *IADs) PutString(name string, val string) error {
	return ole.NewError(ole.E_NOTIMPL)
}

// SetInfo saves the cached property values of the ADSI object to the underlying directory store.
func (v *IADs) SetInfo() error {
	return ole.NewError(ole.E_NOTIMPL)
}
