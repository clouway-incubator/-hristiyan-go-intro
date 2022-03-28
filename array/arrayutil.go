package array

import "fmt"

// Returns the smallest element of the slice
func MinElement(array []int) int {
	min := array[0]

	for i := 0; i < len(array); i++ {
		if min > array[i] {
			min = array[i]
		}
	}
	return min
}

// Returns the sum of all elements in the slice
func Sum(array []int) int {
	sum := 0

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

// Prints all elements of the slice to the screen
func Print(array []int) {
	fmt.Println(array)
}
