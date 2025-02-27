// Description: This package provides a client to interact with the Bitso API.
// The Bitso API documentation can be found at https://bitso.com/api_info.
// The Bitso API provides the following endpoints:
// - Ticker: Retrieve the ticker for the given cryptocurrency.
// - Order Book: Retrieve the order book for the given cryptocurrency.
// - Trades: Retrieve the trades for the given cryptocurrency.
// - Available Books: Retrieve the available books.
// The Bitso API returns the data in JSON format.
//
// author: Uriel Marquez (uriel.marquez@wizeline.com)
// version: 0.1.0
package bitso_client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const bitsoProdBaseUrl = "https://api-stage.bitso.com/api/v3"
const bitsoSandboxBaseUrl = "https://api-sandbox.bitso.com/api/v3"

// BitsoError represents the error data.
type BitsoError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type bitsoBaseResponse struct {
	Success bool       `json:"success"`
	Error   BitsoError `json:"error,omitempty"`
}

type bitsoClient struct {
	baseUrl string
}

// TickerName represents the name of a ticker.
type TickerName string

const (
	BTC_MXN TickerName = "btc_mxn" // BTC_MXN for Bitcoin to Mexican Pesos.
	ETH_MXN TickerName = "eth_mxn" // ETH_MXN for Ethereum to Mexican Pesos.
	XRP_MXN TickerName = "xrp_mxn" // XRP_MXN for Ripple to Mexican Pesos.
	BTC_USD TickerName = "btc_usd" // BTC_USD for Bitcoin to US Dollars.
	ETH_USD TickerName = "eth_usd" // ETH_USD for Ethereum to US Dollars.
	XRP_USD TickerName = "xrp_usd" // XRP_USD for Ripple to US Dollars.
)

type bitsoPayload struct {
	High                 string            `json:"high"`
	Last                 string            `json:"last"`
	CreatedAt            time.Time         `json:"created_at"`
	Book                 string            `json:"book"`
	Volume               string            `json:"volume"`
	Vwap                 string            `json:"vwap"`
	Low                  string            `json:"low"`
	Ask                  string            `json:"ask"`
	Bid                  string            `json:"bid"`
	Change24             string            `json:"change_24"`
	RollingAverageChange map[string]string `json:"rolling_average_change"`
}

// Ticker represents the ticker data.
type Ticker struct {
	bitsoBaseResponse
	Payload bitsoPayload `json:"payload"`
}

// OrderBook represents the order book data.
// TODO: Implement the OrderBook struct.
type OrderBook struct{}

// Trade represents the trade data.
// TODO: Implement the Trade struct.
type Trade struct{}

// Book represents the book data.
// TODO: Implement the Book struct.
type Book struct{}

// Client represents the Bitso client.
type Client interface {
	GetTicker(ticker TickerName) (Ticker, error)
	GetOrderBook(ticker TickerName) (OrderBook, error)
	GetTrades(ticker TickerName) ([]Trade, error)
	GetAvailableBooks() ([]Book, error)
}

// NewClient creates a new instance of the Bitso client.
// If productionMode is true, it creates a client for the production environment.
// Otherwise, it creates a client for the sandbox environment.
// Please use the sandbox environment for dev and testing purposes, and the
// production environment for real-world scenarios only.
// Example:
//
//	client := NewClient(productionMode=false)
//	ticker, err := client.GetTicker("btc_mxn")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Last: %s\n", ticker.Last)
//	fmt.Printf("Volume: %s\n", ticker.Volume)
//	fmt.Printf("CreatedAt: %s\n", ticker.CreatedAt)
//	fmt.Printf("Error: %v\n", ticker.Error)
//	fmt.Printf("Success: %v\n", ticker.Success)
//	fmt.Printf("High: %s\n", ticker.High)
//	fmt.Printf("Book: %s\n", ticker.Book)
//	fmt.Printf("Vwap: %s\n", ticker.Vwap)
//	fmt.Printf("Low: %s\n", ticker.Low)
//	fmt.Printf("Ask: %s\n", ticker.Ask)
//	fmt.Printf("Bid: %s\n", ticker.Bid)
//	fmt.Printf("Change24: %s\n", ticker.Change24)
//	fmt.Printf("RollingAverageChange: %v\n", ticker.RollingAverageChange)
func NewClient(productionMode bool) Client {
	if productionMode {
		return &bitsoClient{baseUrl: bitsoProdBaseUrl}
	} else {
		return &bitsoClient{baseUrl: bitsoSandboxBaseUrl}
	}
}

