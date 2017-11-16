// Package comiid provides component object model interface identifiers for
// active directory domain services.
package comiid

import (
	"github.com/google/uuid"
)

// See: https://github.com/blueardour/pxe-server/blob/master/0.winxp-drivers-via-binlsrv/inf/activeds.inf
// See: https://msdn.microsoft.com/library/aa746443

var (
	// IADs is the component object model identifier of the IADs interface.
	//
	// IID_IADs
	// {FD8256D0-FD15-11CE-ABC4-02608C9E7553}
	IADs = uuid.UUID{0xFD, 0x82, 0x56, 0xD0, 0xFD, 0x15, 0x11, 0xCE, 0xAB, 0xC4, 0x02, 0x60, 0x8C, 0x9E, 0x75, 0x53}

	// IADsNamespaces is the component object model identifier of the
	// IADsNamespaces interface.
	//
	// IID_IADsNamespaces
	// {28B96BA0-B330-11CF-A9AD-00AA006BC149}
	IADsNamespaces = uuid.UUID{0x28, 0xB9, 0x6B, 0xA0, 0xB3, 0x30, 0x11, 0xCF, 0xA9, 0xAD, 0x00, 0xAA, 0x00, 0x6B, 0xC1, 0x49}

	// IADsOpenDSObject is the component object model identifier of the
	// IADsOpenDSObject interface.
	//
	// IID_IADsOpenDSObject
	// {DDF2891E-0F9C-11D0-8AD4-00C04FD8D503}
	IADsOpenDSObject = uuid.UUID{0xDD, 0xF2, 0x89, 0x1E, 0x0F, 0x9C, 0x11, 0xD0, 0x8A, 0xD4, 0x00, 0xC0, 0x4F, 0xD8, 0xD5, 0x03}

	// IADsContainer is the component object model identifier of the
	// IADsContainer interface.
	//
	// IID_IADsContainer
	// {001677D0-FD16-11CE-ABC4-02608C9E7553}
	IADsContainer = uuid.UUID{0x00, 0x16, 0x77, 0xD0, 0xFD, 0x16, 0x11, 0xCE, 0xAB, 0xC4, 0x02, 0x60, 0x8C, 0x9E, 0x75, 0x53}

	// IADsComputer is the component object model identifier of the
	// IADsComputer interface.
	//
	// IID_IADsComputer
	// {EFE3CC70-1D9F-11CF-B1F3-02608C9E7553}
	IADsComputer = uuid.UUID{0xEF, 0xE3, 0xCC, 0x70, 0x1D, 0x9F, 0x11, 0xCF, 0xB1, 0xF3, 0x02, 0x60, 0x8C, 0x9E, 0x75, 0x53}

	// IADsGroup is the component object model identifier of the
	// IADsGroup interface.
	//
	// IID_IADsGroup
	// {27636B00-410F-11CF-B1FF-02608C9E7553}
	IADsGroup = uuid.UUID{0x27, 0x63, 0x6B, 0x00, 0x41, 0x0F, 0x11, 0xCF, 0xB1, 0xFF, 0x02, 0x60, 0x8C, 0x9E, 0x75, 0x53}

	// IADsMembers is the component object model identifier of the
	// IADsMembers interface.
	//
	// IID_IADsMembers
	// {451A0030-72EC-11CF-B03B-00AA006E0975}
	IADsMembers = uuid.UUID{0x45, 0x1A, 0x00, 0x30, 0x72, 0xEC, 0x11, 0xCF, 0xB0, 0x3B, 0x00, 0xAA, 0x00, 0x6E, 0x09, 0x75}
)
