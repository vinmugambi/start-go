package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main() {
	log.SetPrefix("Greetings")
	log.SetFlags(0)

	names := []string{"Vincent", "Mugambi", "Wambui"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}