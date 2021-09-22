package main

import (
	"fmt"
	"log"

	gecko "github.com/shortcircuit3/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	list, err := cg.CoinsList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available coins:", len(*list))
}
