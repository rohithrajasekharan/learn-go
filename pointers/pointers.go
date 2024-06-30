package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeropntr(ipntr *int) {
	*ipntr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeropntr(&i)
	fmt.Println("zeropntr:", i)

	fmt.Println("pointer:", &i)
}
