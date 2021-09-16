package adsi

import (
	"encoding/hex"
	"fmt"
	"sync"
	"unsafe"

	"github.com/go-adsi/adsi/api"
	"github.com/go-adsi/adsi/comiid"
	ole "github.com/go-ole/go-ole"
	"github.com/google/uuid"
	"github.com/scjalliance/comshim"
	"github.com/scjalliance/comutil"
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
	o.iface.Release() // FIXME: What happens if release returns an error?
	o.iface = nil
}

// Name retrieves the name of the object.
func (o *object) Name() (name string, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return "", ErrClosed
	}
	name, err = o.iface.Name()
	return
}

// Class retrieves the class of the object.
func (o *object) Class() (class string, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return "", ErrClosed
	}
	class, err = o.iface.Class()
	return
}

// GUID retrieves the globally unique identifier of the object.
func (o *object) GUID() (guid uuid.UUID, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		err = ErrClosed
		return
	}

	var sguid string
	sguid, err = o.iface.GUID() // may return binary octet string in hexadecimal form
	if err != nil {
		return
	}

	switch len(sguid) {
	case 38:
		// 38 character representations include dashes and curly braces
		return uuid.Parse(sguid[1:37]) // Omit the braces
	case 32:
		// 32 character representations lack curly braces and dashes. When the
		// LDAP provider returns a GUID as a string in this form, it is the result
		// of taking the binary octet string of the GUID and converting it to
		// hexadecimal.
		_, err = hex.Decode(guid[:], []byte(sguid))
		if err != nil {
			return
		}

		// Assuming the original binary octet string was in the endianness
		// of the originating system, and that the originating system was
		// little-endian, we need to swap the bytes of the GUID we just parsed
		// to put them back in big-endian order.
		guid[0], guid[1], guid[2], guid[3] = guid[3], guid[2], guid[1], guid[0]
		guid[4], guid[5] = guid[5], guid[4]
		guid[6], guid[7] = guid[7], guid[6]
		return
	default:
		return uuid.Parse(sguid)
	}
}

// Path retrieves the fully qualified path of the object.
func (o *object) Path() (path string, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return "", ErrClosed
	}
	path, err = o.iface.AdsPath()
	return
}

// Parent retrieves the fully qualified path of the object's parent.
func (o *object) Parent() (path string, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return "", ErrClosed
	}
	path, err = o.iface.Parent()
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
	path, err = o.iface.Schema()
	return
}

// Pull causes the given list of attributes to be retrieved from the
// underlying LDAP data store. The retrieved attributes will be cached.
//
// Subsequent calls to the attr retrieval functions will return the cached
// values.
func (o *object) Pull(attrs ...string) (err error) {
	if len(attrs) == 0 {
		return nil
	}

	v, err := comutil.BuildVarArrayStr(attrs...)
	if err != nil {
		return
	}
	defer v.Clear()

	err = o.iface.GetInfoEx(v)
	return
}

