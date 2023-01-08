package parser

import (
	"errors"
	"github.com/c1emon/lemontree/errorc"
	"gorm.io/gorm"
)

type GormParser struct {
}

func (p *GormParser) Do(err any) (string, errorc.ErrorType) {
	if e, ok := err.(error); ok {
		switch {
		case errors.Is(e, gorm.ErrRecordNotFound):
			return "", errorc.ErrResourceNotFound
		}
	}
	return "", errorc.ErrUnknown
}

func (p *GormParser) Support(err any) bool {
	if e, ok := err.(error); ok {
		switch {
		case errors.Is(e, gorm.ErrRecordNotFound), errors.Is(e, gorm.ErrNotImplemented):
			return true
		}
	}
	return false
}

func NewGormParser() *GormParser {
	p := &GormParser{}
	errorc.Parsers.Add("gorm", p)
	return p
}
