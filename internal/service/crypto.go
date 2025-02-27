package service

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/umarquez/cryptocoins-go-challenge/internal/domain"
	"github.com/umarquez/cryptocoins-go-challenge/sdk/bitso_client"
)

// CryptoService represents the service for the crypto domain.
type CryptoService interface {
	GetCryptoValue(crypto domain.CryptoCurrency, currency domain.Currency) (string, error)
}

type cacheItem struct {
	value      string
	expiration time.Time
}

type cryptoService struct {
	cache map[string]cacheItem
}

var cryptoServiceInstance *cryptoService

// GetCryptoService returns a single instance of the crypto service.
func GetCryptoService() CryptoService {
	if cryptoServiceInstance == nil {
		cryptoServiceInstance = &cryptoService{
			cache: make(map[string]cacheItem),
		}
	}

	return cryptoServiceInstance
}

func (s *cryptoService) GetCryptoValue(crypto domain.CryptoCurrency, currency domain.Currency) (string, error) {
	// Simulate a delay to simulate the time it takes to fetch the data.
	delay := time.Duration(rand.Intn(4500)+500) * time.Millisecond // from 0.5 to 5 seconds
	time.Sleep(delay)

	// Check if the value is in the cache.
	cacheKey := fmt.Sprintf("%s_%s", crypto, currency)
	if item, ok := s.cache[cacheKey]; ok && item.expiration.After(time.Now()) {
		return item.value, nil
	}

	// Fetch the value from the API.
	c := bitso_client.NewClient(false)

	// Retry fetching the data up to 3 times in case of an error.
	for retriesCount := 0; retriesCount < 3; retriesCount++ {
		ticker, err := c.GetTicker(bitso_client.TickerName(fmt.Sprintf("%s_%s", strings.ToLower(string(crypto)), strings.ToLower(string(currency)))))
		if err != nil {
			fmt.Printf("(%v) Error fetching crypto_service value: %v\n", retriesCount, err)
			continue
		}

		// Store the value in the cache.
		s.cache[cacheKey] = cacheItem{
			value:      ticker.Payload.Last,
			expiration: time.Now().Add(time.Minute),
		}

		return ticker.Payload.Last, nil
	}

	return "", fmt.Errorf("failed to fetch crypto_service value after 3 retries")
}
