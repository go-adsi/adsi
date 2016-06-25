package adsi

import (
	"strings"
	"sync"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/scjalliance/comshim"
	"gopkg.in/adsi.v0/adspath"
	"gopkg.in/adsi.v0/api"
)

type namespace struct {
	Name    string
	ClassID *ole.GUID
	Iface   *api.IADsOpenDSObject
	Err     error
}

// Client provides access to Active Directory Service Interfaces for
// any namespace supported by a local or remote COM server.
type Client struct {
	m sync.RWMutex
	n []namespace
}

// NewClient creates a new ADSI client. When done with a client it should be
// closed with a call to Close(). If NewClient is successful it will return a
// client and error will be nil, otherwise the returned client will be nil and
// error will be non-nil.
func NewClient() (*Client, error) {
	return NewRemoteClient("")
}

// NewRemoteClient creates a new ADSI client on a remote server. When done with
// a client it should be closed with a call to Close(). If NewClient is
// successful it will return a client and error will be nil, otherwise the
// returned client will be nil and error will be non-nil.
//
// If no server is provided a local client is created instead and the
// resulting behavior is identical to NewClient.
func NewRemoteClient(server string) (*Client, error) {
	comshim.Add(1)
	c := &Client{}
	err := run(func() error {
		return c.init(server)
	})
	if err != nil {
		comshim.Done()
		return nil, err
	}
	// TODO: Add finalizer for ds?
	return c, nil
}

func (c *Client) init(server string) (err error) {
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

	c.n = make([]namespace, 0, 12)

	for child, iterErr := iter.Next(); iterErr == nil; child, iterErr = iter.Next() {
		defer child.Close()

		// Add the entry and whip up a pointer to it
		c.n = append(c.n, namespace{})
		item := &c.n[len(c.n)-1]

		// Name
		item.Name, item.Err = child.Name()
		if item.Err != nil {
			continue
		}
		item.Name = strings.TrimRight(item.Name, ":")

		// GUID
		item.ClassID, item.Err = child.GUID()
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

func (c *Client) closed() bool {
	return (c.n == nil)
}

// Close will release resources consumed by the client. It should be called
// when the client is no longer needed.
func (c *Client) Close() {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return
	}
	defer comshim.Done()
	run(func() error {
		for i := 0; i < len(c.n); i++ {
			if c.n[i].Iface != nil {
				c.n[i].Iface.Release()
			}
		}
		return nil
	})
	c.n = nil
}

// Open opens an ADSI object with the given path. When provided, the
// username and password are used to establish a security context for the
// connection. When credentials are not provided the existing security
// context of the application is used instead.
//
// Open returns the ADSI object as an Object type, which provides
// an idiomatic go wrapper around the underlying component object model
// IADs interface.
//
// Open calls QueryInterface internally to acquire an implementation of
// the IADs interface that is needed by the Object type. If the returned
// directory object does not implement the IADs interface an error is
// returned.
//
// The returned object consumes resources until it is closed. It is the
// caller's responsibilty to call Close on the returned object when it is no
// longer needed.
func (c *Client) Open(path, user, password string, flags uint32) (obj *Object, err error) {
	idispatch, err := c.OpenInterface(path, user, password, flags, api.IID_IADs)
	if err != nil {
		return nil, err
	}
	iface := (*api.IADs)(unsafe.Pointer(idispatch))
	obj = NewObject(iface)
	return
}

// OpenDispatch opens an ADSI object with the given path. When provided, the
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
func (c *Client) OpenDispatch(path, user, password string, flags uint32) (obj *ole.IDispatch, err error) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		obj, err = c.open(path, user, password, flags)
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
func (c *Client) OpenInterface(path, user, password string, flags uint32, iid *ole.GUID) (obj *ole.IDispatch, err error) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		idispatch, err := c.open(path, user, password, flags)
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

func (c *Client) open(path, user, password string, flags uint32) (obj *ole.IDispatch, err error) {
	p, err := adspath.Parse(path)
	if err != nil {
		return
	}

	ns := c.namespace(p.Scheme)
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
func (c *Client) namespace(name string) *namespace {
	for i := 0; i < len(c.n); i++ {
		if c.n[i].Name == name {
			return &c.n[i]
		}
	}
	return nil
}
