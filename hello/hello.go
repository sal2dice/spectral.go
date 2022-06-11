package main

import (
	"fmt"
	"spectralstudio.com/messaging"
)

func main() {
	fmt.Println("Hello wrlds")
	//	fmt.Println(quote.Go())
	message := messaging.Hello("Gladys")
	fmt.Println(message)
}
