package main

import (
	"log"

	"clean-portfolio/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
