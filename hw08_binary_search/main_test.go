package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{}, 5, -1}, // empty array
		{[]int{1}, 1, 0}, // one element
		{[]int{1}, 2, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1, -1}, // less than the minimum
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, -1}, // larget than the maximum
	}

	for _, tt := range tests {
		got := BinarySearch(tt.arr, tt.target)
		assert.Equal(t, tt.want, got, "BinarySearch(%v, %d)", tt.arr, tt.target)
	}
}
