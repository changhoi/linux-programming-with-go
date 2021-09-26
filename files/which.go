package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
	which
	1. path/파일 이름 실존 하는가?
	2. 실행 가능한 파일인가?
*/

func main() {
	minusA := flag.Bool("a", false, "all")
	flag.Parse()

	command := flag.Args()
	if len(command) == 0 {
		fmt.Printf("usage: %s [-a] <command>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	pathString := os.Getenv("PATH")
	path := strings.Split(pathString, ":")

	success := false
	for _, dir := range path {
		fullPath := dir + "/" + command[0]
		f, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := f.Mode()
		if mode.IsRegular() && mode&0111 != 0 {
			fmt.Println(fullPath)
			success = true
			if !*minusA {
				break
			}
		}
	}

	if !success {
		fmt.Printf("%s not found", command[0])
	}
}
