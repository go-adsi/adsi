package adsi

import (
	"sync"

	"gopkg.in/adsi.v0/api"
)

// User provides access to Active Directory users.
type User struct {
	m     sync.RWMutex
	iface *api.IADsUser
}
