package util

type HTTPError struct {
	Code int
	Text string
}

// List with errors
var errors map[int]string

func init() {
	errors = make(map[int]string)

	// Now create the error list
	errors[400] = "Bad Request"
	errors[401] = "Unauthorized"
	errors[403] = "Forbidden"
	errors[404] = "Not Found"
	errors[405] = "Method Not Allowed"
	errors[406] = "Not Acceptable"
	errors[408] = "Request Time-out"
}

// ### Public
func GetHTTPError(code int) *HTTPError {

	hErr := HTTPError{
		Code: code,
		Text: errors[code],
	}

	return &hErr
}
