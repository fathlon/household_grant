package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDatastore(t *testing.T) {
	// When:
	result := NewDatastore()

	// Then:
	require.NotNil(t, result.Households)
	require.NotNil(t, result.Members)
}

func TestNextHseIndex(t *testing.T) {
	for i := 1; i < 3; i++ {
		require.Equal(t, i, nextHseIndex())
	}
}

func TestNextMemIndex(t *testing.T) {
	for i := 1; i < 3; i++ {
		require.Equal(t, i, nextMemIndex())
	}
}
