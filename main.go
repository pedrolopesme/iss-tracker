package main

import (
	"log"
	"fmt"
	"github.com/pedrolopesme/iss-notifier/iss"
)

func main() {

	// Retrieving coordinate
	issCoordinate, err := iss.GetCoordinate()
	if err != nil {
		log.Fatal("iss.GetCoordinate: ", err)
	}
	fmt.Println(issCoordinate)

}
