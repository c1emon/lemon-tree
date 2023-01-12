package errorc

import (
	"github.com/pkg/errors"
)

var (
	ErrInternal            = New(1000, 500, "internal error")
	ErrUnknown             = New(1001, 500, "unknown error")
	ErrIllegalParam        = New(1000, 500, "internal error")
	ErrDuplicateKey        = New(1000, 500, "internal error")
	ErrResourceUnavailable = New(1000, 500, "internal error")
	ErrResourceNotFound    = New(1000, 500, "internal error")
)

type Error struct {
	code       int
	httpStatus int
	message    string
}

func New(code, httpStatus int, message string) *Error {
	return &Error{code: code, httpStatus: httpStatus, message: message}
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Code() int {
	return e.code
}

func (e Error) HttpStatus() int {
	return e.httpStatus
}

func (e Error) Is(err error) bool {
	if x, ok := errors.Cause(err).(Error); ok {
		return e.Code() == x.Code()
	}
	return false
}

func From(err error) *Error {
	if err == nil {
		return nil
	}

	for _, parser := range Parsers.Iter() {
		if parser.Support(err) {
			return parser.Parse(err)
		}
	}

	return ErrUnknown
}
