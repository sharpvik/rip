package main

import (
	"log"

	"github.com/sharpvik/rip/server/handler"
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

	msg = api.Users(handler.UserParams{
		N:          5,
		NameLength: 10,
	})
	log.Println("users response:", msg)
}
