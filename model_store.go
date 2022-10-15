package kvnuts

import (
	"fmt"

	"github.com/xujiajun/nutsdb"
)

// KV is key value for the NoSQL DB.
type KV struct {
	Key   []byte
	Value []byte
}

// BStore Concentrates information defining a KV store.
type KVStore struct {
	Store *nutsdb.DB
}

const _folderDB = "./nutsdb"
const _segmentSizeTests = 4 * 1024 * 1024 // 4 MB

// NewStoreInMemory returns a type containing a store that satisfies store interface.
// With test segment size.
func NewStoreInMemory(mbSegmentSize uint) (*KVStore, error) {
	db, errOpen := nutsdb.Open(nutsdb.DefaultOptions, nutsdb.WithDir(_folderDB), nutsdb.WithSegmentSize(int64(mbSegmentSize)))
	if errOpen != nil {
		return nil, fmt.Errorf("could not create database in folder: %s, %w", _folderDB, errOpen)
	}

	return &KVStore{
		Store: db,
	}, nil
}

// NewStore returns a type containing a store that satisfies store interface.
// With test segment size.
func NewStore(mbSegmentSize uint) (*KVStore, error) {
	db, errOpen := nutsdb.Open(nutsdb.DefaultOptions, nutsdb.WithDir(_folderDB), nutsdb.WithEntryIdxMode(nutsdb.HintKeyAndRAMIdxMode), nutsdb.WithSegmentSize(int64(mbSegmentSize)))
	if errOpen != nil {
		return nil, fmt.Errorf("could not create database in folder: %s, %w", _folderDB, errOpen)
	}

	return &KVStore{
		Store: db,
	}, nil
}

// Close closes the store.
func (s *KVStore) Close() error {
	return s.Store.Close()
}
