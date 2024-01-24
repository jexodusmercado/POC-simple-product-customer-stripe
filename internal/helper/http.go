package helper

import (
	"fmt"
	"net/http"
)

type IDQueryString struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type ApiError struct {
	Field string
	Error string
}

type HTTPError struct {
	Code            int        `json:"code"`
	Message         string     `json:"message,omitempty"`
	Messages        []ApiError `json:"messages,omitempty"`
	InternalError   error      `json:"-"`
	InternalMessage string     `json:"-"`
	ErrorID         string     `json:"error_id,omitempty"`
}

func (e *HTTPError) Error() string {
	if e.InternalMessage != "" {
		return e.InternalMessage
	}
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func (e *HTTPError) Is(target error) bool {
	return e.Error() == target.Error()
}

// Cause returns the root cause error
func (e *HTTPError) Cause() error {
	if e.InternalError != nil {
		return e.InternalError
	}
	return e
}

// WithInternalError adds internal error information to the error
func (e *HTTPError) WithInternalError(err error) *HTTPError {
	e.InternalError = err
	return e
}

// WithInternalMessage adds internal message information to the error
func (e *HTTPError) WithInternalMessage(fmtString string, args ...interface{}) *HTTPError {
	e.InternalMessage = fmt.Sprintf(fmtString, args...)
	return e
}

func httpStruct(code int, fmtString string, args ...interface{}) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: fmt.Sprintf(fmtString, args...),
	}
}

func BadRequestError(fmtString string, args ...interface{}) *HTTPError {
	return httpStruct(http.StatusBadRequest, fmtString, args...)
}

func InternalServerError(fmtString string, args ...interface{}) *HTTPError {
	return httpStruct(http.StatusInternalServerError, fmtString, args...)
}

func NotFoundError(fmtString string, args ...interface{}) *HTTPError {
	return httpStruct(http.StatusNotFound, fmtString, args...)
}

func ExpiredTokenError(fmtString string, args ...interface{}) *HTTPError {
	return httpStruct(http.StatusUnauthorized, fmtString, args...)
}

func UnauthorizedError(fmtString string, args ...interface{}) *HTTPError {
	return httpStruct(http.StatusUnauthorized, fmtString, args...)
}

func ForbiddenError(fmtString string, args ...interface{}) *HTTPError {
	return httpStruct(http.StatusForbidden, fmtString, args...)
}

func UnprocessableEntityError(fmtString string, args ...interface{}) *HTTPError {
	return httpStruct(http.StatusUnprocessableEntity, fmtString, args...)
}

func TooManyRequestsError(fmtString string, args ...interface{}) *HTTPError {
	return httpStruct(http.StatusTooManyRequests, fmtString, args...)
}

func ConflictError(fmtString string, args ...interface{}) *HTTPError {
	return httpStruct(http.StatusConflict, fmtString, args...)
}

func NotImplemented(fmtString string, args ...interface{}) *HTTPError {
	return httpStruct(http.StatusNotImplemented, fmtString, args...)
}

func CreatedOK(fmtString string, args ...interface{}) *HTTPError {
	return httpStruct(http.StatusCreated, fmtString, args...)
}

func OK(fmtString string, args ...interface{}) *HTTPError {
	return httpStruct(http.StatusOK, fmtString, args...)
}

func NoContent(fmtString string, args ...interface{}) *HTTPError {
	return httpStruct(http.StatusNoContent, fmtString, args...)
}
