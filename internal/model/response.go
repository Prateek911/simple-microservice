package model

import "fmt"

type BaseApiError interface {
	Type() ErrorType
	ToError() error
	Error() string
	IsNil() bool
}

type ApiError struct {
	Typ ErrorType
	Err error
}

func (a *ApiError) Type() ErrorType {
	return a.Typ
}

func (a *ApiError) ToError() error {
	if a != nil {
		return a.Err
	}
	return a
}

func (a *ApiError) Error() string {
	if a == nil || a.Err == nil {
		return ""
	}
	return a.Err.Error()
}

func (a *ApiError) IsNil() bool {
	return a == nil || a.Err == nil
}

type ErrorType string

const (
	ErrorNone           ErrorType = ""
	ErrorTimeout        ErrorType = "timeout"
	ErrorCanceled       ErrorType = "canceled"
	ErrorExec           ErrorType = "execution"
	ErrorBadData        ErrorType = "bad_data"
	ErrorInternal       ErrorType = "internal"
	ErrorUnavailable    ErrorType = "unavailable"
	ErrorNotFound       ErrorType = "not_found"
	ErrorNotImplemented ErrorType = "not_implemented"
	ErrorUnauthorized   ErrorType = "unauthorized"
	ErrorForbidden      ErrorType = "forbidden"
	ErrorConflict       ErrorType = "conflict"
)

func BadRequest(err error) *ApiError {
	return &ApiError{
		Typ: ErrorBadData,
		Err: err,
	}
}

func TimeOut(err error) *ApiError {
	return &ApiError{
		Typ: ErrorTimeout,
		Err: err,
	}
}

func BadRequestStr(s string) *ApiError {
	return &ApiError{
		Typ: ErrorBadData,
		Err: fmt.Errorf(s),
	}
}

func InternalError(err error) *ApiError {
	return &ApiError{
		Typ: ErrorInternal,
		Err: err,
	}
}

func NotFoundError(err error) *ApiError {
	return &ApiError{
		Typ: ErrorNotFound,
		Err: err,
	}
}

func UnauthorizedError(err error) *ApiError {
	return &ApiError{
		Typ: ErrorUnauthorized,
		Err: err,
	}
}

func UnavailableError(err error) *ApiError {
	return &ApiError{
		Typ: ErrorUnavailable,
		Err: err,
	}
}

func ForbiddenError(err error) *ApiError {
	return &ApiError{
		Typ: ErrorForbidden,
		Err: err,
	}
}
