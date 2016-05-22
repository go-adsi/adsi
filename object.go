package adsi

import (
	"sync"
	"unsafe"

	"gopkg.in/adsi.v0/api"
	"gopkg.in/adsi.v0/comshim"
)

// ADSI Objects of LDAP:  https://msdn.microsoft.com/library/aa772208
// ADSI Objects of WinNT: https://msdn.microsoft.com/library/aa772211

// Object provides access to Active Directory objects.
type Object struct {
	m     sync.RWMutex
	iface *api.IADs
}

// NewObject returns an object that manages the given COM interface.
func NewObject(iface *api.IADs) *Object {
	comshim.Add(1)
	return &Object{iface: iface}
}

/*
// See: https://msdn.microsoft.com/library/aa772184
func GetObject(path string, iid ole.GUID) (obj *Object, err error) {
	// TODO: Implement this
	return
}

// See: https://msdn.microsoft.com/library/aa772184
func RemoteObject(server, path string) (obj *Object, err error) {
	// TODO: Implement this
	return
}
*/

func (o *Object) closed() bool {
	return (o.iface == nil)
}

// Close will release resources consumed by the object. It should be
// called when the object is no longer needed.
func (o *Object) Close() {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return
	}
	defer comshim.Done()
	run(func() error {
		o.iface.Release()
		return nil
	})
	// FIXME: What happens if the run returns an error?
	o.iface = nil
}

// Name retrieves the name of the object.
func (o *Object) Name() (name string, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return "", ErrClosed
	}
	err = run(func() error {
		name, err = o.iface.Name()
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// Class retrieves the class of the object.
func (o *Object) Class() (class string, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return "", ErrClosed
	}
	err = run(func() error {
		class, err = o.iface.Class()
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// GUID retrieves the globally unique identifier of the object.
func (o *Object) GUID() (guid string, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return "", ErrClosed
	}
	err = run(func() error {
		guid, err = o.iface.GUID()
		if err != nil {
			return err
		}
		// TODO: Cast guid to a proper ole.GUID type
		return nil
	})
	return
}

// Path retrieves the fully qualified path of the object.
func (o *Object) Path() (path string, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return "", ErrClosed
	}
	err = run(func() error {
		path, err = o.iface.AdsPath()
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// Parent retrieves the fully qualified path of the object's parent.
func (o *Object) Parent() (path string, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return "", ErrClosed
	}
	err = run(func() error {
		path, err = o.iface.Parent()
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// Schema retrieves the fully qualified path of the object's schema class
// object.
func (o *Object) Schema() (path string, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return "", ErrClosed
	}
	err = run(func() error {
		path, err = o.iface.Schema()
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// ToContainer attempts to acquire a container interface for the object.
func (o *Object) ToContainer() (c *Container, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		idispatch, err := o.iface.QueryInterface(api.IID_IADsContainer)
		if err != nil {
			return err
		}
		iface := (*api.IADsContainer)(unsafe.Pointer(idispatch))
		c = NewContainer(iface)
		return nil
	})
	return
}
