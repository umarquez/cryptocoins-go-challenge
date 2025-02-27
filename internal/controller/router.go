package controller

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(useCase CryptoUseCase) *gin.Engine {
	cryptoController := NewCryptoController(useCase)
	router := gin.Default()
	router.GET("/cryptos", cryptoController.GetCryptos)
	router.GET("/cryptos/:id", cryptoController.GetCrypto)

	return router
}