// getTicker calls `<bitsoBaseUrl>/ticker?book=<name>` to retrieve the ticker
// data and returns the result if "success" == true.
// Otherwise, it returns an error.
// ref: https://docs.bitso.com/bitso-api/docs/ticker
// Example response:
//
//	{
//		"success": true,
//		"payload": {
//			"high": "12345.67",
//			"last": "12345.67",
//			"created_at": "2021-01-01T00:00:00+00:00",
//			"book": "btc_mxn",
//			"volume": "12345.67",
//			"vwap": "12345.67",
//			"low": "12345.67",
//			"ask": "12345.67",
//			"bid": "12345.67",
//			"change_24": "12345.67",
//			"rolling_average_change": {}
//		}
//	}
func (c *bitsoClient) getTicker(name TickerName) (t Ticker, err error) {
	url := fmt.Sprintf("%s/ticker?book=%s", c.baseUrl, name)
	resp, err := http.Get(url)
	if err != nil {
		return t, fmt.Errorf("failed to get ticker: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return t, fmt.Errorf("failed to get ticker: %s", resp.Status)
	}

	defer resp.Body.Close()

	/*// read the response body and decode it into the Ticker struct
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return t, fmt.Errorf("failed to read response: %w", err)
	}
	_ = content*/

	if err = json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return t, fmt.Errorf("failed to decode response: %w", err)
	}

	if !t.Success {
		return t, fmt.Errorf("failed to get ticker, API response: (%v)%s", t.Error.Code, t.Error.Message)
	}

	return t, nil
}

// getOrderBook calls `<bitsoBaseUrl>/order_book?book=<name>` to retrieve the order book
// data and returns the result if "success" == true.
// Otherwise, it returns an error.
// TODO: Implement the getOrderBook method.
func (c *bitsoClient) getOrderBook(name TickerName) (ob OrderBook, err error) {
	return
}

// getAvailableBooks calls `<bitsoBaseUrl>/available_books` to retrieve the available books
// data and returns the result if "success" == true.
// Otherwise, it returns an error.
// TODO: Implement the getAvailableBooks method.
func (c *bitsoClient) getAvailableBooks() (b []Book, err error) {
	return
}

// getTrades calls `<bitsoBaseUrl>/trades?book=<name>` to retrieve the trades
// data and returns the result if "success" == true.
// Otherwise, it returns an error.
// TODO: Implement the getTrades method.
func (c *bitsoClient) getTrades(name TickerName) (t []Trade, err error) {
	return
}

// GetTicker retrieves the ticker for the given cryptocurrency.
// It returns an error if the request fails or if the API response is not successful.
// Otherwise, it returns the ticker data.
// The TickerName is the name of the cryptocurrency (ex. BTC_MXN).
// The Ticker struct contains the following fields:
// - High
// - Last
// - CreatedAt
// - Book
// - Volume
// - Vwap
// - Low
// - Ask
// - Bid
// - Change24
// - RollingAverageChange
// - Success
// - Error (if the request fails)
// The CreatedAt field is a time.Time value.
// The RollingAverageChange field is a map[string]string.
// The Success field is a boolean.
// The Error field is a BitsoError.
// The BitsoError struct contains the following fields:
// - Code
// - Message
// Example:
//
//	ticker, err := client.GetTicker("btc_mxn")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Last: %s\n", ticker.Last)
//	fmt.Printf("Volume: %s\n", ticker.Volume)
//	fmt.Printf("CreatedAt: %s\n", ticker.CreatedAt)
//	fmt.Printf("Error: %v\n", ticker.Error)
//	fmt.Printf("Success: %v\n", ticker.Success)
//	fmt.Printf("High: %s\n", ticker.High)
//	fmt.Printf("Book: %s\n", ticker.Book)
//	fmt.Printf("Vwap: %s\n", ticker.Vwap)
//	fmt.Printf("Low: %s\n", ticker.Low)
//	fmt.Printf("Ask: %s\n", ticker.Ask)
//	fmt.Printf("Bid: %s\n", ticker.Bid)
//	fmt.Printf("Change24: %s\n", ticker.Change24)
//	fmt.Printf("RollingAverageChange: %v\n", ticker.RollingAverageChange)
//
// The output should be:
//
//	Last: 12345.67
//	Volume: 12345.67
//	CreatedAt: 2021-01-01 00:00:00 +0000 UTC
//	Error: {Code: Message:}
//	Success: true
//	High: 12345.67
//	Book: btc_mxn
//	Vwap: 12345.67
//	Low: 12345.67
//	Ask: 12345.67
//	Bid: 12345.67
//	Change24: 12345.67
//	RollingAverageChange: map[]
//
// The CreatedAt field should be a time.Time value.
// The Error field should be a BitsoError.
// The RollingAverageChange field should be a map[string]string.
// The Success field should be a boolean.
func (c *bitsoClient) GetTicker(ticker TickerName) (Ticker, error) {
	return c.getTicker(ticker)
}

// GetOrderBook retrieves the order book for the given cryptocurrency.
func (c *bitsoClient) GetOrderBook(ticker TickerName) (OrderBook, error) {
	return c.getOrderBook(ticker)
}

// GetTrades retrieves the trades for the given cryptocurrency.
func (c *bitsoClient) GetTrades(ticker TickerName) ([]Trade, error) {
	return c.getTrades(ticker)
}

// GetAvailableBooks retrieves the available books.
func (c *bitsoClient) GetAvailableBooks() ([]Book, error) {
	return c.getAvailableBooks()
}
