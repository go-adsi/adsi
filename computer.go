package adsi

import (
	"sync"

	"github.com/scjalliance/comshim"

	"gopkg.in/adsi.v0/api"
)

// Computer provides access to Active Directory computers.
type Computer struct {
	m     sync.RWMutex
	iface *api.IADsComputer
}

// NewComputer returns a computer that manages the given COM interface.
func NewComputer(iface *api.IADsComputer) *Computer {
	comshim.Add(1)
	return &Computer{iface: iface}
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
	run(func() error {
		c.iface.Release()
		return nil
	})
	c.iface = nil
}

// ID retrieves the ID of the computer.
func (c *Computer) ID() (id string, err error) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return "", ErrClosed
	}
	err = run(func() error {
		id, err = c.iface.ComputerID()
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// Site retrieves the site of the computer.
func (c *Computer) Site() (site string, err error) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return "", ErrClosed
	}
	err = run(func() error {
		site, err = c.iface.Site()
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// OperatingSystem retrieves the operating system of the computer.
func (c *Computer) OperatingSystem() (kind string, err error) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return "", ErrClosed
	}
	err = run(func() error {
		kind, err = c.iface.OperatingSystem()
		if err != nil {
			return err
		}
		return nil
	})
	return
}
