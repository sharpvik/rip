package main

import (
	"fmt"
	"log"

	"github.com/sharpvik/rfip"
)

const addr = ":8000"

type Handler struct{}

func (h *Handler) Greet(age int) string {
	return fmt.Sprintf("Hello, stranger, you are %d y.o.", age)
}

func main() {
	log.Println("server listening at", addr)
	err := rfip.NewServerWithResolver(new(Handler)).ListenAndServe(addr)
	if err != nil {
		panic(err)
	}
}
