package kvnuts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemorySetUpdateDelete(t *testing.T) {
	store, err := NewStoreInMemory(_segmentSizeTests)
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, store.Close())
	}()

	key := []byte("x")
	value := []byte("y")
	bucket := "A"

	assert.NoError(t, store.Set(bucket, key, value))

	fetchedValue1, errGet1 := store.GetVByK(bucket, key)
	assert.NoError(t, errGet1)
	assert.Equal(t, value, fetchedValue1)

	updateValue := []byte("z")
	assert.NoError(t, store.Set(bucket, key, updateValue))

	fetchedValue2, errGet2 := store.GetVByK(bucket, key)

	t.Logf("updated value: %s", updateValue)
	t.Logf("fetched: %s", fetchedValue2)

	assert.NoError(t, errGet2)
	assert.Equal(t, updateValue, fetchedValue2)

	require.NoError(t, store.DeleteKVByK(bucket, key))

	fetchedValue3, errGet3 := store.GetVByK(bucket, key)
	assert.Error(t, errGet3)
	assert.Nil(t, fetchedValue3)
}

func TestDiskSetUpdateDelete(t *testing.T) {
	store, err := NewStore(_segmentSizeTests)
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, store.Close())
	}()

	key := []byte("x")
	value := []byte("y")
	bucket := "A"

	assert.NoError(t, store.Set(bucket, key, value))

	fetchedValue1, errGet1 := store.GetVByK(bucket, key)
	assert.NoError(t, errGet1)
	assert.Equal(t, value, fetchedValue1)

	updateValue := []byte("z")
	assert.NoError(t, store.Set(bucket, key, updateValue))

	fetchedValue2, errGet2 := store.GetVByK(bucket, key)

	t.Logf("updated value: %s", updateValue)
	t.Logf("fetched: %s", fetchedValue2)

	assert.NoError(t, errGet2)
	assert.Equal(t, updateValue, fetchedValue2)

	require.NoError(t, store.DeleteKVByK(bucket, key))

	fetchedValue3, errGet3 := store.GetVByK(bucket, key)
	assert.Error(t, errGet3)
	assert.Nil(t, fetchedValue3)
}
