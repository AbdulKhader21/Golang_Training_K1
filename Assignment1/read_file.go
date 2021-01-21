package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	ctr := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line %d: %s \n", ctr, line)
		ctr++
	}
}
