package main

import (
	"fmt"
	"time"
)

func main() {
	v := 3
	switch v {
	case 3:
		fmt.Println("The variable is 3")
	case 2:
		fmt.Println("The variable is 2")
	default:
		fmt.Println("The variable is something else")
	}

	t := time.Now()
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its a weekend")
	default:
		fmt.Println("Its a weekday")
	}

	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Printf("The type of param %T is boolean", t)
		case int:
			fmt.Printf("The type of param %T is integer", t)
		case string:
			fmt.Printf("The type of param %T is string", t)
		default:
			fmt.Println("Unknown type")
		}
	}

	whatAmI(true)
	whatAmI(4)
	whatAmI("hey")
	whatAmI(4.0)
}
