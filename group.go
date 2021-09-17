package adsi

import (
	"github.com/scjalliance/comshim"

	"github.com/go-adsi/adsi/api"
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

// Add adds an ADSI object to an existing group.
func (g *Group) Add(item string) (err error) {
	g.m.Lock()
	defer g.m.Unlock()
	if g.closed() {
		return ErrClosed
	}
	return g.iface.Add(item)
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
	g.iface.Release()
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
	desc, err = g.iface.Description()
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
	imembers, err := g.iface.Members()
	if err != nil {
		return
	}
	m = NewMembers(imembers)
	return
}

// Remove removes the specified user object from this group. The operation
// does not remove the group object itself even when there is no member remaining in the group.
func (g *Group) Remove(item string) error {
	g.m.Lock()
	defer g.m.Unlock()
	if g.closed() {
		return ErrClosed
	}
	return g.iface.Remove(item)
}
