package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
)

/*
	정규표현식 제외한 -name 기준
	스페셜 f, l, d
*/

func main() {
	minusName := flag.String("name", "", "name")
	minusUser := flag.String("user", "", "user")

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		os.Exit(1)
	}

	root := args[0]

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		format := filepath.Dir(root) + "/" + path

		info, err := os.Stat(path)
		uid := info.Sys().(*syscall.Stat_t).Uid
		conv := strconv.FormatUint(uint64(uid), 10)

		user, err := user.LookupId(conv)
		userName := user.Username


		if *minusName != "" && *minusUser != "" {
			// name과 user 다 사용할때
			if filepath.Base(path) == *minusName && userName == *minusUser {
				fmt.Println(format)
			}
		} else if *minusName != "" && *minusUser == "" {
			// name만 사용할때
			if filepath.Base(path) == *minusName {
				fmt.Println(format)
			}
			fmt.Println(format)
		} else if *minusName == "" && *minusUser != "" {
			// user만 사용할때
			if userName == *minusUser {
				fmt.Println(format)
			}
		} else {
			fmt.Println(format)
		}
		return nil
	})
}
