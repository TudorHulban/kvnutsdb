package kvnuts

import "github.com/xujiajun/nutsdb"

// Set sets or updates a key in store.
func (s *KVStore) Set(bucket string, key, value []byte) error {
	return s.Store.Update(func(txn *nutsdb.Tx) error {
		return txn.Put(bucket, key, value, 0)
	})
}
