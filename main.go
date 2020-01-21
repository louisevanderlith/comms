package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/comms/controllers/message"
	"github.com/louisevanderlith/comms/core"
	"github.com/louisevanderlith/droxo"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	r := gin.Default()

	r.GET("/message/:key", message.View)

	msgs := r.Group("/message")
	msgs.Use(droxo.Authorize())
	msgs.POST("", message.Create)

	r.GET("/messages", message.Get)
	r.GET("/messages/:pagesize/*hash", message.Search)

	err := r.Run(":8085")

	if err != nil {
		panic(err)
	}
}
