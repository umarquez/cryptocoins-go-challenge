package crypto

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/umarquez/cryptocoins-go-challenge/sdk/bitso_client"
)

func GetCryptoValue(crypto, currency string) (string, error) {
	// Simulate a delay to simulate the time it takes to fetch the data.
	delay := time.Duration(rand.Intn(3000)) * time.Millisecond
	time.Sleep(delay)

	c := bitso_client.NewClient(false)

	// Retry fetching the data up to 3 times in case of an error.
	for retriesCount := 0; retriesCount < 3; retriesCount++ {
		ticker, err := c.GetTicker(bitso_client.TickerName(fmt.Sprintf("%s_%s", crypto, currency)))
		if err != nil {
			fmt.Printf("(%v) Error fetching crypto value: %v\n", retriesCount, err)
			continue
		}

		return ticker.Payload.Last, nil
	}

	return "", fmt.Errorf("failed to fetch crypto value after 3 retries")
}
