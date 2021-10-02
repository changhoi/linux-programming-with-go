package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.OpenFile("textfile", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	str := strings.Repeat("test string value!\n", 100)
	f.Write([]byte(str))

	buf := make([]byte, 19, 19)
	f.Seek(0, 0)
	for {
		n, err := f.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print(string(buf[:n]))
	}
}
