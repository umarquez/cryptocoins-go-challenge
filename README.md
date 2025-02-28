# Cryptocoins Go challenge

  [![Test and Build](https://github.com/umarquez/cryptocoins-go-challenge/actions/workflows/test_and_build.yaml/badge.svg)](https://github.com/umarquez/cryptocoins-go-challenge/actions/workflows/test_and_build.yaml)
  [![Test Endpoint](https://github.com/umarquez/cryptocoins-go-challenge/actions/workflows/test_endpoint.yaml/badge.svg)](https://github.com/umarquez/cryptocoins-go-challenge/actions/workflows/test_endpoint.yaml)

## Requirements:

#### Layout

```json
[
    {
        "id": 1,
        "component": "crypto_btc",
        "model": {}
    },
    {
        "id": 2,
        "component": "crypto_eth",
        "model": {}
    },
    {
        "id": 3,
        "component": "crypto_xrp",
        "model": {}
    }
]
```

#### Model

```json
{
    "date": "2025-02-26T17:00:00",
    "name": "Bitcoin",
    "ticker_symbol": "BTC",
    "price": {
        "usd": 123.456,
        "mxn": 123.456
    }
}
```

### ACs

#### API
- design the endpoint
- simulate provider to obtain the JSON with the configuration of the components
- implement provider to obtain the values of the cryptocurrencies
    - coinmarketcap
    - coinbase
    - coingecko
    - whichever is available and free
- obtain the values of the cryptocurrencies concurrently and deposit the values in their respective model
- use best practices and known conventions

## Installation

To install the dependencies, run:
```sh
make generate
```

## Usage

To start the application, run:
```sh
make run
```
