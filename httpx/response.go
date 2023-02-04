package httpx

import "strings"

type Order int

const (
	ASC  Order = 1
	DESC Order = -1
)

func (o Order) String() string {

	if o == ASC {
		return "ASC"
	} else {
		return "DESC"
	}
}

func ParserOrder(val any) Order {
	switch v := val.(type) {
	case int:
		if v < 0 {
			return DESC
		}
	case string:
		if strings.TrimSpace(strings.ToUpper(v)) == DESC.String() {
			return DESC
		}
	case bool:
		if !v {
			return DESC
		}
	}
	return ASC
}

type SortItem struct {
	Field string
	Order Order
}

// ?sort=0,field0&sort=1,field1,DESC
type PaginationParams struct {
	Size  int        `form:"size" json:"size,omitempty" binding:"required_with=page total"`
	Page  int        `form:"page" json:"page,omitempty" binding:"required_with=size total"`
	Total int        `form:"total" json:"total,omitempty" binding:"required_with=page size"`
	Sorts []SortItem `json:"-" binding:"-"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
	PaginationParams
}

func (r *Response) WithMessage(msg string) *Response {
	r.Message = msg
	return r
}

func (r *Response) WithError(err string) *Response {
	r.Error = err
	return r
}

func (r *Response) WithData(data any) *Response {
	r.Data = data
	return r
}

func NewResponse(c int) *Response {
	return &Response{Code: c}
}

func ResponseOK() *Response {
	return NewResponse(0)
}
