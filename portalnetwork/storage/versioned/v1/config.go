package v1

import (
	"database/sql"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/portalnetwork/storage"
)

type IdIndexedV1StoreConfig struct {
	storeType              StoreType
	nodeId                 enode.ID
	storageCapacityInBytes uint64
	sqliteDB               *sql.DB
	log                    log.Logger
}

func NewIdIndexedV1StoreConfig(storeType StoreType, config storage.PortalStorageConfig) IdIndexedV1StoreConfig {
	return IdIndexedV1StoreConfig{
		storeType:              storeType,
		nodeId:                 config.NodeId,
		storageCapacityInBytes: config.StorageCapacityMB,
		sqliteDB:               config.DB,
		log:                    log.New("storage", storeType),
	}
}
