package main

import (
	"bufio"
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
	defer f.Close()
	r := bufio.NewReader(f)

	r.ReadLine()
	for {
		str, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		checkError(err)

		fmt.Print(str)
	}
}
