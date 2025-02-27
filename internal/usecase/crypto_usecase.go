package usecase

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/umarquez/cryptocoins-go-challenge/internal/domain"
	"github.com/umarquez/cryptocoins-go-challenge/internal/service/crypto_service"
)

// CryptoUseCase defines the contract for crypto_service business logic.
type CryptoUseCase interface {
	GetCryptos() ([]domain.Crypto, error)
}

type cryptoUseCase struct {
}

func NewCryptoUseCase() CryptoUseCase {
	return &cryptoUseCase{}
}

// GetCryptos retrieves the list of cryptocurrencies with their average prices.
func (uc *cryptoUseCase) GetCryptos() ([]domain.Crypto, error) {
	// Create a map to store the results.
	results := make(map[domain.CryptoCurrency]map[domain.Currency]string)
	resultsWriterMutex := new(sync.Mutex)

	// Create a wait group to wait for all goroutines to finish.
	wg := new(sync.WaitGroup)
	for _, crypto := range domain.Cryptos {
		for _, currency := range domain.Currencies {
			if _, ok := results[crypto]; !ok {
				results[crypto] = make(map[domain.Currency]string)
			}

			wg.Add(1)
			go uc.getCryptoValueAsync(crypto, currency, wg, results, resultsWriterMutex)
		}
	}

	wg.Wait()
	fmt.Printf("Results: %v\n", results)

	cryptoList := []domain.Crypto{}
	for simbol, prices := range results {
		c := domain.Crypto{
			Date:         time.Now(),
			Name:         domain.CryptoName[simbol],
			TickerSymbol: string(simbol),
			Price: domain.Price{
				USD: prices[domain.USD],
				MXN: prices[domain.MXN],
			},
		}
		cryptoList = append(cryptoList, c)
	}

	return cryptoList, nil
}

// getCryptoValueAsync retrieves the average price of a cryptocurrency in a given currency.
func (uc *cryptoUseCase) getCryptoValueAsync(crypto domain.CryptoCurrency, currency domain.Currency, wg *sync.WaitGroup, results map[domain.CryptoCurrency]map[domain.Currency]string, resultsWriterMutex *sync.Mutex) {
	log.Printf("[%s][%s] thread started\n", crypto, currency)
	startTime := time.Now()
	defer wg.Done()

	value, err := crypto_service.GetCryptoValue(crypto, currency)
	if err != nil {
		value = err.Error()
	}

	log.Printf("[%s][%s] thread writing results", crypto, currency)
	resultsWriterMutex.Lock()
	results[crypto][currency] = value
	defer resultsWriterMutex.Unlock()
	t := time.Now().Sub(startTime).Seconds()
	log.Printf("[%s][%s] thread done after %v seconds", crypto, currency, t)
}
