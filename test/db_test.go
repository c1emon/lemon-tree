package test

import (
	"fmt"
	"github.com/c1emon/lemontree/model"
	"testing"
)

func Test_DbCreate(t *testing.T) {

	d := model.CreateDepartment()
	fmt.Printf("%+v", d)

}
