package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {

	talha := User{"Talha", "talha@abc.com", 23}
	fmt.Printf("%v\n", talha)
	fmt.Printf("%+v\n", talha)
	fmt.Printf("%v\n", talha.Email)

	var fahim = new(User)
	fahim.Name = "Fahim"
	fahim.Email = "fahim@abc.com"
	fahim.Age = 23

	fmt.Printf("%v\n", fahim)
	fmt.Printf("%+v\n", fahim)
	fmt.Printf("%v\n", fahim.Email)

	var rashid = &User{"Rashid", "rashid@abc.com", 24}
	fmt.Printf("%v\n", rashid)
	fmt.Printf("%+v\n", rashid)
	fmt.Printf("%v\n", rashid.Email)

}
