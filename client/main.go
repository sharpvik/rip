package main

import (
	"log"

	"github.com/sharpvik/rfip"
)

const addr = ":8000"

func main() {
	resp, err := rfip.NewClient(addr).Invoke("Greet", 22)
	if err != nil {
		panic(err)
	}
	log.Println("received response:", resp)
}
