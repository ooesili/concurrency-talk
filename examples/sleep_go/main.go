package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	fmt.Println("before")
	go time.Sleep(time.Second) // HL
	fmt.Println("after")
	// END OMIT
}
