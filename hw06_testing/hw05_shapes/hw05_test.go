package hw05

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculateAreaCircle(t *testing.T) {
	const want float64 = 78.53981633974483

	myCircle := &circle{5}
	area, err := CalculateArea(myCircle)
	require.NoError(t, err)
	assert.Equal(t, want, area)
}

func TestCalculateAreaEmptyCircle(t *testing.T) {
	const want float64 = 0

	myCircle := &circle{}
	area, err := CalculateArea(myCircle)
	require.NoError(t, err)
	assert.Equal(t, want, area)
}

func TestCalculateAreaRectangle(t *testing.T) {
	const want float64 = 50

	myRectangle := &rectangle{10, 5}
	area, err := CalculateArea(myRectangle)
	require.NoError(t, err)
	assert.Equal(t, want, area)
}

func TestCalculateAreaEmptyRectangle(t *testing.T) {
	const want float64 = 0

	myRectangle := &rectangle{}
	area, err := CalculateArea(myRectangle)
	require.NoError(t, err)
	assert.Equal(t, want, area)
}

func TestCalculateAreaTriangle(t *testing.T) {
	const want float64 = 24

	myTriangle := &triangle{8, 6}
	area, err := CalculateArea(myTriangle)
	require.NoError(t, err)
	assert.Equal(t, want, area)
}

func TestCalculateAreaEmptyTriangle(t *testing.T) {
	const want float64 = 0

	myTriangle := &triangle{}
	area, err := CalculateArea(myTriangle)
	require.NoError(t, err)
	assert.Equal(t, want, area)
}

func TestCalculateAreaNotShape(t *testing.T) {
	const want float64 = 0
	wanterr := errors.New("this type of shape isn't available")

	notShape := "not a shape"
	area, err := CalculateArea(notShape)
	assert.Equal(t, wanterr, err)
	assert.Equal(t, want, area)
}
