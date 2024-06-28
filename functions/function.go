package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func multi() (int, int) {
	return 3, 7
}

func variadic(nums ...int) {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total:", total)
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)

	a, b := multi()
	fmt.Println(a, b)

	_, c := multi()
	fmt.Println("c:", c)

	variadic(1, 2, 3, 4)
}
