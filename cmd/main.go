package main

import (
	"fmt"
	"os"

	"github.com/praveenmahasena647/gpsql/cmd/app"
)

func main() {
	var appErr = app.Exc()
	if appErr != nil {
		fmt.Printf("%v", appErr)
		os.Exit(1)
	}
	fmt.Println("sucessfully migrated")
}
