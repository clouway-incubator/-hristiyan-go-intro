package mathoper_test

import (
	"hristiyan-go-intro/mathoper"
	"testing"
)

func TestGCD(t *testing.T) {
	actual := mathoper.GCD(75, 125)
	if actual != 25 {
		t.Errorf("expected 25, got '%d'", actual)
	}
}

func TestLCM(t *testing.T) {
	actual := mathoper.LCM(4, 6)
	if actual != 12 {
		t.Errorf("expected 12, got '%d'", actual)
	}
}
