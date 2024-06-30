package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 28})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 23})

	fmt.Println(newPerson("John"))

	s := person{name: "Sean", age: 70}
	fmt.Println("Sean's age", s.age)

	sp := &s
	fmt.Println(sp.name)

	sp.age = 41
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		false,
	}

	fmt.Println("Rex info:", dog)
}
