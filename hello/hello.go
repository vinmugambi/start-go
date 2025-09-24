package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main() {
	log.SetPrefix("Greetings")
	log.SetFlags(0)

	message, err := greetings.Hello("Vincent")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}