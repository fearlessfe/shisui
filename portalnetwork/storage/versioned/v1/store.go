package v1

import (
	"database/sql"
	"errors"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/portalnetwork/storage"
	"github.com/holiman/uint256"
)

var _ storage.ContentStorage = &IdIndexV1Store{}

type IdIndexV1Store struct {
	config IdIndexedV1StoreConfig
	radius atomic.Value
}

func NewIdIndexV1Store(config IdIndexedV1StoreConfig) (*IdIndexV1Store, error) {
	store := &IdIndexV1Store{
		config: config,
	}
	store.radius.Store(storage.MaxDistance)
	_, err := config.sqliteDB.Exec(CreateTableSql(config.storeType))
	if err != nil {
		return nil, err
	}
	return store, nil
}

// Get implements storage.ContentStorage.
func (i *IdIndexV1Store) Get(contentKey []byte, contentId []byte) ([]byte, error) {
	var res []byte
	err := i.config.sqliteDB.QueryRow(LookupValueSql(i.config.storeType), contentId).Scan(&res)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, storage.ErrContentNotFound
	}
	return res, err
}

// Put implements storage.ContentStorage.
func (i *IdIndexV1Store) Put(contentKey []byte, contentId []byte, content []byte) error {
	panic("unimplemented")
}

// Radius implements storage.ContentStorage.
func (i *IdIndexV1Store) Radius() *uint256.Int {
	radius := i.radius.Load()
	val := radius.(*uint256.Int)
	return val
}
