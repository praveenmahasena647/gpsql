package main

import (
	"fmt"
	"uri/cmd/internal"
)

func main() {
	var err error = internal.Start()
	if err != nil {
		fmt.Printf("%v", err)
	}
}
