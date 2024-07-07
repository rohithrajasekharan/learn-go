package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"fruits", "banana", "kiwi"}
	lenComp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenComp)
	fmt.Println("Fruits:", fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "Josh", age: 45},
		{name: "Sara", age: 20},
		{name: "Sid", age: 34},
	}
	ageComp := func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	}

	slices.SortFunc(people, ageComp)
	fmt.Println("People:", people)
}
