package enums

type ErrorCode string

func (ec ErrorCode) Error() string {
	return string(ec)
}

var Error = struct {
	BookNotFound        ErrorCode
	DataNotFound        ErrorCode
	DisabledAccount     ErrorCode
	FileNotFound        ErrorCode
	InvalidFileFormat   ErrorCode
	InvalidInputData    ErrorCode
	InvalidPassword     ErrorCode
	InvalidTimeValue    ErrorCode
	MissingJwtAuth      ErrorCode
	NonActivatedAccount ErrorCode
	PayloadTooLarge     ErrorCode
	PermissionDenied    ErrorCode
	UserNotFound        ErrorCode
}{
	BookNotFound:        "BOOK_NOT_FOUND",
	DataNotFound:        "DATA_NOT_FOUND",
	DisabledAccount:     "DISABLED_ACCOUNT",
	FileNotFound:        "FILE_NOT_FOUND",
	InvalidFileFormat:   "INVALID_FILE_FORMAT",
	InvalidInputData:    "INVALID_INPUT_DATA",
	InvalidPassword:     "INVALID_PASSWORD",
	InvalidTimeValue:    "INVALID_TIME_VALUE",
	MissingJwtAuth:      "MISSING_JWT_AUTH",
	NonActivatedAccount: "NON_ACTIVATED_ACCOUNT",
	PayloadTooLarge:     "PAYLOAD_TOO_LARGE",
	PermissionDenied:    "PERMISSION_DENIED",
	UserNotFound:        "USER_NOT_FOUND",
}
