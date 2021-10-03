package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	f, err := os.Create("file")
	if err != nil {
		fmt.Println("err")
	}

	fmt.Println("Hello STDIO!")

	defer f.Close()

	syscall.Close(1)
	fd, err := syscall.Dup(int(f.Fd()))
	fmt.Println("Hello World!")
	fmt.Printf("This is on %d", fd)
}
