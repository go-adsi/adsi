package api

import "github.com/go-ole/go-ole"

type CoAuthInfo struct {
	AuthenticationService uint32
	AuthorizationService  uint32
	ServerPrincipalName   *int16
	AuthenticationLevel   uint32
	ImpersonationLevel    uint32
}

type CoServerInfo struct {
	_        uint32 // reserved
	Name     *int16
	AuthInfo *CoAuthInfo
	_        uint32 // reserved
}

type MultiQI struct {
	IID       *ole.GUID
	Interface *ole.IUnknown
	HR        uintptr
}
