package test

import (
	"fmt"
	"github.com/c1emon/lemontree/errorc"
	"github.com/c1emon/lemontree/ginx"
	"github.com/c1emon/lemontree/log"
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_DbCreate(t *testing.T) {
	//config.SetConfig(8080, "postgres", "host=10.0.0.70 port=5432 user=postgres dbname=lemon_tree password=123456 sslmode=disable")
	log.Init("info")

	r := ginx.GetGinEngine()

	r.GET("/ping/:ok", func(c *gin.Context) {
		//c.JSON(400, gin.H{
		//	"message": "pong",
		//})
		//c.Error(errors.New("errx"))
		switch c.Param("ok") {
		case "t":

			c.JSON(200, gin.H{
				"message": "pong",
			})
		case "f":
			c.Error(errorc.ErrResourceUnavailable)
		default:
			panic("panic")
		}
		return
	})

	r.Run(fmt.Sprintf(":8080"))

}
