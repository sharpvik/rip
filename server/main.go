package main

import (
	"log"
	"net"

	"github.com/sharpvik/rfip"
)

const addr = ":8000"

func main() {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Println("server listening at", addr)

	conn, err := l.Accept()
	if err != nil {
		panic(err)
	}
	log.Print("new connection")

	r, err := rfip.ReadRequest(conn)
	if err != nil {
		panic(err)
	}
	log.Println("received request:", r)

	_, err = rfip.ResponseString("Hello there").Send(conn)
	if err != nil {
		panic(err)
	}
	log.Print("response sent")
}
