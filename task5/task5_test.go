package main

import "testing"

func TestReverseArray(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	actual := reverseArray(data)
	for i := range data {
		if actual[i] != expected[i] {
			t.Errorf("expected %d, got %d", expected[i], actual[i])
		}
	}

}
