package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func isExists(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <target/file/path>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	path := os.Args[1]

	if !isExists(path) {
		fmt.Println("file not exists")
		os.Exit(1)
	}

	err := os.Remove(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
