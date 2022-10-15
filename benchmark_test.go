package kvnuts

import (
	"testing"
)

func Benchmark_InMemory(b *testing.B) {
	b.ResetTimer()

	key := []byte("x")
	value := []byte("y")
	bucket := "A"

	store, _ := NewStoreInMemory(_segmentSizeTests)

	for i := 0; i < b.N; i++ {
		store.Set(bucket, key, value)
		store.GetVByK(bucket, key)
		store.DeleteKVByK(bucket, key)
	}
}

func Benchmark_KeyInMemory(b *testing.B) {
	b.ResetTimer()

	key := []byte("x")
	value := []byte("y")
	bucket := "A"

	store, _ := NewStoreInMemory(_segmentSizeTests)

	for i := 0; i < b.N; i++ {
		store.Set(bucket, key, value)
		store.GetVByK(bucket, key)
		store.DeleteKVByK(bucket, key)
	}
}
