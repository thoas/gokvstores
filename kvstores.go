package gokvstores

import (
	"sort"

	conv "github.com/cstockton/go-conv"
)

// KVStore is the KV store interface.
type KVStore interface {
	// Get returns value for the given key.
	Get(key string) (interface{}, error)

	// Set sets value for the given key.
	Set(key string, value interface{}) error

	// GetMap returns map for the given key.
	GetMap(key string) (map[string]interface{}, error)

	// GetMaps returns maps for the given keys.
	GetMaps(keys []string) (map[string]map[string]interface{}, error)

	// SetMap sets map for the given key.
	SetMap(key string, value map[string]interface{}) error

	// GetSlice returns slice for the given key.
	GetSlice(key string) ([]interface{}, error)

	// SetSlice sets slice for the given key.
	SetSlice(key string, value []interface{}) error

	// AppendSlice appends values to an existing slice.
	// If key does not exist, creates slice.
	AppendSlice(key string, values ...interface{}) error

	// Exists checks if the given key exists.
	Exists(key string) (bool, error)

	// Delete deletes the given key.
	Delete(key string) error

	// Flush flushes the store.
	Flush() error

	// Return all keys matching pattern
	Keys(pattern string) ([]interface{}, error)

	// Close closes the connection to the store.
	Close() error
}

func stringSlice(values []interface{}) ([]string, error) {
	converted := []string{}

	for _, v := range values {
		if v != nil {
			val, err := conv.String(v)
			if err != nil {
				return nil, err
			}

			converted = append(converted, val)
		}
	}

	sort.Strings(converted)

	return converted, nil
}
