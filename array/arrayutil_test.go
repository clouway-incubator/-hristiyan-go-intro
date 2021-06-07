package array_test

import (
	"hristiyan-go-intro/array"
	"testing"
)

func TestMinElement(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	actual := array.MinElement(arr)
	if actual != 1 {
		t.Errorf("expected 1, got '%d", actual)
	}
}

func TestSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	actual := array.Sum(arr)
	if actual != 15 {
		t.Errorf("expected 15, got '%d", actual)
	}
}
