package main

import "fmt"

func main() {
	if 8%5 == 0 {
		fmt.Println("divisible")
	} else {
		fmt.Println("not divisible")
	}

	if i := 7; i%3 == 0 {
		fmt.Println("the initialized value is a factor of 3")
	} else if i%4 == 0 {
		fmt.Println("the initialized value is a factor of 4")
	} else {
		fmt.Println("the initialized value is something else")
	}
}
