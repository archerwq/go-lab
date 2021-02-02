package main

import (
	"fmt"
)

// Store .
type Store interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
	Incr(key string) error
}

// MockStore .
// If you embed interfaces on a struct, you are actually adding new fields to
// the struct with the name of the interface so if you don't init those,
// you will get panics because they are nil.
// Embedding promotes all of the methods of the embedded type (struct or interface, doesn't matter)
// to be methods of the parent type, but called using the embedded object as the receiver.
type MockStore struct {
	Store
}

// NewMockStore .
func NewMockStore() *MockStore {
	return &MockStore{}
}

// Del .
func (m *MockStore) Del(key string) error {
	fmt.Println("delete key", key)
	return nil
}

func main() {
	store := NewMockStore()
	store.Del("A")

	// Trying to call any of those methods without defining something to fill
	// that interface field in the struct, however, will panic, as that
	// interface field defaults to nil.
	store.Set("A", 1)
	store.Get("A")
	store.Incr("A")
}
