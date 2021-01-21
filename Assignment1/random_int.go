package main

import (
	"fmt"
	"math/rand"
)

func main() {

	num1 := rand.Intn(5)
	num2 := rand.Intn(3)
	num3 := rand.Intn(9)

	fmt.Println("Random No.1: ", num1)
	fmt.Println("Random No.2: ", num2)
	fmt.Println("Random No.3: ", num3)

}
