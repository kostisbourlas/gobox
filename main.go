package main

import (
	"log"

	"github.com/kostisbourlas/gobox/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3030")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
