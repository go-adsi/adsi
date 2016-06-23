package adspath

import "errors"

var (
	// ErrFragmentNotPermitted is returned when an attempt is made to parse a
	// path that has a URL fragment. URL fragments are not permitted in ADS paths.
	ErrFragmentNotPermitted = errors.New("The URL contains a fragment, which is not permitted within Active Directory paths.")
	// ErrUserInfoNotPermitted is returned when an attempt is made to parse a
	// path that has a user information section. User info cannot be specified in
	// ADS paths.
	ErrUserInfoNotPermitted = errors.New("The URL contains a username or password, which is not permitted within Active Directory paths.")
)

const (
	// LDAP is the scheme of the LDAP namespace provider.
	LDAP = "LDAP"
	// WinNT is the scheme of the WinNT namespace provider.
	WinNT = "WinNT"
	// IIS is the scheme of the IIS namespace provider.
	IIS = "IIS"
	// GC is the scheme of the GC namespace provider.
	GC = "GC"
)
