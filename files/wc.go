package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("need argument")
		os.Exit(1)
	}

	paths := flag.Args()

	for _, path := range paths {
		words := 0
		lines := 0
		var size int64 = 0
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		info, err := (*file).Stat()
		if err != nil {
			log.Fatal(err)
		}
		size = info.Size()

		wordScanner := bufio.NewScanner(file)
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			words++
		}

		file.Seek(0, 0)

		lineScanner := bufio.NewScanner(file)
		lineScanner.Split(bufio.ScanLines)
		for lineScanner.Scan() {
			lines++
		}

		fmt.Println(fmt.Sprint(lines, " ", words, " ", size))
	}
}
