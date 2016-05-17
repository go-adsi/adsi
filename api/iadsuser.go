package api

import "unsafe"

// IADsUserVtbl represents the component object model virtual
// function table for the IADsUser interface.
type IADsUserVtbl struct {
	IADsVtbl
	BadLoginAddress           uintptr
	BadLoginCount             uintptr
	LastLogin                 uintptr
	LastLogoff                uintptr
	LastFailedLogin           uintptr
	PasswordLastChanged       uintptr
	Description               uintptr
	SetDescription            uintptr
	Division                  uintptr
	SetDivision               uintptr
	Department                uintptr
	SetDepartment             uintptr
	EmployeeID                uintptr
	SetEmployeeID             uintptr
	FullName                  uintptr
	SetFullName               uintptr
	FirstName                 uintptr
	SetFirstName              uintptr
	LastName                  uintptr
	SetLastName               uintptr
	OtherName                 uintptr
	SetOtherName              uintptr
	NamePrefix                uintptr
	SetNamePrefix             uintptr
	NameSuffix                uintptr
	SetNameSuffix             uintptr
	Title                     uintptr
	SetTitle                  uintptr
	Manager                   uintptr
	SetManager                uintptr
	TelephoneHome             uintptr
	SetTelephoneHome          uintptr
	TelephoneMobile           uintptr
	SetTelephoneMobile        uintptr
	TelephoneNumber           uintptr
	SetTelephoneNumber        uintptr
	TelephonePager            uintptr
	SetTelephonePager         uintptr
	FaxNumber                 uintptr
	SetFaxNumber              uintptr
	OfficeLocations           uintptr
	SetOfficeLocations        uintptr
	PostalAddresses           uintptr
	SetPostalAddresses        uintptr
	PostalCodes               uintptr
	SetPostalCodes            uintptr
	SeeAlso                   uintptr
	SetSeeAlso                uintptr
	AccountDisabled           uintptr
	SetAccountDisabled        uintptr
	AccountExpirationDate     uintptr
	SetAccountExpirationDate  uintptr
	GraceLoginsAllowed        uintptr
	SetGraceLoginsAllowed     uintptr
	GraceLoginsRemaining      uintptr
	SetGraceLoginsRemaining   uintptr
	IsAccountLocked           uintptr
	SetIsAccountLocked        uintptr
	LoginHours                uintptr
	SetLoginHours             uintptr
	LoginWorkstations         uintptr
	SetLoginWorkstations      uintptr
	MaxLogins                 uintptr
	SetMaxLogins              uintptr
	MaxStorage                uintptr
	SetMaxStorage             uintptr
	PasswordExpirationDate    uintptr
	SetPasswordExpirationDate uintptr
	PasswordMinimumLength     uintptr
	SetPasswordMinimumLength  uintptr
	PasswordRequired          uintptr
	SetPasswordRequired       uintptr
	RequireUniquePassword     uintptr
	SetRequireUniquePassword  uintptr
	EmailAddress              uintptr
	SetEmailAddress           uintptr
	HomeDirectory             uintptr
	SetHomeDirectory          uintptr
	Languages                 uintptr
	SetLanguages              uintptr
	Profile                   uintptr
	SetProfile                uintptr
	LoginScript               uintptr
	SetLoginScript            uintptr
	Picture                   uintptr
	SetPicture                uintptr
	HomePage                  uintptr
	SetHomePage               uintptr
	Groups                    uintptr
	SetPassword               uintptr
	ChangePassword            uintptr
}

// IADsUser represents the component object model interface for
// active directory users.
type IADsUser struct {
	IADs
}

// VTable returns the component object model virtual function table for the
// user.
func (v *IADsUser) VTable() *IADsUserVtbl {
	return (*IADsUserVtbl)(unsafe.Pointer(v.RawVTable))
}
