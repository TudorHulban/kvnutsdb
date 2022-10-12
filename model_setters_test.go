package kvnuts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	store, err := NewStore(nil)
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, store.Close())
	}()

	key := []byte("x")
	value := []byte("y")
	bucket := "A"

	assert.NoError(t, store.Set(bucket, key, value))

	fetchedValue, errGet := store.GetVByK(bucket, key)
	assert.NoError(t, errGet)
	assert.Equal(t, value, fetchedValue)

	// updateValue := []byte("z")
	// assert.NoError(t, inMemoryStore.Set(key, updateValue))

	// t.Logf("value: %s", updateValue)
	// t.Logf("fetched: %s", fetchedValue)

}
