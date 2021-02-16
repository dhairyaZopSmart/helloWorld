package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 8; i++ {
		fmt.Println(rand.Int())
	}
}