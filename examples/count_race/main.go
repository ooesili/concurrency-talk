package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	sum := 0

	for i := 0; i < 100; i++ {
		go func() { // HL
			for j := 0; j < 100; j++ {
				sum = sum + 1
			}
		}() // HL
	}

	time.Sleep(time.Second) // HL
	fmt.Println(sum)
	// END OMIT
}
