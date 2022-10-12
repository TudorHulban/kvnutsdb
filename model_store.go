package kvnuts

import (
	"fmt"
	"log"

	"github.com/xujiajun/nutsdb"
)

// KV is key value for the NoSQL DB.
type KV struct {
	Key   []byte
	Value []byte
}

// BStore Concentrates information defining a KV store.
type KVStore struct {
	logger *log.Logger // logger needed only for package logging
	Store  *nutsdb.DB
}

const _folderDB = "./nutsdb"
const _segmentSizeTesting = 4 * 1024 * 1024 // 4 MB

// NewStore returns a type containing a store that satisfies store interface.
// With test segment size.
func NewStore(l *log.Logger) (*KVStore, error) {
	db, errOpen := nutsdb.Open(nutsdb.DefaultOptions, nutsdb.WithDir(_folderDB), nutsdb.WithSegmentSize(_segmentSizeTesting))
	if errOpen != nil {
		return nil, fmt.Errorf("could not create database in folder: %s, %w", _folderDB, errOpen)
	}

	return &KVStore{
		logger: l,
		Store:  db,
	}, nil
}

// Close closes the store.
func (s *KVStore) Close() error {
	return s.Store.Close()
}
