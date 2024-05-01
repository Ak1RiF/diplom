package app

import (
	"github.com/gin-gonic/gin"
	"github.com/todoApp/pkg/handlers"
	"github.com/todoApp/pkg/repository"
	"github.com/todoApp/pkg/service"
)

func Start() {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handlers.NewHandler(service)

	// paths
	router := gin.Default()

	router.POST("/sign-in", handler.SignIn)
	router.POST("/sign-up", handler.SignUp)

	secure := router.Group("/protected", handler.AuthMiddleware())
	{
		secure.GET("/greet", func(ctx *gin.Context) {
			ctx.JSON(200, "Hello authorize user!")
		})
	}

	router.Run(":8000")
}
