package test

import (
	"fmt"
	"github.com/c1emon/lemontree/config"
	"github.com/c1emon/lemontree/log"
	"github.com/c1emon/lemontree/model"
	"testing"
)

func Hello(o *model.Organization) {
	fmt.Println(o)
}

func Test_DbCreate(t *testing.T) {
	config.SetConfig(8080, "postgres", "host=10.0.0.70 port=5432 user=postgres dbname=lemon_tree password=123456 sslmode=disable")
	log.Init("trace")

	//r.CreateOne(context.Background(), &model.Organization{Name: "taaest"})

}
