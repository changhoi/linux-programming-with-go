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

	bytes, err := io.ReadAll(src)
	checkError(err)

	err = os.WriteFile(path+".copy", bytes, 0644)
	checkError(err)
}
