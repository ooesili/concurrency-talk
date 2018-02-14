package main

import (
	"fmt"
)

func main() {
	// START OMIT
	sum := 0
	ch := make(chan int) // HL

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				ch <- 1 // HL
			}
		}()
	}

	for i := 0; i < 100*100; i++ {
		sum += <-ch // HL
	}

	fmt.Println(sum)
	// END OMIT
}
