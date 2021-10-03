package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("usage: permission <filepath>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	file, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("there is no file named %s\n", filePath)
		os.Exit(1)
	}

	fmt.Printf("%s : %s\n", filePath, file.Mode())
}
