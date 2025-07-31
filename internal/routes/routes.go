package routes

import (
	"dreamus/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	controller := controllers.NewBuildController()

	buildGroup := r.Group("/builds")
	{
		buildGroup.POST("/", controller.Create)
		buildGroup.GET("/:id", controller.Get)
		buildGroup.PUT("/:id", controller.Update)
		buildGroup.DELETE("/:id", controller.Delete)
	}
	return r
}
