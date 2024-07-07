package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"a", "c", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{5, 4, 8}
	slices.Sort(ints)
	fmt.Println("Integers:", ints)

	isSorted := slices.IsSorted(ints)
	fmt.Println("Already Sorted?", isSorted)
}
