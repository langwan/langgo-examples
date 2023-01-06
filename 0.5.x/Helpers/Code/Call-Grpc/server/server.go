package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/langwan/langgo"
	helper_code "github.com/langwan/langgo/helpers/code"
	helperGin "github.com/langwan/langgo/helpers/gin"
	"io"
	"langwan/langgo-examples/Helpers/Code/Call-Grpc/server/service/server"
)

const addr = "localhost:8000"

func main() {
	langgo.Run()
	g := gin.New()
	api := g.Group("api")
	api.Any("/*uri", requestProxy())
	g.Run()
}

func requestProxy() gin.HandlerFunc {
	server := server.Server{}

	return func(c *gin.Context) {
		methodName := c.Param("uri")[1:]

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			return
		}

		if err != nil {
			c.AbortWithStatus(500)
			return
		}

		response, code, err := helper_code.Call(context.Background(), &server, methodName, string(body))

		if err != nil {
			c.AbortWithStatus(500)
			return
		} else if code != 0 {
			helperGin.SendBad(c, code, err.Error(), nil)
		}

		helperGin.SendOk(c, response)
	}
}
