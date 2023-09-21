package test

import (
	"fmt"
	"testing"

	"github.com/c1emon/lemontree/pkg/ginx"
	"github.com/c1emon/lemontree/pkg/httpx"
	"github.com/gin-gonic/gin"
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
