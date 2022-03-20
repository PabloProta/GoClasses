//main package means that file is the main context
package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Pablo")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
