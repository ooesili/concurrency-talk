package main

import (
	"fmt"
)

func main() {
	// START OMIT
	sum := 0

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			sum = sum + 1
		}
	}

	fmt.Println(sum)
	// END OMIT
}
