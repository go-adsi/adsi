package adspath

import (
	"net/url"
	"strings"
)

// Path represents the location of an Active Directory Services object.
//
// Paths are commonly represented in the form of a URL.
type Path struct {
	// Scheme indicates the type of namespace provider used to access the path.
	// It is sometimes also referred to as the protocol or as the ProgID.
	//
	// It is important to note that many namespace providers require a
	// case-sensitive match of the scheme.
	Scheme string
	// Host may be the computer, domain, or nothing for serverless binding. It may
	// be in the form "host" or "host:port".
	Host string
	// Path is the address of the resource within the namespace. The rules for its
	// interpretation vary between schemes.
	Path string
}

// Parse will attempt to interpret the raw path provided as a string in URL
// form.
//
// Internally it relies on the URL parsing capabilities of the `url` library.
func Parse(rawpath string) (path *Path, err error) {
	var u *url.URL
	u, err = url.Parse(rawpath)
	if err != nil {
		return
	}

	if u.User != nil {
		return nil, &url.Error{Op: "parse", URL: rawpath, Err: ErrUserInfoNotPermitted}
	}

	if u.Fragment != "" {
		return nil, &url.Error{Op: "parse", URL: rawpath, Err: ErrFragmentNotPermitted}
	}

	// Enforce correct capitalization of the scheme, because many ADSI schemes are
	// case-sensitive.
	//
	// TODO: Decide whether its really appropriate to do this auto-correction.
	switch strings.ToLower(u.Scheme) {
	case "ldap":
		u.Scheme = LDAP
	case "winnt":
		u.Scheme = WinNT
	case "iis":
		u.Scheme = IIS
	case "gc":
		u.Scheme = GC
	}

	return &Path{
		Scheme: u.Scheme,
		Host:   u.Host,
		Path:   u.Path,
	}, nil
}

// URL converts the path to a url.URL.
func (p *Path) URL() (u *url.URL) {
	return &url.URL{
		Scheme: p.Scheme,
		Host:   p.Host,
		Path:   p.Path,
	}
}

// String returns the path as a string in URL form.
func (p *Path) String() string {
	return p.URL().String()
}
