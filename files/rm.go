package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("usage: rm <filename>")
		os.Exit(1)
	}

	file := os.Args[1]
	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
		return
	}
}
