package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// START OMIT
	messages := make(chan string)

	go func() {
		limit := rand.Intn(10)
		for i := 0; i < limit; i++ {
			messages <- "Hello!"
		}
		close(messages) // HL
	}()

	for message := range messages { // HL
		fmt.Println(message)
	}
	// END OMIT
}
