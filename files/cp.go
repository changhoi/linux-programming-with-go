package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func isExists(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func main() {
	flag.Parse()

	command := flag.Args()

	if len(command) == 0 {
		fmt.Printf("usage: %s <src/file/path> <dest/file/path>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	src, dest := command[0], command[1]

	// src 로 넘어온 경로에 파일이 없거나 폴더일 경우
	if !isExists(src) {
		fmt.Println("wrong file path or not a file")
		os.Exit(1)
	}

	in, err := os.Open(src)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// in 파일을 copy가 끝난 후에 닫아줌
	defer in.Close()

	out, err := os.Create(dest)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// out 파일을 copy가 끝난 후에 닫아줌
	defer func() {
		err := out.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	// copy 중에 오류가 생겼다면
	if _, err := io.Copy(out, in); err != nil {
		fmt.Println("error while copying files")
		os.Exit(1)
	}

}
