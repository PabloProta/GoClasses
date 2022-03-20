//main package means that file is the main context
package main

import (
	"example/greetings"

	"fmt"
)

func main() {
	message := greetings.Hello("Prota")
	fmt.Println(message)
}
