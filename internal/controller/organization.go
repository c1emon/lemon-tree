package controller

import (
	"github.com/c1emon/lemontree/internal/service"
	"github.com/c1emon/lemontree/pkg/httpx"
	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	service service.OrganizationService
}

func (s *OrganizationController) Create(c *gin.Context) {
	//q := httpx.PaginationFromQuery(c.Request)
}

func (s *OrganizationController) GetAll(c *gin.Context) {
	name := c.Query("name")
	pagination := httpx.PaginationFromQuery(c.Request)
	orgs := s.service.GetByNameAll(pagination, name)

	c.JSON(200, httpx.ResponseOK().WithPagination(pagination).WithData(orgs))
}

func BuildRouter() {

}
