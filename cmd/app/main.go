package main

import (
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	"github.com/tidwall/buntdb"

	"github.com/umarquez/cryptocoins-go-challenge/internal/controller"
	"github.com/umarquez/cryptocoins-go-challenge/internal/repository"
	"github.com/umarquez/cryptocoins-go-challenge/internal/service"
	"github.com/umarquez/cryptocoins-go-challenge/internal/usecase"

	_ "github.com/umarquez/cryptocoins-go-challenge/docs" // Import generated docs
)

const dataPath = "./data"

// @title CryptoCoins API
// @version 1.0
// @description This is a sample server for managing cryptocurrencies.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	m := new(sync.Mutex)

	_ = os.Mkdir(dataPath, os.ModePerm)
	dbCnn, err := buntdb.Open(path.Join(dataPath, "cryptos.db"))
	if err != nil {
		panic(err)
	}

	defer dbCnn.Close()
	cryptoRepo := repository.NewCryptoRepository(dbCnn, m, time.Minute)
	cryptoService := service.GetCryptoService()
	cryptoUseCase := usecase.NewCryptoUseCase(cryptoService, cryptoRepo)

	err = http.ListenAndServe(":8080", controller.NewRouter(cryptoUseCase))
	if err != nil {
		panic(err)
	}
}
