// Package server provee las funcionas para levantar el servidor basado en el framework web gin-gonic
package server // import "github.com/aferrercrafter/operacion-fuego-quazar/api/server"

import (
	"github.com/aferrercrafter/operacion-fuego-quazar/api/server/controllers"
	"github.com/aferrercrafter/operacion-fuego-quazar/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// Start server
func Start() {

	router := SetupRouter()
	router.Run(":8080")
}

// SetupRouter set topsecret routes
func SetupRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	docs.SwaggerInfo.Title = "operacion-quazar API"
	docs.SwaggerInfo.Description = "Servicio topsecret para operación fuego quazar."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "fuego-quazar.rj.r.appspot.com"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		topsecret := v1.Group("/topsecret")
		{
			topsecret.POST("", controllers.TopSecretPost)
		}

		topsecretsplit := v1.Group("/topsecret_split")
		{
			topsecretsplit.POST(":name", controllers.TopSecretSplitPost)
			topsecretsplit.GET("", controllers.TopSecretSplitGet)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router

}
