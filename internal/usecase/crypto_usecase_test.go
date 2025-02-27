package usecase

import (
	"reflect"
	"testing"

	"github.com/umarquez/cryptocoins-go-challenge/internal/domain"
)

func TestNewCryptoUseCase(t *testing.T) {
	tests := []struct {
		name string
		want CryptoUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCryptoUseCase(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCryptoUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cryptoUseCase_GetCryptos(t *testing.T) {
	tests := []struct {
		name    string
		want    []domain.Crypto
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &cryptoUseCase{}
			got, err := uc.GetCryptos()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCryptos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCryptos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cryptoUseCase_getCryptoValueAsync(t *testing.T) {
	type args struct {
		crypto             domain.CryptoCurrency
		currency           domain.Currency
		wg                 *sync.WaitGroup
		results            map[domain.CryptoCurrency]map[domain.Currency]string
		resultsWriterMutex *sync.Mutex
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &cryptoUseCase{}
			uc.getCryptoValueAsync(tt.args.crypto, tt.args.currency, tt.args.wg, tt.args.results, tt.args.resultsWriterMutex)
		})
	}
}
