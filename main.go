package main

import (
	"fmt"
	"log"
	"time"

	"golang-ninja/httpclient/got"
)

func main() {
	coincapClient, err := got.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	hero, err := coincapClient.GetAsset("Daenerys")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hero.Info())
}
