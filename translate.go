/* NameTranslator translates distinguished names (DNs) among the various formats supported by Active Directory objects.

The translation is performed by the directory server.

Example, converting an NT4-style object to an RFC 1779 name:

tr, err := adsi.NameTranslator("")
if err != nil {
	return err
}
if err := tr.Init("mydomain", api.ADS_NAME_INITTYPE_DOMAIN); err != nil {
	return err
}
if err := tr.Set(`mydomain\john`, api.ADS_NAME_TYPE_NT4); err != nil {
	return err
}
x, err := tr.Get(api.ADS_NAME_TYPE_1779)
if err != nil {
	return err
}

See also https://docs.microsoft.com/en-us/windows/win32/api/iads/nn-iads-iadsnametranslate
*/
package adsi

import (
	"sync"

	"github.com/go-adsi/adsi/api"
	"github.com/scjalliance/comshim"
)

// NameTranslator provides active directory name translation services via the
// IADsNameTranslate interface.
type NameTranslator struct {
	m     sync.RWMutex
	iface *api.IADsNameTranslate
}

func (nt *NameTranslator) closed() bool {
	return (nt.iface == nil)
}

// NewNameTranslator creates a new NameTranslator bound to server.
// For the local host, leave server blank ("").
func NewNameTranslator(server string) (*NameTranslator, error) {
	comshim.Add(1)

	tr, err := api.NewIADsNameTranslate(server)
	if err != nil {
		return &NameTranslator{}, err
	}
	return &NameTranslator{iface: tr}, nil
}

// Close will release resources consumed by the translator. It should be
// called when the translator is no longer needed.
func (nt *NameTranslator) Close() {
	nt.m.Lock()
	defer nt.m.Unlock()
	if nt.closed() {
		return
	}
	nt.iface.Release()
	comshim.Done()
}

// Init initializes the translator to use either Domain, Server, or GC for translation.
// See ADS_NAME_INITTYPE_ENUM for possible values. adsPath is required for
// Domain and Server types, corresponding to the domain name or the server name respectively.
// Init must be called first.
func (nt *NameTranslator) Init(adsPath string, initType uint32) error {
	nt.m.Lock()
	defer nt.m.Unlock()
	if nt.closed() {
		return ErrClosed
	}
	return nt.iface.Init(adsPath, initType)
}

// Get attempts to get the result of a translation in the specified format. See
// ADS_NAME_TYPE_ENUM for the available options. Get must be called last.
func (nt *NameTranslator) Get(formatType uint32) (string, error) {
	nt.m.Lock()
	defer nt.m.Unlock()
	if nt.closed() {
		return "", ErrClosed
	}
	return nt.iface.Get(formatType)
}

// Set sets the input object to be translated. The adsPath must resolve to a valid domain object,
// in the format matching setType (an ADS_NAME_TYPE_ENUM). Set must be called after Init and
// before Get.
func (nt *NameTranslator) Set(adsPath string, setType uint32) error {
	nt.m.Lock()
	defer nt.m.Unlock()
	if nt.closed() {
		return ErrClosed
	}
	return nt.iface.Set(adsPath, setType)
}
