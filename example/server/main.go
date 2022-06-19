package main

import (
	"log"

	"github.com/sharpvik/rip/example/random"
)

func main() {
	log.Println("server listening at", random.Addr)
	if err := random.Server().ListenAndServeTCP(random.Addr); err != nil {
		log.Fatal(err)
	}
}
