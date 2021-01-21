package main

import (
	"fmt"
)

type Student struct {
	rollNo int
	name   string
	marks  int
}

func main() {

	var s1 Student = Student{14202, "ABDUL KHADER P", 744}
	fmt.Println(s1)
	fmt.Println(s1.name)

	var s2 = Student{rollNo: 14201, marks: 1000, name: "ABBU BACKAR A"}
	fmt.Println(s2)
}
