package adsi

// Open opens an ADSI object with the given path. It creates an ephemeral
// local client and uses it to open the requested object. The connection is made
// using the security context of the application and the default client flags
// specifying that it be encrypted and read-only.
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
	c, err := NewClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()

	return c.Open(path)
}

// OpenSC opens an ADSI object with the given path. Most users will use Open
// instead. It creates an ephemeral local client and uses it to open the
// requested object.
//
// When provided, the username and password are used to establish a security
// context for the connection. When they are not provided the existing
// security context of the application is used instead. The provided flags will
// be used to make the connection.
//
// OpenSC returns the ADSI object as an Object type, which provides
// an idiomatic go wrapper around the underlying component object model
// IADs interface.
//
// If the requested ADSI object does not implement the IADs interface an error
// is returned.
//
// The returned object consumes resources until it is closed. It is the
// caller's responsibilty to call Close on the returned object when it is no
// longer needed.
func OpenSC(path, user, password string, flags uint32) (obj *Object, err error) {
	c, err := NewClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.OpenSC(path, user, password, flags)
}

// OpenContainer opens an ADSI container with the given path. It creates an
// ephemeral local client and uses it to open the requested container. The
// connection is made using the security context of the application and the
// default client flags specifying that it be encrypted and read-only.
//
// OpenContainer returns the ADSI container as a Container type, which provides
// an idiomatic go wrapper around the underlying component object model
// IADsContainer interface.
//
// If the returned directory object does not implement the IADsContainer
// interface an error is returned.
//
// The returned container consumes resources until it is closed. It is the
// caller's responsibilty to call Close on the returned container when it is no
// longer needed.
func OpenContainer(path string) (container *Container, err error) {
	c, err := NewClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.OpenContainer(path)
}

// OpenContainerSC opens an ADSI container with the given path. Most users will
// use OpenContainer instead. It creates an ephemeral local client and uses it
// to open the requested container.
//
// When provided, the username and password are used to establish a security
// context for the connection. When they are not provided the existing
// security context of the application is used instead. The provided flags will
// be used to make the connection.
//
// OpenContainerSC returns the ADSI container as a Container type, which
// provides an idiomatic go wrapper around the underlying component object model
// IADsContainer interface.
//
// If the returned directory object does not implement the IADsContainer
// interface an error is returned.
//
// The returned container consumes resources until it is closed. It is the
// caller's responsibilty to call Close on the returned container when it is no
// longer needed.
func OpenContainerSC(path, user, password string, flags uint32) (container *Container, err error) {
	c, err := NewClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.OpenContainerSC(path, user, password, flags)
}

// OpenComputer opens an ADSI computer with the given path. The existing
// security context of the application and any flags specified via SetFlags will
// be used when making the connection. The default flags specify an encrypted
// read-only connection.
//
// OpenComputer returns the ADSI computer as a Computer type, which provides
// an idiomatic go wrapper around the underlying component object model
// IADsComputer interface.
//
// If the returned directory object does not implement the IADsComputer
// interface an error is returned.
//
// The returned computer consumes resources until it is closed. It is the
// caller's responsibilty to call Close on the returned computer when it is no
// longer needed.
func OpenComputer(path string) (computer *Computer, err error) {
	c, err := NewClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.OpenComputer(path)
}

// OpenComputerSC opens an ADSI computer with the given path. Most users will
// use OpenComputer instead. It creates an ephemeral local client and uses it
// to open the requested computer.
//
// When provided, the username and password are used to establish a security
// context for the connection. When they are not provided the existing
// security context of the application is used instead. The provided flags will
// be used to make the connection.
//
// OpenComputerSC returns the ADSI computer as a Computer type, which
// provides an idiomatic go wrapper around the underlying component object model
// IADsComputer interface.
//
// If the returned directory object does not implement the IADsComputer
// interface an error is returned.
//
// The returned computer consumes resources until it is closed. It is the
// caller's responsibilty to call Close on the returned computer when it is no
// longer needed.
func OpenComputerSC(path, user, password string, flags uint32) (computer *Computer, err error) {
	c, err := NewClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.OpenComputerSC(path, user, password, flags)
}
