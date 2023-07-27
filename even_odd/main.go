package main

import (
	"fmt"
	"math/rand"
)

func main() {

	nums := rand.Perm(10)

	for _, n := range nums {
		if n%2 == 0 {
			fmt.Printf("O numero %v é par.\n", n)
		} else {
			fmt.Printf(" O numero %v é impar.\n", n)
		}

	}

}
