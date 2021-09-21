package adsi

import (
	"github.com/go-adsi/adsi/api"
	"github.com/scjalliance/comshim"
)

// User provides access to Active Directory users.
type User struct {
	object
	iface *api.IADsUser
}

// NewUser returns a user that manages the given COM interface.
func NewUser(iface *api.IADsUser) *User {
	comshim.Add(1)
	return &User{iface: iface, object: object{iface: &iface.IADs}}
}

func (u *User) closed() bool {
	return (u.iface == nil)
}

// Close will release resources consumed by the user. It should be
// called when the user is no longer needed.
func (u *User) Close() {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return
	}
	defer comshim.Done()
	u.iface.Release()
	u.object.iface = nil
	u.iface = nil
}

// AccountDisabled retrieves the disablement status of a user account.
func (u *User) AccountDisabled() (disabled bool, err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return false, ErrClosed
	}
	return u.iface.AccountDisabled()
}

// SetAccountDisabled sets an account as disabled.
func (u *User) SetAccountDisabled(val bool) (err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return ErrClosed
	}
	return u.iface.SetAccountDisabled(val)
}

// FullName returns the user's FullName property.
func (u *User) FullName() (string, error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return "", ErrClosed
	}
	return u.iface.FullName()
}
