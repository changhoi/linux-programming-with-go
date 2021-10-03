package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

/*
	1. 소스 파일과 목적지를 받음 (하나씩)
	2. 소스 파일이 존재 -> 없으면 에러
	3. 목적지가
		1. 디렉토리면, 그 아래 파일을 붙인 이름으로 파일을 확인
			1. 파일이 이미 있으면 덮어씌움 옵션 확인
		2. 파일이면, 해당 패스로 같은 파일이 있는지 확인
			1. 같은 파일이 있다면 Overwrite 옵션이 켜진 경우만 덮어씌움, 아니면 에러
			2. 없으면 그냥 그 자리에 파일을 옮겨놓음
		3. 없으면, Rename 하려는 의도, 그냥 수행되도록 내려보내면 됨
*/

func main() {
	minusO := flag.Bool("overwrite", false, "overwrite")
	flag.Parse()

	args := flag.Args()

	if len(args) != 2 {
		fmt.Println("usage: mv <src> <dest>")
		os.Exit(1)
	}

	srcFile := args[0]
	destFile := args[1]

	// 소스 파일 있는지, 파일 타입 가져오기
	f, err := os.Stat(srcFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	newPath := destFile

	f, err = os.Stat(destFile)

	// 목적지가 존재하고, 디렉토리면 그 아래에 파일 넣기
	if err == nil && f.IsDir() {
		newPath += "/" + filepath.Base(srcFile)
	}

	_, err = os.Stat(newPath)

	// 목적지가 이미 존재하고, 오버라이트가 꺼져있다면 에러
	if err == nil && !*minusO {
		fmt.Println("Destination file already exists!")
		os.Exit(1)
	}

	// 옮겨주기
	err = os.Rename(srcFile, newPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
