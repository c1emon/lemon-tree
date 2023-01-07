package test

import (
	"fmt"
	"github.com/c1emon/lemontree/errorc"
	"github.com/pkg/errors"
	"testing"
)

func Dao() error {
	return errors.Wrap(errorc.New("", 0), "not found")
}

func Service() error {
	err := Dao()
	if err != nil {
		return errors.WithMessage(err, "service xxxx")
	}
	return nil
}

func Ctrl() error {
	err := Service()
	return errors.WithMessage(err, "controller xxxxx")
}

func Test_Err(t *testing.T) {
	err := Ctrl()
	if errorc.Is(err, errorc.ErrInternal) {
		//fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)

	}

}
