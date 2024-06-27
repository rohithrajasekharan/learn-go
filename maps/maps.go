package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("value1:", v1)
	v2 := m["k2"]
	fmt.Println("value2:", v2)

	fmt.Println("length:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k1"]

	fmt.Println("Present", prs)

	n := map[string]int{"k1": 4, "k2": 5}
	n2 := map[string]int{"k1": 4, "k2": 5}

	if maps.Equal(n, n2) {
		fmt.Println("n==n2")
	}
}
