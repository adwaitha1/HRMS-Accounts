package customError

import (
	"fmt"
	"net/http"
)

// CustomError represents a custom error type with an additional HTTP status code.
type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

// DatabaseError represents an error from the database layer.
type DatabaseError struct {
	Reason string
}

func (e *DatabaseError) Error() string {
	return e.Reason
}

// ServiceError represents an error from the service layer.
type ServiceError struct {
	Reason string
}

func (e *ServiceError) Error() string {
	return e.Reason
}

// AnotherServiceError represents another type of error from the service layer.
type AnotherServiceError struct {
	Details string
}

func (e *AnotherServiceError) Error() string {
	return e.Details
}

// MapErrorToHTTPCode returns the appropriate HTTP status code for a given error.
func MapErrorToHTTPCode(err error) int {
	switch e := err.(type) {
	case *CustomError:
		return e.Code
	case *DatabaseError:
		return http.StatusInternalServerError
	case *ServiceError:
		return http.StatusBadRequest
	case *AnotherServiceError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

// HandleError formats and returns an HTTP error response.
func HandleError(err error) (int, string) {
	statusCode := MapErrorToHTTPCode(err)
	errorMessage := err.Error()

	return statusCode, fmt.Sprintf(`{"error": "%s"}`, errorMessage)
}
