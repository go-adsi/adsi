package adsi

import "gopkg.in/adsi.v0/api"

// Open opens an ADSI object with the given path. It creates an ephemeral
// local client and uses it to open the requested object. The connection is made
// using the security context of the application. The object is opened readonly.
//
// Open returns the ADSI object as an Object type, which provides
// an idiomatic go wrapper around the underlying component object model
// IADs interface.
//
// If the requested ADSI object does not implement the IADs interface an error
// is returned.
//
// The returned object consumes resources until it is closed. It is the
// caller's responsibilty to call Close on the returned object when it is no
// longer needed.
func Open(path string) (obj *Object, err error) {
	return OpenObject(path, "", "", api.ADS_READONLY_SERVER|api.ADS_SECURE_AUTHENTICATION|api.ADS_USE_SEALING)
}

// OpenObject opens an ADSI object with the given path. Most users will use Open
// instead. It creates an ephemeral local client and uses it to open the
// requested object.
//
// When provided, the username and password are used to establish a security
// context for the connection. When credentials are not provided the existing
// security context of the application is used instead.
//
// OpenObject returns the ADSI object as an Object type, which provides
// an idiomatic go wrapper around the underlying component object model
// IADs interface.
//
// If the requested ADSI object does not implement the IADs interface an error
// is returned.
//
// The returned object consumes resources until it is closed. It is the
// caller's responsibilty to call Close on the returned object when it is no
// longer needed.
func OpenObject(path, user, password string, flags uint32) (obj *Object, err error) {
	var c *Client

	c, err = NewClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()

	obj, err = c.Open(path, user, password, flags)
	return
}
