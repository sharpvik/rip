package main

import (
	"log"
	"net"

	"github.com/sharpvik/rfip"
)

const addr = ":8000"

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Println("connected to server at", addr)

	req, err := rfip.NewRequest("greet", map[string]string{"name": "Viktor"})
	if err != nil {
		panic(err)
	}

	_, err = req.Send(conn)
	if err != nil {
		panic(err)
	}
	log.Print("sent request:", req)

	resp, err := rfip.ReadResponse(conn)
	if err != nil {
		panic(err)
	}
	log.Println("received response:", resp)
}
