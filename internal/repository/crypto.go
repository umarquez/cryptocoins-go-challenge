package repository

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/tidwall/buntdb"
)

type Crypto interface {
	GetValue(id string) (string, error)
	StoreValue(id string, value string) error
}

type crypto struct {
	dbcnn   *buntdb.DB
	mutex   *sync.Mutex
	dataTTL time.Duration
}

func NewCryptoRepository(dbcnn *buntdb.DB, mutex *sync.Mutex, dataTTL time.Duration) Crypto {
	instance := new(crypto)
	instance.dbcnn = dbcnn
	instance.mutex = mutex
	instance.dataTTL = dataTTL

	return instance
}

func (repo *crypto) GetValue(id string) (string, error) {
	result := ""
	err := repo.dbcnn.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(id, true)
		if err != nil && errors.Is(err, buntdb.ErrNotFound) {
			return nil
		} else if err != nil {
			return fmt.Errorf("db.Get(%v) error: %v", id, err)
		}

		result = val
		return nil
	})
	if err != nil {
		return "", err
	}

	return result, nil
}

func (repo *crypto) StoreValue(id string, val string) error {
	repo.mutex.Lock()
	err := repo.dbcnn.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(id, val, &buntdb.SetOptions{
			Expires: true,
			TTL:     repo.dataTTL,
		})
		if err != nil {
			return fmt.Errorf("db.Set: %v", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("update(%v): %v", id, err)
	}
	repo.mutex.Unlock()

	return nil
}
