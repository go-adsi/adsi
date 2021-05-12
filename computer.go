package adsi

import (
	"github.com/scjalliance/comshim"

	"github.com/go-adsi/adsi/api"
)

// Computer provides access to Active Directory computers.
type Computer struct {
	object
	iface *api.IADsComputer
}

// NewComputer returns a computer that manages the given COM interface.
func NewComputer(iface *api.IADsComputer) *Computer {
	comshim.Add(1)
	return &Computer{iface: iface, object: object{iface: &iface.IADs}}
}

func (c *Computer) closed() bool {
	return (c.iface == nil)
}

// Close will release resources consumed by the computer. It should be
// called when the computer is no longer needed.
func (c *Computer) Close() {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return
	}
	defer comshim.Done()
	c.iface.Release()
	c.object.iface = nil
	c.iface = nil
}

// ID retrieves the ID of the computer.
func (c *Computer) ID() (id string, err error) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return "", ErrClosed
	}
	id, err = c.iface.ComputerID()
	return
}

// Site retrieves the site of the computer.
func (c *Computer) Site() (site string, err error) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return "", ErrClosed
	}
	site, err = c.iface.Site()
	return
}

// OperatingSystem retrieves the operating system of the computer.
func (c *Computer) OperatingSystem() (kind string, err error) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return "", ErrClosed
	}
	kind, err = c.iface.OperatingSystem()
	return
}
