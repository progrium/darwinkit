package main

import (
	"log"
	"os"

	"github.com/progrium/macdriver/bridge"
)

func main() {
	h, err := bridge.NewHost(os.Stderr)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(h.Run())
}
