package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 9; i++ {
		fmt.Println(rand.Int())
	}
}