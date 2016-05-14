package adsi

import (
	"sync"
	"unsafe"

	"github.com/go-ole/go-ole"
	"gopkg.in/adsi.v0/api"
)

// DirectoryService provides access to Active Directory Service Interfaces for
// a namespace.
type DirectoryService struct {
	m     sync.RWMutex
	iface *api.IADsOpenDSObject
}

// NewDirectoryService creates a new directory service. When done with a
// directory service it should be closed with a call to Close(). If New is
// successful it will return a directory service and error will be nil,
// otherwise the returned directory service will be nil and error will be
// non-nil.
func NewDirectoryService(server string) (*DirectoryService, error) {
	ds := &DirectoryService{}
	err := run(func() error {
		return ds.init(server)
	})
	if err != nil {
		return nil, err
	}
	// TODO: Add finalizer for ds?
	return ds, nil
}

func (ds *DirectoryService) init(server string) (err error) {
	ds.iface, err = api.NewIADsOpenDSObject(server)
	return
}

func (ds *DirectoryService) closed() bool {
	return (ds.iface == nil)
}

// Close will release resources consumed by the directory service. It should be
// called when the directory service is no longer needed.
func (ds *DirectoryService) Close() {
	ds.m.Lock()
	defer ds.m.Unlock()
	if ds.closed() {
		return
	}
	run(func() error {
		ds.iface.Release()
		return nil
	})
	// FIXME: What happens if the run returns an error?
	ds.iface = nil
}

// Open opens a directory object with the given path. When provided, the
// username and password are used to establish a security context for the
// connection. When credentials are not provided the existing security
// context of the application is used instead.
//
// Open returns a generic IDispatch interface for the object, which can be
// further interrogated to find out which component object model interfaces it
// implements.
//
// To return an object that has already been wrapped in the more convenient
// and safer Object type, use OpenObject instead.
//
// To open an object with a specific interface ID, use OpenInterface instead.
//
// The returned interface consumes resources until it is released. It is the
// caller's responsibilty to call Release on the returned object when it is no
// longer needed.
func (ds *DirectoryService) Open(path, user, password string, flags uint32) (obj *ole.IDispatch, err error) {
	ds.m.Lock()
	defer ds.m.Unlock()
	if ds.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		obj, err = ds.iface.OpenDSObject(path, user, password, flags)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// OpenInterface opens a directory object with the given path. When provided,
// the username and password are used to establish a security context for the
// connection. When credentials are not provided the existing security
// context of the application is used instead.
//
// OpenInterface calls QueryInterface internally to return a pointer to an
// object implementing the requested interface ID. If the returned object
// does not implement the requested interface an error is returned. The object
// is returned as a pointer to an IDispatch interface; it is expected that the
// caller will recast it as a pointer to the requested implementation.
//
// To return an object that has already been wrapped in the more convenient
// and safer Object type, use OpenObject instead.
//
// The returned interface consumes resources until it is released. It is the
// caller's responsibilty to call Release on the returned object when it is no
// longer needed.
func (ds *DirectoryService) OpenInterface(path, user, password string, flags uint32, iid *ole.GUID) (obj *ole.IDispatch, err error) {
	ds.m.Lock()
	defer ds.m.Unlock()
	if ds.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		idispatch, err := ds.iface.OpenDSObject(path, user, password, flags)
		if err != nil {
			return err
		}
		defer idispatch.Release()
		obj, err = idispatch.QueryInterface(iid)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// OpenObject opens a directory object with the given path. When provided,
// the username and password are used to establish a security context for the
// connection. When credentials are not provided the existing security
// context of the application is used instead.
//
// OpenObject returns the directory object as an Object type, which provides
// an idiomatic go wrapper around the underlying component object model
// interface.
//
// OpenObject calls QueryInterface internally to acquire an implementation of
// the IADs interface that is needed by the Object type. If the returned
// directory object does not implement the IADs interface an error is
// returned.
//
// The returned object consumes resources until it is closed. It is the
// caller's responsibilty to call Close on the returned object when it is no
// longer needed.
func (ds *DirectoryService) OpenObject(path, user, password string, flags uint32) (obj *Object, err error) {
	idispatch, err := ds.OpenInterface(path, user, password, flags, api.IID_IADs)
	if err != nil {
		return nil, err
	}
	iface := (*api.IADs)(unsafe.Pointer(idispatch))
	obj = &Object{iface: iface}
	return
}
