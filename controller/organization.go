package controller

import (
	"github.com/c1emon/lemontree/service"
	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	service service.OrganizationService
}

func (s *OrganizationController) Create(c *gin.Context) {

	c.JSON(1, s.service.Create())
}

func BuildRouter() {

}
