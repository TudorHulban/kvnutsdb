package kvnuts

import (
	"github.com/xujiajun/nutsdb"
)

// DeleteKVByK Deletes KV by key.
func (s *KVStore) DeleteKVByK(bucket string, key []byte) error {
	return s.Store.Update(func(txn *nutsdb.Tx) error {
		return txn.Delete(bucket, []byte(key))
	})
}
