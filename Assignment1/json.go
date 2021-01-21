package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	e1 := employee{
		FirstName: "Abdul",
		LastName:  "Khader",
		Age:       24,
	}

	e2 := employee{"Mohan", "Kumar", 37}

	employees := []employee{e1, e2}
	fmt.Println(employees)

	bs, err := json.Marshal(employees)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
