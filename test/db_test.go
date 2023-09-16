package test

import (
	"fmt"
	"github.com/c1emon/lemontree/ginx"
	"github.com/c1emon/lemontree/httpx"
	"github.com/gin-gonic/gin"
	"testing"
)

func SetPage(pageable httpx.Pageable) {
	pageable.SetTotal(109)
}

func Test_DbCreate(t *testing.T) {
	start()
	r := ginx.GetGinEngine()

	r.GET("/page", func(c *gin.Context) {
		pagination := httpx.PaginationFromQuery(c.Request)
		SetPage(pagination)

		c.JSON(200, httpx.ResponseOK().WithPagination(pagination))
	})

	r.Run(fmt.Sprintf(":8080"))

}

func Test_Sort(t *testing.T) {
}
