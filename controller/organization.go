package controller

import (
	"github.com/c1emon/lemontree/httpx"
	"github.com/c1emon/lemontree/service"
	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	service service.OrganizationService
}

func (s *OrganizationController) Create(c *gin.Context) {
	q := httpx.PaginationFromQuery(c.Request)
}

func BuildRouter() {

}
