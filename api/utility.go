package api

import (
	"strings"

	"github.com/go-ole/go-ole"
)

// convertHresultToError converts syscall to error, if call is unsuccessful.
func convertHresultToError(hr uintptr) (err error) {
	if hr != 0 {
		err = ole.NewError(hr)
		if strings.Contains(err.Error(), "FormatMessage failed") {
			switch hr {
			case E_INVALID_NAMESPACE:
				err = ErrInvalidNamespace
			case E_ACCESS_DENIED:
				err = ErrAccessDenied
			case S_ADS_ERRORSOCCURRED:
				err = ErrQueryFailed
			case S_ADS_NOMORE_ROWS:
				err = ErrNoMoreRows
			case S_ADS_NOMORE_COLUMNS:
				err = ErrNoMoreColumns
			case E_ADS_BAD_PATHNAME:
				err = ErrBadPathname
			case E_ADS_INVALID_DOMAIN_OBJECT:
				err = ErrInvalidDomainObject
			case E_ADS_INVALID_USER_OBJECT:
				err = ErrInvalidUserObject
			case E_ADS_INVALID_COMPUTER_OBJECT:
				err = ErrInvalidComputerObject
			case E_ADS_UNKNOWN_OBJECT:
				err = ErrUnknownObject
			case E_ADS_PROPERTY_NOT_SET:
				err = ErrPropertyNotSet
			case E_ADS_PROPERTY_NOT_SUPPORTED:
				err = ErrPropertyNotSupported
			case E_ADS_PROPERTY_INVALID:
				err = ErrPropertyInvalid
			case E_ADS_BAD_PARAMETER:
				err = ErrBadParameter
			case E_ADS_OBJECT_UNBOUND:
				err = ErrObjectUnbound
			case E_ADS_PROPERTY_NOT_MODIFIED:
				err = ErrPropertyNotModified
			case E_ADS_PROPERTY_MODIFIED:
				err = ErrPropertyModified
			case E_ADS_CANT_CONVERT_DATATYPE:
				err = ErrCantConvertDatatype
			case E_ADS_PROPERTY_NOT_FOUND:
				err = ErrPropertyNotFound
			case E_ADS_OBJECT_EXISTS:
				err = ErrObjectExists
			case E_ADS_SCHEMA_VIOLATION:
				err = ErrSchemaViolation
			case E_ADS_COLUMN_NOT_SET:
				err = ErrColumnNotSet
			case E_ADS_INVALID_FILTER:
				err = ErrInvalidFilter
			}
		}
	}
	return
}
