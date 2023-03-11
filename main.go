package main

import (
	"blog/config"
	"blog/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	router.Router(r)
	r.Run(config.Config.GetString("server.port"))
}
