package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/umarquez/cryptocoins-go-challenge/internal/domain"
	"github.com/umarquez/cryptocoins-go-challenge/internal/dto"
)

type CryptoUseCase interface {
	GetCryptos() ([]domain.Crypto, error)
}

type CryptoController interface {
	GetCryptos(*gin.Context)
	GetCrypto(*gin.Context)
}

type cryptoController struct {
	cryptoUseCase CryptoUseCase
}

func NewCryptoController(cryptoUseCase CryptoUseCase) CryptoController {
	return &cryptoController{
		cryptoUseCase: cryptoUseCase,
	}
}

func (cc *cryptoController) GetCryptos(ctx *gin.Context) {
	cryptos, err := cc.cryptoUseCase.GetCryptos()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Normalize the cryptos.
	var normalizedCryptos []dto.NormalizedCrypto
	for _, crypto := range cryptos {
		normalizedCryptos = append(normalizedCryptos, dto.NormalizeCrypto(crypto))
	}

	ctx.JSON(http.StatusOK, normalizedCryptos)
}

func (cc *cryptoController) GetCrypto(ctx *gin.Context) {
	cryptos, err := cc.cryptoUseCase.GetCryptos()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// retrieve the id from the url
	id := ctx.Param("id")
	for _, c := range cryptos {
		cryptoId := domain.IdByCryptoTable[c.TickerSymbol]
		nid, err := strconv.Atoi(id)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if cryptoId != nid {
			continue
		}

		normalizedCrypto := dto.NormalizeCrypto(c)
		ctx.JSON(http.StatusOK, normalizedCrypto)
		return
	}

	ctx.AbortWithStatus(http.StatusNotFound)
}
