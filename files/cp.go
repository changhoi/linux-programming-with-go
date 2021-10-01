package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: cp source_file target_file")
		os.Exit(64)
	}

	source := os.Args[1]
	target := os.Args[2]

	in, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}

	defer func(in *os.File) {
		err := in.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(in)

	out, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		log.Fatal(err)
	}
	err = out.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
