package gokvstores

import (
	"sort"
	"testing"

	conv "github.com/cstockton/go-conv"
	"github.com/stretchr/testify/assert"
)

func testStore(t *testing.T, store KVStore) {
	is := assert.New(t)

	err := store.Flush()
	is.NoError(err)

	// Set

	err = store.Set("key", "value")
	is.NoError(err)

	// Get

	v, err := store.Get("key")
	val, err := conv.String(v)
	is.NoError(err)
	is.Equal("value", val)

	// Exists

	exists, err := store.Exists("key")
	is.NoError(err)
	is.True(exists)

	// Delete

	err = store.Delete("key")
	is.NoError(err)

	v, _ = store.Get("key")
	is.Nil(v)

	exists, err = store.Exists("key")
	is.NoError(err)
	is.False(exists)

	// Map

	mapResults := map[string]map[string]interface{}{
		"key1": {"language": "go"},
		"key2": {"integer": "1"},
		"key3": {"float": "20.2"},
	}

	for key, expected := range mapResults {
		err = store.SetMap(key, expected)
		is.NoError(err)

		v, err := store.GetMap(key)
		is.Equal(expected, v)

		exists, err := store.Exists(key)
		is.NoError(err)
		is.True(exists)
	}

	keys := []string{"key1", "key2", "key3"}

	results, err := store.GetMaps(keys)
	is.NoError(err)

	for key, result := range results {
		is.Equal(result, mapResults[key])
	}

	for key := range mapResults {
		err = store.Delete(key)
		is.NoError(err)

		v, _ = store.GetMap(key)
		is.Nil(v)

		exists, err = store.Exists(key)
		is.NoError(err)
		is.False(exists)
	}

	// Slices

	sliceResults := map[string][]interface{}{
		"key1": {"one", "two", "three", "four"},
		"key2": {"1", "2", "3", "4"},
		"key3": {"1.0", "1.1", "1.2", "1.3"},
	}

	for key, expected := range sliceResults {
		err = store.SetSlice(key, expected)
		is.NoError(err)

		expectedStrings, err := stringSlice(expected)
		is.NoError(err)

		v, err := store.GetSlice(key)
		is.NoError(err)
		strings, err := stringSlice(v)
		is.NoError(err)
		is.Equal(expectedStrings, strings)

		exists, err := store.Exists(key)
		is.NoError(err)
		is.True(exists)

		err = store.AppendSlice(key, "append1", "append2")
		is.NoError(err)

		v, err = store.GetSlice(key)
		is.NoError(err)

		expectedStrings = append(expectedStrings, []string{"append1", "append2"}...)
		sort.Strings(expectedStrings)
		values, err := stringSlice(v)
		is.NoError(err)
		is.Equal(expectedStrings, values)

		err = store.Delete(key)
		is.NoError(err)

		v, _ = store.GetSlice(key)
		is.Nil(v)

		exists, err = store.Exists(key)
		is.NoError(err)
		is.False(exists)

	}

}
