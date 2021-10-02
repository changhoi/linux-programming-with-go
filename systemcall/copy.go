package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("path, bufsize")
		os.Exit(1)
	}

	path := os.Args[1]
	size, err := strconv.Atoi(os.Args[2])
	checkError(err)

	srcFd, err := syscall.Open(path, syscall.O_RDONLY, 0644)
	checkError(err)
	defer syscall.Close(srcFd)
	destFd, err := syscall.Open(path+".copy", syscall.O_RDWR|syscall.O_TRUNC|syscall.O_CREAT, 0644)
	checkError(err)
	defer syscall.Close(destFd)

	buf := make([]byte, size)
	for {
		n, err := syscall.Read(srcFd, buf)
		if n != 0 {
			syscall.Write(destFd, buf[:n])
		}

		if n == 0 {
			break
		}
		checkError(err)
	}
}
