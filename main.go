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

	hero, err := coincapClient.GetAsset("2baf70d1-42bb-4437-b551-e5fed5a87abe")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hero.Info())
}
