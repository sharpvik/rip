package main

import (
	"log"

	"github.com/sharpvik/rip/example/random"
)

func main() {
	i, err := random.Service(random.Addr).RandInt()
	if err != nil {
		log.Fatalln("response error:", err)
	}
	log.Println("random integer:", i)
}
