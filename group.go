package adsi

import (
	"sync"

	"gopkg.in/adsi.v0/api"
)

// Group provides access to Active Directory groups.
type Group struct {
	m     sync.RWMutex
	iface *api.IADsGroup
}
