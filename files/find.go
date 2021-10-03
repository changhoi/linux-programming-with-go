package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

/*
	정규표현식 제외한 -name 기준
	스페셜 f, l, d
*/

func main() {
	minusName := flag.String("name", "", "name")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		os.Exit(1)
	}

	root := args[0]

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		format := filepath.Dir(root) + "/" + path

		if *minusName != "" {
			if filepath.Base(path) == *minusName {
				fmt.Println(format)
			}
		} else {
			fmt.Println(format)
		}
		return nil
	})
}
