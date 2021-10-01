package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	minusP := flag.Bool("P", false, "피지컬 링크")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !*minusP {
		fmt.Println(dir)
	} else {
		readlink, err := os.Readlink(dir)
		if err != nil {
			log.Fatal(err)
		}
		err = os.Chdir(readlink)
		if err != nil {
			log.Fatal(err)
		}
		dir, err = os.Getwd()
		fmt.Println(dir)
	}
}
