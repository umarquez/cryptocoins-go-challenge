package dto

import (
	"strings"

	"github.com/davecgh/go-spew/spew"

	"github.com/umarquez/cryptocoins-go-challenge/internal/domain"
)

type NormalizedCrypto struct {
	Id        int           `json:"id"`
	Component string        `json:"component"`
	Model     domain.Crypto `json:"model"`
}

func NormalizeCrypto(crypto domain.Crypto) (NormalizedCrypto, error) {
	id, ok := domain.CryptoIdEnum[domain.CryptoCurrency(crypto.TickerSymbol)]
	if !ok {
		return NormalizedCrypto{}, domain.ErrCryptoIdNotFound
	}
	return NormalizedCrypto{
		Id:        id,
		Component: spew.Sprintf("crypto_%v", strings.ToLower(string(crypto.TickerSymbol))),
		Model:     crypto,
	}, nil
}
