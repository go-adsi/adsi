package adsi

import (
	"errors"
	"unsafe"

	ole "github.com/go-ole/go-ole"
	"github.com/scjalliance/comutil"
	"github.com/go-adsi/adsi/v2/api"
	"github.com/go-adsi/adsi/v2/comiid"
)

func reverseUint16(v uint16) uint16 {
	return v>>8&0x00ff | v<<8&0xff00
}

func reverseUint32(v uint32) uint32 {
	return v>>24&0x000000ff | v>>8&0x0000ff00 | v<<8&0x00ff0000 | v<<24&0xff000000
}

func dispatchToInt64(v *ole.IDispatch) (value int64, err error) {
	if v == nil {
		return 0, errors.New("nil IDispatch interface")
	}

	iface, err := v.QueryInterface(comutil.GUID(comiid.IADsLargeInteger))
	if err == nil {
		defer iface.Release()
		largeInt := (*api.IADsLargeInteger)(unsafe.Pointer(iface))
		return largeInt.Value()
	}

	iface, err = v.QueryInterface(comutil.GUID(comiid.IADsPropertyValue))
	if err == nil {
		defer iface.Release()
		err = errors.New("ADSI property values are not yet supported")
		return
	}

	return 0, errors.New("unsupported COM interface for integer conversion")
}
