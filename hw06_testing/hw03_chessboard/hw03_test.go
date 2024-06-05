package hw03

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChessboardEmpty(t *testing.T) {
	const want = ""

	got, err := GetStrChessboard(0, 0)
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestChessboardPositive(t *testing.T) {
	const want = " # # \n# # #\n # # \n# # #\n # # "

	got, err := GetStrChessboard(5, 5)
	require.NoError(t, err)
	assert.Equal(t, want, got)
}
