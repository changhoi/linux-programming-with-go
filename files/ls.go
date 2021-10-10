package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
	"syscall"
)

func applyAllFlag(results []string, longFlag bool) []string {
	if longFlag != true {
		results = append([]string{".", ".."}, results...)
	}
	return results
}

func applyLongFlag(file *os.FileInfo) string {
	fileStat := (*file).Sys().(*syscall.Stat_t)
	perms := ""              // 권한
	var linkCount uint64 = 0 // 링크 개수
	ownerName := ""          // 소유자
	groupName := ""          // 그룹
	var size int64 = 0       // 파일 크기
	modTime := ""            // 수정 일시
	fileName := ""           // 파일명

	// 권한
	perms = (*file).Mode().String()

	// 하드 링크 개수
	linkCount = uint64(fileStat.Nlink)

	// 소유자, 그룹
	uid := fmt.Sprint(fileStat.Uid)
	gid := fmt.Sprint(fileStat.Gid)
	owner, err := user.LookupId(uid)
	if err != nil {
		ownerName = uid
	} else {
		ownerName = owner.Username
	}
	group, err := user.LookupGroupId(gid)
	if err != nil {
		groupName = gid
	} else {
		groupName = group.Name
	}

	// 파일 크기
	size = (*file).Size()

	// 수정 일시
	modTime = (*file).ModTime().Format("01 02 03:04")

	// 파일명
	fileName = (*file).Name()

	return fmt.Sprint(perms, " ", linkCount, " ", ownerName, " ", groupName, " ", size, " ", modTime, " ", fileName)
}

func main() {
	// 실행 옵션 파싱
	allFlag := flag.Bool("a", false, "Include directory entries whose names begin with a dot (.).")
	longFlag := flag.Bool("l", false, "(The lowercase letter ``ell''.)  List in long format.")
	flag.Parse()
	dirs := flag.Args()
	// 하나도 없다면, 기본값으로 현재 실행 위치를 슬라이스에 추가
	if len(dirs) == 0 {
		dirs = append(dirs, ".")
	}

	for _, path := range dirs {
		results := []string{}
		files, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		// if *allFlag {
		// 	files = append([]string{".", ".."}, files...)
		// }

		for _, file := range files {
			result := ""
			info, err := file.Info()
			if err != nil {
				log.Fatal(err)
			}
			if *longFlag {
				result = applyLongFlag(&info)
			} else {
				result = info.Name()
			}
			results = append(results, result)
		}

		if *allFlag {
			results = applyAllFlag(results, *longFlag)
		}

		if *longFlag {
			fmt.Println(strings.Join(results, "\n"))
		} else {
			fmt.Println(strings.Join(results, " "))
		}
	}
}
