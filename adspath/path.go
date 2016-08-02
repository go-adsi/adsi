package adspath

import (
	"bytes"
	"errors"
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
func Parse(rawpath string) (path *Path, err error) {
	var rest string

	if rawpath == "" {
		return nil, errors.New("empty path")
	}

	path = new(Path)

	if path.Scheme, rest, err = parseScheme(rawpath); err != nil {
		return nil, err
	}

	// Enforce correct capitalization of the scheme, because many ADSI schemes are
	// case-sensitive.
	//
	// TODO: Decide whether it's really appropriate to do this auto-correction.
	switch strings.ToLower(path.Scheme) {
	case "ldap":
		path.Scheme = LDAP
	case "winnt":
		path.Scheme = WinNT
	case "iis":
		path.Scheme = IIS
	case "gc":
		path.Scheme = GC
	}

	if rest == "" {
		// We're binding to the root of the namespace
		return
	}

	if !strings.HasPrefix(rest, "//") {
		return nil, errors.New("Invalid ADS path")
	}

	rest = rest[2:]

	authority, rest := split(rest, "/")
	if rest == "" && path.Scheme == LDAP && strings.ContainsRune(authority, '=') {
		// This is serverless LDAP binding
		rest = authority
		authority = ""
	}

	path.Host, path.Path = authority, rest

	return
}

func parseScheme(rawpath string) (scheme, path string, err error) {
	for i := 0; i < len(rawpath); i++ {
		c := rawpath[i]
		switch {
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z':
		// do nothing
		case '0' <= c && c <= '9' || c == '+' || c == '-' || c == '.':
			if i == 0 {
				return "", rawpath, nil
			}
		case c == ':':
			if i == 0 {
				return "", "", errors.New("missing protocol scheme")
			}
			return rawpath[:i], rawpath[i+1:], nil
		default:
			// we have encountered an invalid character,
			// so there is no valid scheme
			return "", rawpath, nil
		}
	}
	return "", rawpath, nil
}

func split(input string, cut string) (string, string) {
	i := strings.Index(input, cut)
	if i < 0 {
		return input, ""
	}
	return input[:i], input[i+len(cut):]
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
	var buf bytes.Buffer
	if p.Scheme != "" {
		buf.WriteString(p.Scheme)
		buf.WriteByte(':')
	}
	if p.Scheme != "" && (p.Host != "" || p.Path != "") {
		buf.WriteString("//")
	}
	if p.Host != "" {
		buf.WriteString(p.Host)
		if p.Path != "" {
			buf.WriteByte('/')
		}
	}
	buf.WriteString(p.Path)
	return buf.String()
}
