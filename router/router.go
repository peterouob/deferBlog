package router

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	user := r.Group("/user")
	{
		//all tags
		user.GET("/tags")
		//insert tags
		user.POST("/tags")
		//update tags
		user.PUT("/tags/:id")
		//delete tags
		user.DELETE("/tags/:id")
	}
}
