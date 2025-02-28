package usecase

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/umarquez/cryptocoins-go-challenge/internal/domain"
	"github.com/umarquez/cryptocoins-go-challenge/internal/repository"
	"github.com/umarquez/cryptocoins-go-challenge/internal/service"
)

// CryptoService defines the contract for crypto_service business logic.
type CryptoService interface {
	GetValue(crypto domain.CryptoCurrency, currency domain.Currency) (string, error)
}

// CryptoRepo defines the contract for crypto_repo business logic.
type CryptoRepo interface {
	GetValue(key string) (string, error)
	StoreValue(key, value string) error
}

// CryptoUseCase defines the contract for crypto_service business logic.
type CryptoUseCase interface {
	GetAllCryptos() ([]domain.Crypto, error)
	GetCryptoById(id int) (domain.Crypto, error)
}

type cryptoUseCase struct {
	cryptoService service.Crypto
	cryptoRepo    repository.Crypto
}

func NewCryptoUseCase(srv CryptoService, repo CryptoRepo) CryptoUseCase {
	return &cryptoUseCase{
		cryptoService: srv,
		cryptoRepo:    repo,
	}
}

// getCryptoValueAsync retrieves the last price of a cryptocurrency in a given currency.
func (uc *cryptoUseCase) getCryptoValueAsync(crypto domain.CryptoCurrency, currency domain.Currency, wg *sync.WaitGroup, results map[domain.CryptoCurrency]map[domain.Currency]string, resultsWriterMutex *sync.Mutex) {
	log.Printf("[%s][%s] thread started\n", crypto, currency)
	startTime := time.Now()
	defer wg.Done()

	value, err := uc.cryptoRepo.GetValue(fmt.Sprintf("%s_%s", crypto, currency))
	if err != nil {
		log.Printf("[%s][%s] cryptoRepo.GetValue: %v", crypto, currency, err)
		return
	}

	if value == "" {
		log.Printf("[%s][%s] cryptoRepo.GetValue returned an empty value", crypto, currency)
		log.Printf("[%s][%s] fetching value from cryptoService", crypto, currency)
		value, err = uc.cryptoService.GetValue(crypto, currency)
		if err != nil {
			log.Printf("[%s][%s] cryptoService.GetValue: %v", crypto, currency, err)
			return
		}

		log.Printf("[%s][%s] storing value in cryptoRepo", crypto, currency)
		err = uc.cryptoRepo.StoreValue(fmt.Sprintf("%s_%s", crypto, currency), value)
		if err != nil {
			log.Printf("[%s][%s] cryptoRepo.StoreValue: %v", crypto, currency, err)
			return
		}
	} else {
		log.Printf("[%s][%s] value found in cache", crypto, currency)
	}

	log.Printf("[%s][%s] thread writing results", crypto, currency)
	resultsWriterMutex.Lock()
	results[crypto][currency] = value
	defer resultsWriterMutex.Unlock()
	t := time.Now().Sub(startTime).Seconds()
	log.Printf("[%s][%s] thread done after %v seconds", crypto, currency, t)
}

func (uc *cryptoUseCase) GetAllCryptos() ([]domain.Crypto, error) {
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

// GetCryptoById returns a single instance of the crypto service.
func (uc *cryptoUseCase) GetCryptoById(id int) (domain.Crypto, error) {
	// Create a map to store the results.
	results := make(map[domain.CryptoCurrency]map[domain.Currency]string)
	resultsWriterMutex := new(sync.Mutex)

	// Create a wait group to wait for all goroutines to finish.
	wg := new(sync.WaitGroup)
	crypto, ok := domain.CryptoByIdEnum[id]
	if !ok {
		return domain.Crypto{}, fmt.Errorf("crypto not found")
	}

	for _, currency := range domain.Currencies {
		if _, ok := results[crypto]; !ok {
			results[crypto] = make(map[domain.Currency]string)
		}

		wg.Add(1)
		go uc.getCryptoValueAsync(crypto, currency, wg, results, resultsWriterMutex)
	}

	wg.Wait()
	fmt.Printf("Results: %v\n", results)

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
		return c, nil
	}

	return domain.Crypto{}, fmt.Errorf("crypto not found") // should never reach this point
}
