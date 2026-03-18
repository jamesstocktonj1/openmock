package main

import (
	"log"

	"github.com/checkr/openmock/app"
)

func main() {
	s, err := app.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
