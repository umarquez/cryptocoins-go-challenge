package domain

import (
	"time"
)

/*var CryptoByIdTable map[int]string
var IdByCryptoTable map[string]int

func init() {
	for i, crypto := range Cryptos {
		if CryptoByIdTable == nil {
			CryptoByIdTable = make(map[int]string)
		}

		if IdByCryptoTable == nil {
			IdByCryptoTable = make(map[string]int)
		}
		CryptoByIdTable[i] = string(crypto)
		IdByCryptoTable[string(crypto)] = i
	}

	println("IdByCryptoTable: ", IdByCryptoTable)
}*/

type Currency string

const MXN Currency = "MXN"
const USD Currency = "USD"

type CryptoCurrency string

const BTC CryptoCurrency = "BTC"
const ETH CryptoCurrency = "ETH"
const XRP CryptoCurrency = "XRP"

var CryptoByIdEnum = map[int]CryptoCurrency{
	0: BTC,
	1: ETH,
	2: XRP,
}

var CryptoIdEnum = map[CryptoCurrency]int{
	BTC: 0,
	ETH: 1,
	XRP: 2,
}

var Currencies = []Currency{MXN, USD}
var Cryptos = []CryptoCurrency{BTC, ETH, XRP}

// CryptoName represents the name of a cryptocurrency.
var CryptoName = map[CryptoCurrency]string{
	BTC: "Bitcoin",
	ETH: "Ethereum",
	XRP: "Ripple",
}

// Price represents the pricing details for a cryptocurrency.
type Price struct {
	USD string `json:"usd"`
	MXN string `json:"mxn"`
}

// Crypto represents the core domain entity for a cryptocurrency.
type Crypto struct {
	Date         time.Time `json:"date"`          // The time at which the price is fetched.
	Name         string    `json:"name"`          // The name of the cryptocurrency (e.g., Bitcoin).
	TickerSymbol string    `json:"ticker_symbol"` // The ticker symbol (e.g., BTC).
	Price        Price     `json:"price"`         // The price in different currencies.
}
