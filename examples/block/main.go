package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	messages := make(chan string)

	go func() {
		time.Sleep(time.Second)
		messages <- "Don't put ketchup on your hot dog"
	}()

	message := <-messages
	fmt.Println(message)
	// END OMIT
}
