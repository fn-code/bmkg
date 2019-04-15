package main

import (
	"fmt"
	"log"

	"github.com/fn-code/bmkg"
)

const (
	baseUrl         = "http://data.bmkg.go.id/"
	gempaTerkiniUrl = "gempaterkini.xml"
)

func main() {

	info, err := bmkg.GempaTerkini(baseUrl + gempaTerkiniUrl)
	if err != nil {
		log.Printf("failed get gempa terkini:%v\n", err)
	}
	for _, v := range info.Gempa {
		fmt.Println(v)
	}
}