// Attr attempts to retrieve the attribute with the given name and
// return its values as a slice of interfaces. Each value is an interface{}
// that holds a Go native type that is the best match for the underlying
// variant.
//
// If the attribute contains IUnknown or IDispatch members, it is the
// caller's responsibility to release them.
func (o *object) Attr(name string) (values []interface{}, err error) {
	variant, err := o.iface.GetEx(name)
	if err != nil {
		return nil, err
	}
	defer variant.Clear()

	array := variant.ToArray()
	if array == nil {
		return nil, ErrNonArrayAttribute
	}

	values, err = comutil.SafeArrayToVariantSlice(array)
	if err != nil {
		return nil, fmt.Errorf("unable to read \"%s\" attribute: %v", name, err)
	}
	return values, nil
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
		switch v := element.(type) {
		case string:
			values = append(values, v)
		case *ole.IUnknown:
			v.Release()
		case *ole.IDispatch:
			v.Release()
		default:
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

// AttrBytesSlice attempts to retrieve the attribute with the given name and
// return its values as a slice of byte slices.
//
// Any non-byte values contained in the attribute will be ommitted.
func (o *object) AttrBytesSlice(name string) (values [][]byte, err error) {
	elements, err := o.Attr(name)
	if err != nil {
		return
	}
	for _, element := range elements {
		switch v := element.(type) {
		case []byte:
			values = append(values, v)
		case *ole.IUnknown:
			v.Release()
		case *ole.IDispatch:
			v.Release()
		default:
			// TODO: Consider returning error
		}
	}
	return
}

// AttrBytes attempts to retrieve the attribute with the given name and
// return its value as a string. If the attribute holds more than one value,
// only the first value is returned.
//
// Any non-string values contained in the attribute will be ignored.
func (o *object) AttrBytes(name string) (attr []byte, err error) {
	array, err := o.AttrBytesSlice(name)
	if err != nil {
		return
	}
	if len(array) > 0 {
		return array[0], nil
	}
	return nil, nil
}

// AttrBoolSlice attempts to retrieve the attribute with the given name and
// return its values as a slice of bools.
//
// Any non-bool values contained in the attribute will be ommitted.
func (o *object) AttrBoolSlice(name string) (values []bool, err error) {
	elements, err := o.Attr(name)
	if err != nil {
		return
	}
	for _, element := range elements {
		switch v := element.(type) {
		case bool:
			values = append(values, v)
		case *ole.IUnknown:
			v.Release()
		case *ole.IDispatch:
			v.Release()
		default:
			// TODO: Consider returning error
		}
	}
	return
}

// AttrBool attempts to retrieve the attribute with the given name and
// return its value as a bool. If the attribute holds more than one value,
// only the first value is returned.
//
// Any non-bool values contained in the attribute will be ignored.
func (o *object) AttrBool(name string) (attr bool, err error) {
	array, err := o.AttrBoolSlice(name)
	if err != nil {
		return
	}
	if len(array) > 0 {
		return array[0], nil
	}
	return false, nil
}

// AttrIntSlice attempts to retrieve the attribute with the given name and
// return its values as a slice of integers.
//
// Any non-integer values contained in the attribute will be ommitted.
//
// Unsigned integer values will be coerced into signed types.
//
// 64-bit integer types will be coerced into integer types, which may
// overflow the value on 32-bit systems.
func (o *object) AttrIntSlice(name string) (values []int, err error) {
	elements, err := o.Attr(name)
	if err != nil {
		return
	}
	for i, element := range elements {
		switch v := element.(type) {
		case int:
			values = append(values, v)
		case uint:
			values = append(values, int(v))
		case int16:
			values = append(values, int(v))
		case uint16:
			values = append(values, int(v))
		case int32:
			values = append(values, int(v))
		case uint32:
			values = append(values, int(v))
		case int64:
			values = append(values, int(v))
		case uint64:
			values = append(values, int(v))
		case *ole.IUnknown:
			v.Release()
		case *ole.IDispatch:
			var value int64
			value, err = dispatchToInt64(v)
			v.Release()
			if err != nil {
				return nil, fmt.Errorf("attribute \"%s\" value %d: %v", name, i, err)
			}
			values = append(values, int(value))
		default:
			// TODO: Consider returning error
		}
	}
	return
}

// AttrInt attempts to retrieve the attribute with the given name and
// return its value as an integer. If the attribute holds more than one value,
// only the first value is returned.
//
// Any non-integer values contained in the attribute will be ignored.
func (o *object) AttrInt(name string) (attr int, err error) {
	array, err := o.AttrIntSlice(name)
	if err != nil {
		return
	}
	if len(array) > 0 {
		return array[0], nil
	}
	return 0, nil
}

// AttrInt64Slice attempts to retrieve the attribute with the given name and
// return its values as a slice of 64-bit integers.
//
// Any non-integer values contained in the attribute will be ommitted.
//
// Unsigned integer values will be coerced into signed types.
func (o *object) AttrInt64Slice(name string) (values []int64, err error) {
	elements, err := o.Attr(name)
	if err != nil {
		return nil, err
	}
	for i, element := range elements {
		switch v := element.(type) {
		case int:
			values = append(values, int64(v))
		case uint:
			values = append(values, int64(v))
		case int16:
			values = append(values, int64(v))
		case uint16:
			values = append(values, int64(v))
		case int32:
			values = append(values, int64(v))
		case uint32:
			values = append(values, int64(v))
		case int64:
			values = append(values, v)
		case uint64:
			values = append(values, int64(v))
		case *ole.IUnknown:
			v.Release()
		case *ole.IDispatch:
			var value int64
			value, err = dispatchToInt64(v)
			v.Release()
			if err != nil {
				return nil, fmt.Errorf("attribute \"%s\" value %d: %v", name, i, err)
			}
			values = append(values, value)
		default:
			// TODO: Consider returning error
		}
	}
	return
}

// AttrInt64 attempts to retrieve the attribute with the given name and
// return its value as a 64-bit integer. If the attribute holds more than one
// value, only the first value is returned.
//
// Any non-integer values contained in the attribute will be ignored.
func (o *object) AttrInt64(name string) (attr int64, err error) {
	array, err := o.AttrInt64Slice(name)
	if err != nil {
		return
	}
	if len(array) > 0 {
		return array[0], nil
	}
	return 0, nil
}

// AttrGUIDSlice attempts to retrieve the attribute with the given name and
// return its values as a slice of GUIDs.
//
// Any non-GUID values contained in the attribute will be ommitted.
//
// Values are returned as-is, without any byte ordering adjustment.
func (o *object) AttrGUIDSlice(name string) (values []uuid.UUID, err error) {
	elements, err := o.Attr(name)
	if err != nil {
		return
	}
	for _, element := range elements {
		switch v := element.(type) {
		case string:
			value, parseErr := uuid.Parse(v)
			if parseErr == nil {
				values = append(values, value)
			} else {
				// TODO: Consider returning error
			}
		case []byte:
			if len(v) == 16 {
				value, parseErr := uuid.FromBytes(v)
				if parseErr == nil {
					values = append(values, value)
				} else {
					// TODO: Consider returning error
				}
			} else {
				// TODO: Consider returning error
			}
		case *ole.IUnknown:
			v.Release()
		case *ole.IDispatch:
			v.Release()
		default:
			// TODO: Consider returning error
		}
	}
	return
}

// AttrGUID attempts to retrieve the attribute with the given name and
// return its value as a GUID in string format. If the attribute holds more
// than one value, only the first value is returned.
//
// Any non-string values contained in the attribute will be ignored.
func (o *object) AttrGUID(name string) (attr uuid.UUID, err error) {
	array, err := o.AttrGUIDSlice(name)
	if err != nil {
		return
	}
	if len(array) > 0 {
		attr = array[0]
	}
	return
}

// Put sets the values of an attribute in the ADSI attribute cache.
func (o *object) Put(name string, val interface{}) error {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return ErrClosed
	}
	return o.iface.Put(name, val)
}

// SetInfo saves the cached property values of the ADSI object to the underlying
// directory store.
func (o *object) SetInfo() error {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return ErrClosed
	}
	return o.iface.SetInfo()
}

// ToContainer attempts to acquire a container interface for the object.
func (o *object) ToContainer() (c *Container, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return nil, ErrClosed
	}
	idispatch, err := o.iface.QueryInterface(comutil.GUID(comiid.IADsContainer))
	if err != nil {
		return
	}
	iface := (*api.IADsContainer)(unsafe.Pointer(idispatch))
	c = NewContainer(iface)
	return
}

// ToComputer attempts to acquire a computer interface for the object.
func (o *object) ToComputer() (c *Computer, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return nil, ErrClosed
	}
	idispatch, err := o.iface.QueryInterface(comutil.GUID(comiid.IADsComputer))
	if err != nil {
		return
	}
	iface := (*api.IADsComputer)(unsafe.Pointer(idispatch))
	c = NewComputer(iface)
	return
}

// ToGroup attempts to acquire a group interface for the object.
func (o *object) ToGroup() (g *Group, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return nil, ErrClosed
	}
	idispatch, err := o.iface.QueryInterface(comutil.GUID(comiid.IADsGroup))
	if err != nil {
		return
	}
	iface := (*api.IADsGroup)(unsafe.Pointer(idispatch))
	g = NewGroup(iface)
	return
}
