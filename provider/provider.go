package provider

import (
	"test01/cmd/server"
	"test01/internals/handler"
	"test01/internals/repository"
	"test01/internals/routes"
	"test01/internals/services"

	"gorm.io/gorm"
)

func NewProvider(db *gorm.DB, server server.GinServer) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	routes.RegisterUserRoutes(server, userHandler)
}
