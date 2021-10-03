package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func isFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func checkOptions(options []string) bool {
	if len(options) != 2 {
		fmt.Println("usage: cp source_file_path target_file_path")
		return false
	}
	return true
}

func main() {
	flag.Parse()
	options := flag.Args()

	// 인자가 2개 나란히 넘어왓는지 검사
	if !checkOptions(options) {
		os.Exit(1)
	}
	source, target := options[0], options[1]

	// copy 를 수행할 파일이 path 위치 상에 존재하는지 검사
	// 파일이 존재하면, 해당 파일 열고, 아니면 에러 출력하며 종료
	sourceFile, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer func() {
		err := sourceFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// copy 목적지 파일이 path 위치 상에 이미 존재하는지 검사
	// 존재하면 에러 출력하며 종료
	if isFileExists(target) {
		err = fmt.Errorf("cp: target file already exists (not copied).")
		log.Fatal(err)
		os.Exit(3)
	}

	// copy 목적지 파일을 새로 생성
	targetFile, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
		os.Exit(4)
	}
	defer func() {
		err := targetFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// 복사 실시
	// 복사 중 오류 발생했으면 에러 발생
	_, err = io.Copy(targetFile, sourceFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(5)
	}
}
