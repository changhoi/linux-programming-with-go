package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: rm file")
		os.Exit(64)
	}
	err := os.Remove(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
