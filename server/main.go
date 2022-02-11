package main

import (
	"log"

	"github.com/sharpvik/rfip"
	"github.com/sharpvik/rfip/server/handler"
)

const addr = ":8000"

func main() {
	log.Println("server listening at", addr)
	err := rfip.NewServerWithResolver(handler.New()).ListenAndServe(addr)
	if err != nil {
		log.Fatal(err)
	}
}
