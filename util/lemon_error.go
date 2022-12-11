package util

type LemonErrorType int

const (
	ResNotFound LemonErrorType = iota
	DupKeyError
	InterError
)

type LemonError struct {
	t   LemonErrorType
	msg string
	res string
}

func (e *LemonError) Error() string {

	return e.msg
}

func LemonNotFoundError(res string) *LemonError {
	return &LemonError{
		t:   ResNotFound,
		res: res,
		msg: "res not found",
	}
}
