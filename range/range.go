package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("The index is:", i)
		}
	}

	kvs := map[string]string{"k1": "apple", "k2": "orange", "k3": "banana"}
	for key, val := range kvs {
		fmt.Printf("%s -> %s\n", key, val)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}
	for i, c := range "rohith" {
		fmt.Println(i, c)
	}
}
