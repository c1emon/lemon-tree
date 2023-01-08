package errorc

import (
	"fmt"
	"github.com/pkg/errors"
)

type ErrorType int

const (
	ErrInternal ErrorType = iota - 1
	ErrUnknown
	ErrIllegalParam
	ErrDuplicateKey
	ErrResourceUnavailable
	ErrResourceNotFound
)

func (t ErrorType) String() string {
	switch t {
	case ErrIllegalParam:
		return "illegal param error"
	case ErrDuplicateKey:
		return "duplicate key error"
	case ErrResourceUnavailable:
		return "resource unavailable error"
	case ErrResourceNotFound:
		return "resource notFound error"
	case ErrInternal:
		return "internal error"
	default:
		return "unknown error"
	}

}

type Error struct {
	errMsg  string
	errType ErrorType
	from    error
}

func (e Error) Error() string {
	return e.errMsg
}

func (e Error) Type() ErrorType {
	return e.errType
}

func (e Error) TypeIs(t ErrorType) bool {
	return e.errType == t
}

func New(msg string, t ErrorType) *Error {
	return &Error{errMsg: msg, errType: t}
}

func Is(err error, t ErrorType) bool {
	if e, ok := errors.Cause(err).(Error); ok {
		return e.TypeIs(t)
	}
	return false
}

func From(err any) *Error {
	if err == nil {
		return nil
	}
	ec := &Error{}
	if e, ok := err.(error); ok {
		ec.from = e
	} else {
		ec.from = fmt.Errorf("%+v", err)
	}

	for _, parser := range Parsers.Iter() {
		if parser.Support(err) {
			ec.errMsg, ec.errType = parser.Do(err)
			return ec
		}
	}

	ec.errType = ErrUnknown
	if e, ok := err.(error); ok {
		ec.errMsg = e.Error()
	} else {
		ec.errMsg = fmt.Sprintf("%s", err)
	}
	return ec
}
