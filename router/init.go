package router

import "github.com/gin-gonic/gin"

func RouterInit() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {})

	r.Run(":8080")
}
