package api

import "errors"

const (
	E_INVALID_NAMESPACE = 0x8004100E
	E_ACCESS_DENIED     = 0x80041003

	// See https://msdn.microsoft.com/en-us/library/aa705940

	S_ADS_ERRORSOCCURRED          = 0x00005011
	S_ADS_NOMORE_ROWS             = 0x00005012
	S_ADS_NOMORE_COLUMNS          = 0x00005013
	E_ADS_BAD_PATHNAME            = 0x80005000
	E_ADS_INVALID_DOMAIN_OBJECT   = 0x80005001
	E_ADS_INVALID_USER_OBJECT     = 0x80005002
	E_ADS_INVALID_COMPUTER_OBJECT = 0x80005003
	E_ADS_UNKNOWN_OBJECT          = 0x80005004
	E_ADS_PROPERTY_NOT_SET        = 0x80005005
	E_ADS_PROPERTY_NOT_SUPPORTED  = 0x80005006
	E_ADS_PROPERTY_INVALID        = 0x80005007
	E_ADS_BAD_PARAMETER           = 0x80005008
	E_ADS_OBJECT_UNBOUND          = 0x80005009
	E_ADS_PROPERTY_NOT_MODIFIED   = 0x8000500A
	E_ADS_PROPERTY_MODIFIED       = 0x8000500B
	E_ADS_CANT_CONVERT_DATATYPE   = 0x8000500C
	E_ADS_PROPERTY_NOT_FOUND      = 0x8000500D
	E_ADS_OBJECT_EXISTS           = 0x8000500E
	E_ADS_SCHEMA_VIOLATION        = 0x8000500F
	E_ADS_COLUMN_NOT_SET          = 0x80005010
	E_ADS_INVALID_FILTER          = 0x80005014
)

const (
	// See https://msdn.microsoft.com/library/aa772247

	ADS_SECURE_AUTHENTICATION = 0x1
	ADS_USE_ENCRYPTION        = 0x2
	ADS_USE_SSL               = 0x2
	ADS_READONLY_SERVER       = 0x4
	ADS_PROMPT_CREDENTIALS    = 0x8
	ADS_NO_AUTHENTICATION     = 0x10
	ADS_FAST_BIND             = 0x20
	ADS_USE_SIGNING           = 0x40
	ADS_USE_SEALING           = 0x80
	ADS_USE_DELEGATION        = 0x100
	ADS_SERVER_BIND           = 0x200
	ADS_NO_REFERRAL_CHASING   = 0x400
	ADS_AUTH_RESERVED         = 0x80000000
)

var (
	ErrInvalidNamespace = errors.New("The provided name or namespace is invalid.")
	ErrAccessDenied     = errors.New("Access denied.")

	// See https://msdn.microsoft.com/en-us/library/aa705940

	ErrQueryFailed           = errors.New("During a query, one or more errors occurred.")
	ErrNoMoreRows            = errors.New("The search operation has reached the last row.")
	ErrNoMoreColumns         = errors.New("The search operation has reached the last column for the current row.")
	ErrBadPathname           = errors.New("An invalid ADSI pathname was passed.")
	ErrInvalidDomainObject   = errors.New("An unknown ADSI domain object was requested.")
	ErrInvalidUserObject     = errors.New("An unknown ADSI user object was requested.")
	ErrInvalidComputerObject = errors.New("An unknown ADSI computer object was requested.")
	ErrUnknownObject         = errors.New("An unknown ADSI object was requested.")
	ErrPropertyNotSet        = errors.New("The specified ADSI property was not set.")
	ErrPropertyNotSupported  = errors.New("The specified ADSI property is not supported.")
	ErrPropertyInvalid       = errors.New("The specified ADSI property is invalid.")
	ErrBadParameter          = errors.New("One or more input parameters are invalid.")
	ErrObjectUnbound         = errors.New("The specified ADSI object is not bound to a remote resource.")
	ErrPropertyNotModified   = errors.New("The specified ADSI object has not been modified.")
	ErrPropertyModified      = errors.New("The specified ADSI object has been modified.")
	ErrCantConvertDatatype   = errors.New("The data type cannot be converted to/from a native DS data type.")
	ErrPropertyNotFound      = errors.New("The property cannot be found in the cache.")
	ErrObjectExists          = errors.New("The ADSI object exists.")
	ErrSchemaViolation       = errors.New("The attempted action violates the directory service schema rules.")
	ErrColumnNotSet          = errors.New("The specified column in the ADSI was not set.")
	ErrInvalidFilter         = errors.New("The specified search filter is invalid.")
)
