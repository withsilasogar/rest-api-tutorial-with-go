package routes

import (
	"test01/cmd/server"
	"test01/internals/handler"
	"test01/x/interfacesx"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterUserRoutes(server server.GinServer, userHandler *handler.UserHandler) {
	server.RegisterGroupRoute("/api/v1/user", []interfacesx.RouteDefinition{
		{Method: "POST", Path: "/register", Handler: userHandler.CreateUser},
		{Method: "GET", Path: "/:email", Handler: userHandler.GetUser},
	}, func(c *gin.Context) {
		//You can add your middlewares here
		logrus.Infof("Request on %s", c.Request.URL.Path)
	})
}
