package hasher_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/marcelofabianov/identity-gateway/pkg/hasher"
)

func TestHash(t *testing.T) {
	h := hasher.NewHasher()

	data := "mysecretpassword"
	hashed, err := h.Hash(data)
	require.NoError(t, err)

	parts := strings.Split(hashed, "$")
	require.Len(t, parts, 3)
}

func TestCompare(t *testing.T) {
	h := hasher.NewHasher()

	data := "mysecretpassword"
	hashed, err := h.Hash(data)
	require.NoError(t, err)

	match, err := h.Compare(data, hashed)
	require.NoError(t, err)
	assert.True(t, match)

	incorrectData := "anotherpassword"
	match, err = h.Compare(incorrectData, hashed)
	require.NoError(t, err)
	assert.False(t, match)
}
