package util

type LemonErrorType int

const (
	ResourceNotFound LemonErrorType = iota + 1000
	DuplicateKey
	InternalError
	ResourceUnavailable
)

type LemonError struct {
	Type LemonErrorType
	Msg  string
	Res  string
}

func (e *LemonError) Error() string {

	return e.Msg
}

func LemonNotFoundError(res string) *LemonError {
	return &LemonError{
		Type: ResourceNotFound,
		Res:  res,
		Msg:  "res not found",
	}
}
