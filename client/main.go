package main

import (
	"log"

	"github.com/sharpvik/rfip/server/handler"
)

const addr = ":8000"

func main() {
	api := handler.NewAPI(addr)
	msg := api.Greet("Viktor")
	log.Println("received response:", msg)
	msg = api.MustInvoke("DoesNotExist", nil)
	if err := msg.Err(); err != nil {
		log.Println(err.Status(), err.Error())
	}
}
