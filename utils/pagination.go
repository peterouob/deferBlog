package utils

import (
	"blog/config"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		panic(err)
	}
	if page > 0 {
		result = (page - 1) * config.Config.GetInt("app.page_size")
	}
	return result
}
