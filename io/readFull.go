package main

import (
	"fmt"
	"io"
	"os"
)

const path = "testfile"

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	f, err := os.Open(path)
	checkError(err)

	buf := make([]byte, 20)
	n, err := io.ReadFull(f, buf)
	checkError(err)
	fmt.Println("read:", n)

	fmt.Printf("%s\n", buf)
}
