package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/umarquez/cryptocoins-go-challenge/internal/domain"
	"github.com/umarquez/cryptocoins-go-challenge/internal/dto"
)

type CryptoUseCase interface {
	GetAllCryptos() ([]domain.Crypto, error)
	GetCryptoById(id int) (domain.Crypto, error)
}

type CryptoController interface {
	GetCryptos(*gin.Context)
	GetCryptoById(*gin.Context)
}

type cryptoController struct {
	cryptoUseCase CryptoUseCase
}

func NewCryptoController(cryptoUseCase CryptoUseCase) CryptoController {
	return &cryptoController{
		cryptoUseCase: cryptoUseCase,
	}
}

// GetCryptos godoc
// @Summary Get all cryptos
// @Description Returns all cryptocurrencies with normalized data.
// @Tags cryptocoin
// @Accept json
// @Produce json
// @Success 200 {array} dto.NormalizedCrypto
// @Failure 500 {object} nil
// @Router /cryptos [get]
func (cc *cryptoController) GetCryptos(ctx *gin.Context) {
	cryptos, err := cc.cryptoUseCase.GetAllCryptos()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var normalizedCryptos []dto.NormalizedCrypto
	for _, crypto := range cryptos {
		nCrypto, err := dto.NormalizeCrypto(crypto)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			log.Println(fmt.Errorf("error normalizing crypto (%v): %v", crypto.TickerSymbol, err))
			return
		}
		normalizedCryptos = append(normalizedCryptos, nCrypto)
	}

	ctx.JSON(http.StatusOK, normalizedCryptos)
}

// GetCryptoById godoc
// @Summary Get crypto by id
// @Description Returns a cryptocurrency with normalized data.
// @Tags cryptocoin
// @Accept json
// @Produce json
// @Param id path int true "Crypto ID"
// @Success 200 {object} dto.NormalizedCrypto
// @Failure 400 {object} nil
// @Failure 500 {object} nil
// @Router /cryptos/{id} [get]
func (cc *cryptoController) GetCryptoById(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c, err := cc.cryptoUseCase.GetCryptoById(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	normalizedCrypto, err := dto.NormalizeCrypto(c)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		log.Println(fmt.Errorf("error normalizing crypto (%v): %v", c.TickerSymbol, err))
		return
	}

	ctx.JSON(http.StatusOK, normalizedCrypto)
	return
}
