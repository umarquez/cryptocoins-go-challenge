package usecase_test

import (
	"testing"

	"github.com/umarquez/cryptocoins-go-challenge/internal/domain"
	"github.com/umarquez/cryptocoins-go-challenge/internal/usecase"
)

func Test_cryptoUseCase_GetCryptos(t *testing.T) {
	tests := []struct {
		name    string
		want    []domain.Crypto
		wantErr bool
	}{
		{
			name:    "GetCryptos()",
			want:    make([]domain.Crypto, len(domain.Cryptos)),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := usecase.NewCryptoUseCase()
			got, err := uc.GetCryptos()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCryptos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("GetCryptos() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
