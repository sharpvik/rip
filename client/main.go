package main

import (
	"log"

	"github.com/sharpvik/rfip/server/handler"
)

const addr = ":8000"

func main() {
	msg := handler.NewAPI(addr).Greet("Viktor")
	log.Println("received response:", msg)
}
