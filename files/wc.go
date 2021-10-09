package main

import (
	"bufio"
	"fmt"
	"os"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	args := os.Args

	if len(args) != 2 {
		fmt.Println("usage: wc [option] <filename>")
		os.Exit(1)
	}

	filepath := args[1]

	f, err := os.Open(filepath)
	handleError(err)

	defer f.Close()

	lineS := bufio.NewScanner(f)
	// 전체 라인의 수 구분
	lineCount := 0

	for lineS.Scan() {
		lineCount++
	}

	// 첫 지점으로
	f.Seek(0, 0)

	// 전체 단어의 수 구분
	wordS := bufio.NewScanner(f)
	wordCount := 0
	wordS.Split(bufio.ScanWords)

	for wordS.Scan() {
		wordCount++
	}

	// 첫 지점으로
	f.Seek(0, 0)

	// 전체 문자의 수 구분
	byteS := bufio.NewScanner(f)
	byteCount := 0
	byteS.Split(bufio.ScanBytes)

	for byteS.Scan() {
		byteCount++
	}

	fmt.Println(lineCount, wordCount, byteCount, f.Name())

}
