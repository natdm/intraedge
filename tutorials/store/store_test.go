package store_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/natdm/intraedge/tutorials/store"
)

func TestStore(t *testing.T) {
	storage := store.New()
	defer storage.Close()
	err := storage.Add("Hello", "World")
	require.NoError(t, err)
	v, err := storage.Val("Hello")
	require.NoError(t, err)
	require.Equal(t, "World", v)
}
