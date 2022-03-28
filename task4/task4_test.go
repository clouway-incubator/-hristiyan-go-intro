package main

import (
	"testing"
)

func TestQuicksort(t *testing.T) {
	slice := []int{3, 6, 2, 8, 1, 9, 7, 4, 5}
	actual := quicksort(slice)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range slice {
		if actual[i] != expected[i] {
			t.Errorf("expected %d, got %d", expected[i], actual[i])
		}
	}
}
