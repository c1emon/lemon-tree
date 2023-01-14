package test

import (
	"fmt"
	"github.com/c1emon/lemontree/errorc"
	"github.com/c1emon/lemontree/errorc/parser"
	"gorm.io/gorm"
	"testing"
)

func Test_Err(t *testing.T) {
	parser.NewGormParser()
	err := errorc.From(gorm.ErrRecordNotFound)
	fmt.Printf("%s", err.(error))
}
