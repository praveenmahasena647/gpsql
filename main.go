package main

import (
	"log"

	"github.com/praveenmahasena647/gpsql/cmd"
)

func main() {
	var appErr = cmd.Start()
	if appErr != nil {
		log.Fatal(appErr)
	}
	log.Println("successfully migrated")
}
