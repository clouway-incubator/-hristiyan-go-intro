package main

import "fmt"

func reverseArray(data []int) []int {
	l := len(data)
	for i := 0; i < l/2; i++ {
		temp := data[i]
		data[i] = data[l-1-i]
		data[l-1-i] = temp
	}

	return data
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(reverseArray(data))
}
