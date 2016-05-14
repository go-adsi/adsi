package adsi

import (
	"sync"
	"unsafe"

	"github.com/go-ole/go-ole"
	"gopkg.in/adsi.v0/api"
)

// Container provides access to Active Directory container objects.
type Container struct {
	m     sync.RWMutex
	iface *api.IADsContainer
}

func (c *Container) closed() bool {
	return (c.iface == nil)
}

// Close will release resources consumed by the container. It should be
// called when the container is no longer needed.
func (c *Container) Close() {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return
	}
	run(func() error {
		c.iface.Release()
		return nil
	})
	// FIXME: What happens if the run returns an error?
	c.iface = nil
}

// Children returns an object iterator that provides access to the immediate
// children of the container.
func (c *Container) Children() (iter *ObjectIter, err error) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		iunknown, err := c.iface.NewEnum()
		if err != nil {
			return err
		}
		defer iunknown.Release()
		idispatch, err := iunknown.QueryInterface(ole.IID_IEnumVariant)
		if err != nil {
			return err
		}
		iface := (*ole.IEnumVARIANT)(unsafe.Pointer(idispatch))
		iter = NewObjectIter(iface)
		return nil
	})
	return

}

// ObjectIter provides an iterator for a set of objects.
type ObjectIter struct {
	iface *ole.IEnumVARIANT
}

// NewObjectIter returns an object iterator that provides access to the objects
// contained in the given enumerator.
func NewObjectIter(enumerator *ole.IEnumVARIANT) *ObjectIter {
	// TODO: Call ADsBuildEnumerator here?
	return &ObjectIter{iface: enumerator}
}

// TODO: Add iterator functions
