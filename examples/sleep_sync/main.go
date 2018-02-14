package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	fmt.Println("before")
	time.Sleep(time.Second)
	fmt.Println("after")
	// END OMIT
}
