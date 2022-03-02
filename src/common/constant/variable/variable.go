package variable

import "net/http"

// HttpMethod http method enum
type HttpMethod string

const (
	GET    = HttpMethod(http.MethodGet)
	POST   = HttpMethod(http.MethodPost)
	PUT    = HttpMethod(http.MethodPut)
	DELETE = HttpMethod(http.MethodDelete)
	OPTION = HttpMethod(http.MethodOptions)
)
