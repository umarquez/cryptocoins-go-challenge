package bitso_client

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetTicker(t *testing.T) {
	tests := []struct {
		name         string
		tickerName   TickerName
		serverStatus int
		response     interface{}
		wantErr      bool
		errMsg       string
	}{
		{
			name:         "successful response",
			tickerName:   BTC_MXN,
			serverStatus: http.StatusOK,
			response: Ticker{
				bitsoBaseResponse: bitsoBaseResponse{Success: true},
				Payload: bitsoPayload{
					Book: string(BTC_MXN),
				},
			},
			wantErr: false,
		},
		{
			name:         "bad request",
			tickerName:   TickerName("btc_btc"),
			serverStatus: http.StatusBadRequest,
			response:     nil,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(false)

			got, err := client.GetTicker(tt.tickerName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTicker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if !got.Success {
					fmt.Printf("error: %s\n", err)
				}
				return
			}

			expected := tt.response.(Ticker)
			if got.Payload.Book != expected.Payload.Book ||
				got.Success != expected.Success {
				t.Errorf("getTicker() got = %v, want %v", got, expected)
			}
		})
	}
}
