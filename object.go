package adsi

import (
	"sync"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/scjalliance/comshim"
	"github.com/scjalliance/comutil"
	"gopkg.in/adsi.v0/api"
)

// ADSI Objects of LDAP:  https://msdn.microsoft.com/library/aa772208
// ADSI Objects of WinNT: https://msdn.microsoft.com/library/aa772211

// Object provides access to Active Directory objects.
type Object struct {
	object
}

// NewObject returns an object that manages the given COM interface.
func NewObject(iface *api.IADs) *Object {
	comshim.Add(1)
	return &Object{object{iface: iface}}
}

type object struct {
	m     sync.RWMutex
	iface *api.IADs
}

func (o *object) closed() bool {
	return (o.iface == nil)
}

// Close will release resources consumed by the object. It should be
// called when the object is no longer needed.
func (o *object) Close() {
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
func (o *object) Name() (name string, err error) {
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
func (o *object) Class() (class string, err error) {
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
func (o *object) GUID() (guid *ole.GUID, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		var sguid string
		sguid, err = o.iface.GUID() // may return binary octet string in hexadecimal form
		if err != nil {
			return err
		}

		guid = ole.NewGUID(sguid)
		if guid == nil {
			return ErrInvalidGUID
		}

		if len(sguid) == 32 {
			// 32 character representations lack curly braces and dashes. When the
			// LDAP provider returns a GUID as a string in this form, it is the result
			// of taking the binary octet string of the GUID and converting it to
			// hexadecimal.
			//
			// Assuming that the original binary octet string was in the endianness
			// of the originating system, and that the originating system was
			// little-endian, we need to swap the bytes of the GUID we just parsed.
			guid.Data1 = reverseUint32(guid.Data1)
			guid.Data2 = reverseUint16(guid.Data2)
			guid.Data3 = reverseUint16(guid.Data3)
		}

		return nil
	})
	return
}

// Path retrieves the fully qualified path of the object.
func (o *object) Path() (path string, err error) {
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
func (o *object) Parent() (path string, err error) {
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
func (o *object) Schema() (path string, err error) {
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

// Attr attempts to retrieve the attribute with the given name and
// return its values as a slice of interfaces. Each value is an interface{}
// that holds a Go native type that is the best match for the underlying
// variant.
func (o *object) Attr(name string) (values []interface{}, err error) {
	variant, err := o.iface.GetEx(name)
	if err != nil {
		return
	}
	defer variant.Clear()

	array := variant.ToArray()
	if array == nil {
		return nil, ErrNonArrayAttribute
	}

	dims, _ := comutil.SafeArrayGetDim(array.Array)
	if dims != 1 {
		return nil, ErrMultiDimArrayAttribute
	}

	vt, err := array.GetType()
	if err != nil {
		return
	}
	if ole.VT(vt) != ole.VT_VARIANT {
		return nil, ErrNonVariantArrayAttribute
	}

	elems, err := array.TotalElements(0)
	if err != nil {
		return
	}

	for i := int64(0); i < elems; i++ {
		element := &ole.VARIANT{}
		ole.VariantInit(element)
		defer ole.VariantClear(element)
		err = comutil.SafeArrayGetElement(array.Array, i, unsafe.Pointer(element))
		if err != nil {
			return
		}
		values = append(values, element.Value())
	}

	return
}

// AttrStringSlice attempts to retrieve the attribute with the given name and
// return its values as a slice of strings.
//
// Any non-string values contained in the attribute will be ommitted.
func (o *object) AttrStringSlice(name string) (values []string, err error) {
	elements, err := o.Attr(name)
	if err != nil {
		return
	}
	for _, element := range elements {
		if s, ok := element.(string); ok {
			values = append(values, s)
		} else {
			// TODO: Consider returning error
		}
	}
	return
}

// AttrString attempts to retrieve the attribute with the given name and
// return its value as a string. If the attribute holds more than one value,
// only the first value is returned.
//
// Any non-string values contained in the attribute will be ignored.
func (o *object) AttrString(name string) (attr string, err error) {
	array, err := o.AttrStringSlice(name)
	if err != nil {
		return
	}
	if len(array) > 0 {
		return array[0], nil
	}
	return "", nil
}

// ToContainer attempts to acquire a container interface for the object.
func (o *object) ToContainer() (c *Container, err error) {
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

// ToComputer attempts to acquire a computer interface for the object.
func (o *object) ToComputer() (c *Computer, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		idispatch, err := o.iface.QueryInterface(api.IID_IADsComputer)
		if err != nil {
			return err
		}
		iface := (*api.IADsComputer)(unsafe.Pointer(idispatch))
		c = NewComputer(iface)
		return nil
	})
	return
}

// ToGroup attempts to acquire a group interface for the object.
func (o *object) ToGroup() (g *Group, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		idispatch, err := o.iface.QueryInterface(api.IID_IADsGroup)
		if err != nil {
			return err
		}
		iface := (*api.IADsGroup)(unsafe.Pointer(idispatch))
		g = NewGroup(iface)
		return nil
	})
	return
}
