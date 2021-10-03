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
	src, err := os.Open(path)
	checkError(err)
	defer src.Close()

	dest, err := os.Create(path + ".copy")
	checkError(err)
	defer dest.Close()

	n, err := io.Copy(dest, src)
	checkError(err)
	fmt.Println("bytes", n)
}
