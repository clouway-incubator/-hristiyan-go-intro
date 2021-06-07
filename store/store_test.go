package store_test

import (
	"hristiyan-go-intro/store"
	"math"
	"testing"
)

func TestRegister(t *testing.T) {
	item := store.InventoryItem{"Apples", 5, 2.00, 2.30}
	s := store.Store{}
	s.Products = make(map[string]*store.InventoryItem)
	s.Register(item)

	actualLenght := len(s.Products)

	if actualLenght != 1 {
		t.Errorf("expected lenght of 1, got %d", actualLenght)
	}
}

func TestSell(t *testing.T) {
	item := store.InventoryItem{"Apples", 5, 2.00, 2.30}
	s := store.Store{}
	s.Products = make(map[string]*store.InventoryItem)
	s.Revenue = make(map[string]float64)
	s.Register(item)
	actualResult := s.Sell(item.ProductType, 3)

	actualCapacity := s.Products[item.ProductType].Capacity

	if actualCapacity != 2 {
		t.Errorf("expected 2, got %d", actualCapacity)
	}

	expectedResult := "order has 3 * Apples * 2.30"
	if actualResult != expectedResult {
		t.Errorf("expected %s, got %s", expectedResult, actualResult)
	}
}

func TestRevenueFrom(t *testing.T) {
	item := store.InventoryItem{"Apples", 5, 2.00, 2.30}
	s := store.Store{}
	s.Products = make(map[string]*store.InventoryItem)
	s.Revenue = make(map[string]float64)
	s.Register(item)
	s.Sell(item.ProductType, 3)

	actualRevenue := s.RevenueFrom(item.ProductType)
	actualRevenue = math.Round(actualRevenue*100) / 100
	if actualRevenue != 0.90 {
		t.Errorf("expected 0.90, got %f", actualRevenue)
	}
}
