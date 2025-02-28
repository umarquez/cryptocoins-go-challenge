package usecase_test

import (
	"sync"
	"testing"
	"time"

	"github.com/umarquez/cryptocoins-go-challenge/internal/domain"
	"github.com/umarquez/cryptocoins-go-challenge/internal/repository"
	"github.com/umarquez/cryptocoins-go-challenge/internal/service"
	"github.com/umarquez/cryptocoins-go-challenge/internal/usecase"

	"github.com/tidwall/buntdb"
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
			db, err := buntdb.Open(":memory:")
			if err != nil {
				t.Errorf("GetCryptos() error = %v", err)
				return
			}
			defer db.Close()

			srv := service.GetCryptoService()
			repo := repository.NewCryptoRepository(db, new(sync.Mutex), time.Minute)

			uc := usecase.NewCryptoUseCase(srv, repo)
			got, err := uc.GetAllCryptos()
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
