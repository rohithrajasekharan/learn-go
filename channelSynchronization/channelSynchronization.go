package main

import (
	"fmt"
	"time"
)

func Worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	Worker(done)
	<-done
}
