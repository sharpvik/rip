package main

import (
	"log"

	"github.com/sharpvik/rip"
	"github.com/sharpvik/rip/server/handler"
)

const addr = ":8000"

func main() {
	log.Println("server listening at", addr)
	err := rip.NewServerWithResolver(handler.New()).ListenAndServe(addr)
	if err != nil {
		log.Fatal(err)
	}
}
