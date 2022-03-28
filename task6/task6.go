package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateString(size int) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	res := make([]byte, size)
	for i := range res {
		res[i] = characters[rand.Intn(len(characters))]
	}
	return string(res)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(generateString(12))
}
