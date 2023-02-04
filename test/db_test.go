package test

import (
	"fmt"
	"github.com/c1emon/lemontree/ginx"
	"github.com/c1emon/lemontree/httpx"
	"github.com/c1emon/lemontree/log"
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_DbCreate(t *testing.T) {
	//config.SetConfig(8080, "postgres", "host=10.0.0.70 port=5432 user=postgres dbname=lemon_tree password=123456 sslmode=disable")
	log.Init("info")

	r := ginx.GetGinEngine()

	r.GET("/page", func(c *gin.Context) {
		q := httpx.PaginationFromQuery(c.Request)
		fmt.Printf("%+v", q)
		return
	})

	r.Run(fmt.Sprintf(":8080"))

}

func Test_Sort(t *testing.T) {
}
