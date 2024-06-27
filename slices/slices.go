package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninitialized:", s, s == nil, len(s))

	s = make([]string, 3, 5)
	fmt.Println("initialized:", s, len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	fmt.Println("append:", s)
	s = append(s, "e", "f", "g")
	fmt.Println("append:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)
	l = s[2:]
	fmt.Println("sl2:", l)
	l = s[:5]
	fmt.Println("sl3:", l)

	t := []string{"a", "b", "c"}
	t2 := []string{"a", "b", "c"}

	if slices.Equal(t, t2) {
		fmt.Println("t==t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}
