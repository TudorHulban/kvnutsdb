package kvnuts

import "github.com/xujiajun/nutsdb"

// GetVByK fetches value from store based on passed key.
func (s *KVStore) GetVByK(bucket string, key []byte) ([]byte, error) {
	var res []byte

	errView := s.Store.View(func(txn *nutsdb.Tx) error {
		value, errGet := txn.Get(bucket, []byte(key))
		if errGet != nil {
			return errGet
		}

		res = value.Value

		return nil
	})

	return res, errView
}
