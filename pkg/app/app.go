package app

import (
	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
	}))

	router.POST("/sign-in", handler.SignIn)
	router.POST("/sign-up", handler.SignUp)

	api := router.Group("/api", handler.AuthMiddleware())
	{
		api.POST("/logout", handler.LogOut)
		account := api.Group("/account")
		{
			account.GET("/info", handler.GetInfo)

			pets := account.Group("/pets")
			{
				pets.GET("/", handler.GetPets)
				pets.GET("/:id", handler.GetPetsById)
				pets.POST("/:id", handler.PostPets)
			}

			eggs := account.Group("/eggs")
			{
				eggs.GET("/", handler.GetEggs)
				eggs.POST("/:id", handler.PostEggs)
				eggs.POST("/add/:id", handler.AddToCountEggs)
				eggs.POST("/remove/:id", handler.TakeFromCountEggs)
			}

			quests := account.Group("/quests")
			{
				quests.GET("/", handler.AllQuests)
				quests.GET("/:id", handler.ByIdQuest)
				quests.POST("/", handler.PostQuest)
				quests.PUT("/:id", handler.PutQuest)
				quests.DELETE("/:id", handler.DeleteQuest)
			}

		}
	}

	router.Run(":8000")
}
