package main

import (
	"net/http"

	"github.com/umarquez/cryptocoins-go-challenge/internal/controller"
	"github.com/umarquez/cryptocoins-go-challenge/internal/usecase"
)

func main() {
	err := http.ListenAndServe(":8080", controller.NewRouter(usecase.NewCryptoUseCase()))
	if err != nil {
		panic(err)
	}
}
