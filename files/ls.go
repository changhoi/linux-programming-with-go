package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func getCurrentWorkingDirectoryRel() string {
	ex, err := os.Executable()
	handleError(err)

	curdir := filepath.Dir(ex)

	dir, err := filepath.Rel(curdir, curdir)
	handleError(err)

	return dir

}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func isStartWithDot(path string) bool {
	if path[0:1] == "." {
		return true
	}
	return false
}

func getContentInfo(minusL *bool, fi fs.FileInfo) string {
	if *minusL {
		ret := make([]string, 0)
		// 파일 mode
		ret = append(ret, fi.Mode().String())
		// 파일의 하드링크 개수
		ret = append(ret, strconv.FormatInt(int64(countHardLink(fi)), 10))
		// 파일의 유저이름, 그룹이름
		ret = append(ret, getUserName(fi), getGroupName(fi))
		// 파일의 사이즈
		ret = append(ret, strconv.FormatInt(int64(fi.Size()), 10))
		// 마지막 수정된 시각 -> 월 일 시간(24H)
		mt := fi.ModTime()
		ret = append(ret, formatTime(mt.Month(), mt.Day(), mt.Hour(), mt.Minute()))
		// 파일의 이름
		ret = append(ret, fi.Name())

		return strings.Join(ret, "\t")
	} else {
		return fi.Name()
	}
}

// fileInfo로 넘겨온 파일의 hard link의 개수를 세서 반환
func countHardLink(fi fs.FileInfo) uint64 {
	nlink := uint64(0)

	if sys := fi.Sys(); sys != nil {
		if stat, ok := sys.(*syscall.Stat_t); ok {
			nlink = uint64(stat.Nlink)
		}
	}
	return nlink
}

func getUserName(fi fs.FileInfo) string {
	// 파일의 uid 가져옴
	uid := fi.Sys().(*syscall.Stat_t).Uid
	// 파일의 uid를 string으로 변환
	uidStr := strconv.FormatUint(uint64(uid), 10)

	user, err := user.LookupId(uidStr)
	handleError(err)
	return user.Username
}

func getGroupName(fi fs.FileInfo) string {
	// 파일의 gid 가져옴
	gid := fi.Sys().(*syscall.Stat_t).Gid
	// 파일의 gid를 string으로 변환
	gidStr := strconv.FormatUint(uint64(gid), 10)

	group, err := user.LookupGroupId(gidStr)
	handleError(err)
	return group.Name
}

func formatTime(m time.Month, d int, h int, min int) string {
	ret := make([]string, 0)

	ret = append(ret, m.String())
	ret = append(ret, strconv.Itoa(d))
	ret = append(ret, strconv.Itoa(h)+":"+strconv.Itoa(min))

	return strings.Join(ret, " ")
}

func main() {
	minusL := flag.Bool("l", false, "Listing format")
	minusA := flag.Bool("a", false, "All")

	flag.Parse()

	dir := getCurrentWorkingDirectoryRel()

	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		// fmt.Println(dir)

		fi, err := os.Stat(path)

		// .으로 시작하고 정규파일일 경우에 : 숨겨진 파일
		if isStartWithDot(path) && !fi.IsDir() {
			// 숨겨진 파일은 all 옵션이 켜져있을때만 출력
			if *minusA {
				fmt.Println(getContentInfo(minusL, fi))
			}
		} else {
			// 나머지의 경우 : 정규 파일의 경우나 상위 디렉토리
			fmt.Println(getContentInfo(minusL, fi))
		}

		return nil
	})
}
