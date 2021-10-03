package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("file path")
		os.Exit(1)
	}

	path := os.Args[1]
	size, err := strconv.Atoi(os.Args[2])
	checkError(err)

	src, err := os.Open(path)
	checkError(err)

	defer src.Close()

	dest, err := os.Create(path + ".copy")
	checkError(err)

	defer dest.Close()

	buf := make([]byte, size)

	for {
		n, err := src.Read(buf)
		if n != 0 {
			dest.Write(buf[:n])
		}

		if err == io.EOF {
			break
		}

		checkError(err)
	}
}
