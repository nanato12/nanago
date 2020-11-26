package nana

import (
	"bytes"
	"fmt"
)

// APIError type
type APIError struct {
	Code     int
	Response *ErrorResponse
}

// Error method
func (e *APIError) Error() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "nanago: APIError %d ", e.Code)
	if e.Response != nil {
		fmt.Fprintf(&buf, "\n[code: %d] %s", e.Response.Data.Code, e.Response.Data.Message)
	}
	return buf.String()
}
