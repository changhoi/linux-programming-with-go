package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
)

/*
	정규표현식 제외한 -name 기준
	스페셜 f, l, d
*/

func main() {
	nameFlag := flag.String("name", "", "name")
	userFlag := flag.String("user", "", "user")
	groupFlag := flag.String("group", "", "group")
	typeFlag := flag.String("type", "", "type")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		os.Exit(1)
	}

	root := args[0]

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		format := root + "/" + path
		info, err := os.Stat(format)
		if err != nil {
			log.Fatal(err)
		}

		if *nameFlag != "" {
			if filepath.Base(path) != *nameFlag {
				format = ""
			}
		}

		if *userFlag != "" {
			ownerName := ""
			uid := fmt.Sprint(info.Sys().(*syscall.Stat_t).Uid)
			owner, err := user.LookupId(uid)
			if err == nil {
				ownerName = owner.Username
			}
			if ownerName != *userFlag {
				format = ""
			}
		}

		if *groupFlag != "" {
			groupName := ""
			gid := fmt.Sprint(info.Sys().(*syscall.Stat_t).Gid)
			group, err := user.LookupGroupId(gid)
			if err == nil {
				groupName = group.Name
			}
			if groupName != *groupFlag {
				format = ""
			}
		}

		if *typeFlag != "" {
			typeInfo, err := os.Lstat(format)
			if err != nil {
				log.Fatal(err)
			}
			if *typeFlag == "f" && !typeInfo.Mode().IsRegular() {
				format = ""
			}
			if *typeFlag == "l" && !(typeInfo.Mode()&os.ModeSymlink == os.ModeSymlink) {
				format = ""
			}
			if *typeFlag == "d" && !typeInfo.Mode().IsDir() {
				format = ""
			}
		}

		if format != "" {
			fmt.Println(format)
		}
		return nil
	})
}
