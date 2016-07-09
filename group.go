package adsi

import (
	"github.com/scjalliance/comshim"

	"gopkg.in/adsi.v0/api"
)

// Group provides access to Active Directory groups.
type Group struct {
	object
	iface *api.IADsGroup
}

// NewGroup returns a group that manages the given COM interface.
func NewGroup(iface *api.IADsGroup) *Group {
	comshim.Add(1)
	return &Group{iface: iface, object: object{iface: &iface.IADs}}
}

func (g *Group) closed() bool {
	return (g.iface == nil)
}

// Close will release resources consumed by the group. It should be
// called when the group is no longer needed.
func (g *Group) Close() {
	g.m.Lock()
	defer g.m.Unlock()
	if g.closed() {
		return
	}
	defer comshim.Done()
	run(func() error {
		g.iface.Release()
		return nil
	})
	g.object.iface = nil
	g.iface = nil
}

// Description retrieves the description of the group.
func (g *Group) Description() (desc string, err error) {
	g.m.Lock()
	defer g.m.Unlock()
	if g.closed() {
		return "", ErrClosed
	}
	err = run(func() error {
		desc, err = g.iface.Description()
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// Members returns a membership that provides access to the members of the
// group.
func (g *Group) Members() (m *Members, err error) {
	g.m.Lock()
	defer g.m.Unlock()
	if g.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		imembers, err := g.iface.Members()
		if err != nil {
			return err
		}
		m = NewMembers(imembers)
		return nil
	})
	return
}
