package main

import (
	"log"

	"github.com/checkr/openmock/app"
)

func main() {
	s := &app.Server{
		Addr: ":8080",
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
