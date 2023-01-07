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
		return "illegal param"
	case ErrDuplicateKey:
		return "duplicate key"
	case ErrResourceUnavailable:
		return "resource unavailable"
	case ErrResourceNotFound:
		return "resource notFound"
	case ErrInternal:
		return "internal error"
	default:
		return "unknown error"
	}

}

type Error struct {
	errMsg  string
	errType ErrorType
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
	for _, parser := range Parsers.Iter() {
		if parser.Support(err) {
			msg, t := parser.Do(err)
			return New(msg, t)
		}
	}
	if e, ok := err.(error); ok {
		return New(e.Error(), ErrUnknown)
	} else {
		return New(fmt.Sprintf("+%v", err), ErrUnknown)
	}
}
