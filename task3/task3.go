package main

import (
	"fmt"

	"hristiyan-go-intro/array"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(array.MinElement(arr))
	fmt.Println(array.Sum(arr))
	array.Print(arr)
}
