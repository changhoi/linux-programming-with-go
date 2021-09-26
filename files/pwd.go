package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	minusP := flag.Bool("P", false, "피지컬 링크")
	flag.Parse()
	fmt.Println(*minusP)

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !*minusP {
		fmt.Println(dir)
	} else {
		//
	}
}
