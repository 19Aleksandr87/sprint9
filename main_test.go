package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		len    int
		result int
	}{
		{
			len:    0,
			result: 0,
		},
		{
			len:    1,
			result: 1,
		},
		{
			len:    -5,
			result: 0,
		},
	}
	for _, i := range tests {
		x := generateRandomElements(i.len)
		assert.Equal(t, i.result, len(x))
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		data []int
		max  int
	}{
		{
			data: []int{},
			max:  0,
		},
		{
			data: []int{0},
			max:  0,
		},
		{
			data: []int{1},
			max:  1,
		},
		{
			data: []int{-5},
			max:  -5,
		},
		{
			data: []int{-5, -1},
			max:  -1,
		},
		{
			data: []int{-5, 0},
			max:  0,
		},
		{
			data: []int{5, -5},
			max:  5,
		},
		{
			data: []int{7, 0, 19},
			max:  19,
		},
		{
			data: []int{19, 7, 19},
			max:  19,
		},
	}

	for j, i := range tests {
		max := maximum(i.data)
		assert.Equal(t, i.max, max, j)
	}

}
