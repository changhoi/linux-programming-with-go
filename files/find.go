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

type findHandler struct {
	Args map[string]string
	Path string
	Root string
}

func (h *findHandler) handle() string {
	return h.handleName()
}

func (h *findHandler) handleName() string {
	nameArg := h.Args["name"]

	// name arg가 없을때 혹은 파일이름과 일치하면
	if nameArg == "" || filepath.Base(h.Path) == nameArg {
		return h.handleUser()
	}
	// 파일이름이 일치하지 않으면
	return ""
}

func (h *findHandler) handleUser() string {
	userArg := h.Args["user"]

	// 파일의 정보를 받아옴
	info, err := os.Stat(h.Path)
	if err != nil {
		fmt.Println("error at stat")
		return ""
	}
	// 파일의 uid 가져옴
	uid := info.Sys().(*syscall.Stat_t).Uid
	// 파일의 uid를 string으로 변환
	str := strconv.FormatUint(uint64(uid), 10)

	user, err := user.LookupId(str)
	userName := user.Username

	// user arg가 없을때 혹은 파일의 user가 user arg와 일치하면
	if userArg == "" || userName == userArg {
		return h.handleGroup()
	}
	// 파일의 user가 일치하지 않으면
	return ""

}

func (h *findHandler) handleGroup() string {
	groupArg := h.Args["group"]

	// 파일의 정보를 받아옴
	info, err := os.Stat(h.Path)
	if err != nil {
		fmt.Println("error at stat")
		return ""
	}
	// 파일의 gid 가져옴
	gid := info.Sys().(*syscall.Stat_t).Gid
	// 파일의 gid를 string으로 변환
	str := strconv.FormatUint(uint64(gid), 10)

	group_, err := user.LookupGroupId(str)
	groupName := group_.Name

	// group arg가 없을때 혹은 파일의 groupname이 group arg와 일치하면
	if groupArg == "" || groupName == groupArg {
		return h.handleType()
	}
	// 파일의 group이 일치하지 않으면
	return ""
}

func (h *findHandler) handleType() string {
	typeArg := h.Args["type"]

	// type에 대한 arg가 없으면 path 반환
	if typeArg == "" {
		return h.Path
	}

	// symbolic link를 걸러야 하니 lstat 사용
	f, err := os.Lstat(h.Path)

	if err != nil {
		fmt.Println("error at stat")
		return ""
	}
	switch typeArg {
	case "f":
		if f.Mode().IsRegular() {
			return h.Path
		}
	case "l":
		// 파일이 symbolic link면
		if f.Mode()&os.ModeSymlink == os.ModeSymlink {
			return h.Path
		}

	case "d":
		if f.IsDir() {
			return h.Path
		}
	}
	return ""

}

func getHandler(path string, root string) *findHandler {
	return &findHandler{
		Path: path,
		Root: root,
		Args: make(map[string]string),
	}
}

func main() {
	minusName := flag.String("name", "", "name")
	minusUser := flag.String("user", "", "user")
	minusGroup := flag.String("group", "", "group")
	minusType := flag.String("type", "", "type")

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		os.Exit(1)
	}

	root := args[0]

	h := getHandler("", "")

	h.Args["name"] = *minusName
	h.Args["user"] = *minusUser
	h.Args["group"] = *minusGroup
	h.Args["type"] = *minusType

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {

		/**
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
		**/
		h.Path = path
		h.Root = root

		if ret := h.handle(); ret != "" {
			format := filepath.Dir(h.Root) + "/" + h.Path
			fmt.Println(format)
		}
		return nil
	})
}
