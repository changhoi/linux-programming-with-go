package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func isFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func checkOptions(options []string) bool {
	if len(options) != 1 {
		fmt.Println("usage: rm target_file_path")
		return false
	}
	return true
}

func main() {
	flag.Parse()
	options := flag.Args()

	// 인자가 1개 넘어왔는지 검사
	if !checkOptions(options) {
		os.Exit(1)
	}
	target := options[0]

	// 해당 파일이 존재하는지 검사
	if !isFileExists(target) {
		err := fmt.Errorf("cp: target file does not exist - nothing to remove")
		log.Fatal(err)
		os.Exit(1)
	}

	// 해당 파일 제거
	err := os.Remove(target)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
