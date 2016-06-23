package adsi

import (
	"strings"
	"sync"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/scjalliance/comshim"
	"github.com/scjalliance/comutil"
	"gopkg.in/adsi.v0/adspath"
	"gopkg.in/adsi.v0/api"
)

type namespace struct {
	Name    string
	ClassID *ole.GUID
	Iface   *api.IADsOpenDSObject
	Err     error
}

// DirectoryService provides access to Active Directory Service Interfaces for
// a namespace.
type DirectoryService struct {
	m sync.RWMutex
	n []namespace
}

// NewDirectoryService creates a new directory service. When done with a
// directory service it should be closed with a call to Close(). If New is
// successful it will return a directory service and error will be nil,
// otherwise the returned directory service will be nil and error will be
// non-nil.
func NewDirectoryService(server string) (*DirectoryService, error) {
	comshim.Add(1)
	ds := &DirectoryService{}
	err := run(func() error {
		return ds.init(server)
	})
	if err != nil {
		comshim.Done()
		return nil, err
	}
	// TODO: Add finalizer for ds?
	return ds, nil
}

func (ds *DirectoryService) init(server string) (err error) {
	// Acquiring a container for the CLSID_ADsNamespaces class gives us access to
	// an enumeration of all of the available namespaces.
	iface, err := api.NewIADsContainer(server, api.CLSID_ADsNamespaces)
	if err != nil {
		return err
	}

	root := NewContainer(iface)
	defer root.Close()

	iter, err := root.Children()
	if err != nil {
		return err
	}
	defer iter.Close()

	ds.n = make([]namespace, 0, 12)

	for child, iterErr := iter.Next(); iterErr == nil; child, iterErr = iter.Next() {
		defer child.Close()

		// Add the entry and whip up a pointer to it
		ds.n = append(ds.n, namespace{})
		item := &ds.n[len(ds.n)-1]

		// Name
		item.Name, item.Err = child.Name()
		if item.Err != nil {
			continue
		}
		item.Name = strings.TrimRight(item.Name, ":")

		// GUID
		var guid string
		guid, item.Err = child.GUID()
		if item.Err != nil {
			continue
		}

		item.ClassID, item.Err = comutil.IIDFromString(guid)
		if item.Err != nil {
			continue
		}

		// Interface
		var idisp *ole.IDispatch
		idisp, item.Err = child.iface.QueryInterface(api.IID_IADsOpenDSObject)
		if item.Err != nil {
			continue
		}
		item.Iface = (*api.IADsOpenDSObject)(unsafe.Pointer(idisp))
	}

	// TODO: Check the value of iterErr to see if it returned something other than
	//       io.EOF.

	return
}

func (ds *DirectoryService) closed() bool {
	return (ds.n == nil)
}

// Close will release resources consumed by the directory service. It should be
// called when the directory service is no longer needed.
func (ds *DirectoryService) Close() {
	ds.m.Lock()
	defer ds.m.Unlock()
	if ds.closed() {
		return
	}
	defer comshim.Done()
	run(func() error {
		for i := 0; i < len(ds.n); i++ {
			if ds.n[i].Iface != nil {
				ds.n[i].Iface.Release()
			}
		}
		return nil
	})
	ds.n = nil
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
		obj, err = ds.open(path, user, password, flags)
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
		idispatch, err := ds.open(path, user, password, flags)
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
	obj = NewObject(iface)
	return
}

func (ds *DirectoryService) open(path, user, password string, flags uint32) (obj *ole.IDispatch, err error) {
	p, err := adspath.Parse(path)
	if err != nil {
		return
	}

	ns := ds.namespace(p.Scheme)
	if ns == nil {
		return nil, api.ErrInvalidNamespace
	}
	if ns.Err != nil {
		return nil, ns.Err
	}

	obj, err = ns.Iface.OpenDSObject(path, user, password, flags)
	return
}

// namespace returns information about the namespace with the given name. If
// no namespace has been registered with that name then nil is returend.
//
// The name matching is case-sensitive.
func (ds *DirectoryService) namespace(name string) *namespace {
	for i := 0; i < len(ds.n); i++ {
		if ds.n[i].Name == name {
			return &ds.n[i]
		}
	}
	return nil
}
