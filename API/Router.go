package API

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouteConfig() *gin.Engine{

	r := gin.Default()

	r.Use(func(context *gin.Context) {
		fmt.Println("Apply middleware here")
	})

	return r
}
