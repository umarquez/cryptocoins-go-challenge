package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/umarquez/cryptocoins-go-challenge/docs"
)

func NewRouter(crypto CryptoUseCase) *gin.Engine {
	cryptoController := NewCryptoController(crypto)
	router := gin.Default()

	// Configura el endpoint para servir Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		cryptos := api.Group("/cryptos")
		{
			cryptos.GET("/", cryptoController.GetCryptos)
			cryptos.GET("/:id", cryptoController.GetCryptoById)
		}
	}

	return router
}
