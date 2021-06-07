package store

import (
	"fmt"
	"strconv"
)

type InventoryItem struct {
	ProductType string
	Capacity    int
	BuyPrice    float64
	SellPrice   float64
}

type Store struct {
	inventoryItem InventoryItem
	Products      map[string]*InventoryItem
	Revenue       map[string]float64
}

func (s *Store) Register(item InventoryItem) {
	s.inventoryItem = item
	s.Products[item.ProductType] = &item
}

func (s *Store) Sell(item string, amount int) string {
	product := s.Products[item]
	if product.Capacity >= amount {
		s.Products[item].Capacity -= amount
	}
	s.Revenue[item] = float64(amount) * (product.SellPrice - product.BuyPrice)

	return "order has " + fmt.Sprint(amount) + " * " + item + " * " + strconv.FormatFloat(product.SellPrice, 'f', 2, 64)
}

func (s *Store) RevenueFrom(item string) float64 {
	return s.Revenue[item]
}

func (s *Store) ShowProductSortedByPrice() {
	sortedProducts := make([]string, len(s.Products))
	sortedPrices := make([]float64, len(s.Products))

	for k, v := range s.Products {
		sortedProducts = append(sortedProducts, k)
		sortedPrices = append(sortedPrices, v.SellPrice)
	}

	for i := 0; i < len(sortedPrices); i++ {
		max := i
		for j := i; j < len(sortedPrices); j++ {
			if sortedPrices[max] < sortedPrices[j] {
				max = j
			}
		}
		temp := i
		sortedPrices[i] = sortedPrices[max]
		sortedProducts[i] = sortedProducts[max]
		sortedPrices[max] = sortedPrices[temp]
		sortedProducts[max] = sortedProducts[temp]
	}

	for _, v := range sortedProducts {
		fmt.Println(v)
	}
}
