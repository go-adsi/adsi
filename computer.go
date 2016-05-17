package adsi

import (
	"sync"

	"gopkg.in/adsi.v0/api"
)

// Computer provides access to Active Directory computers.
type Computer struct {
	m     sync.RWMutex
	iface *api.IADsComputer
}
