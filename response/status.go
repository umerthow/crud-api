package response

// Collection of status.
const (
	StatOK                 string = "OK"
	StatCreated            string = "CREATED"
	StatNotFound           string = "NOT_FOUND"
	StatUnexpectedError    string = "UNEXPECTED_ERROR"
	StatInsufficientPoint  string = "INSUFFICIENT_POINT"
	StatusInvalidPayload   string = "INVALID_PAYLOAD"
	StatusInvalidParameter string = "INVALID_PARAMETER"
	StatUnauthorized       string = "UNAUTHORIZED"
	StatTokenExpired       string = "EXPIRED"
	StatAlreadyExist       string = "ALREADY_EXIST"
	StatBadRequest         string = "BAD_REQUEST"
	StatForbidden          string = "FORBIDDEN"
	StatRequestTimeout     string = "REQUEST_TIMEOUT"
)
