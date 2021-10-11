package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
)

/*
 -u 옵션이 기본임
 -g 프라이머리 그룹만
 -G 그룹 전체

 유저 이름이나, 아이디 값으로 서치할 수 있음. 없으면 현재 로그인 된 유저 (프로세스 실행 중인 유저)
*/

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getGroupsString(gids []int) string {
	return strings.Trim(fmt.Sprint(gids), "[]")
}

func main() {
	minusAG := flag.Bool("G", false, "Groups")
	minusG := flag.Bool("g", false, "Primary group")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {

		if *minusAG {
			groups, err := os.Getgroups()
			checkError(err)
			fmt.Println(getGroupsString(groups))
			return
		}

		if *minusG {
			fmt.Println(os.Getegid())
			return
		}
		id := os.Geteuid()
		fmt.Println(id)
		return
	}
	userString := args[0]

	var lookup func(u string) (*user.User, error)

	if _, err := strconv.ParseInt(userString, 10, 32); err == nil {
		// ID값으로 서치 하는 중
		lookup = user.LookupId
	} else {
		// 유저 이름으로 찾기
		lookup = user.Lookup
	}

	u, err := lookup(userString)
	checkError(err)

	if *minusAG {
		gids, err := u.GroupIds()
		checkError(err)
		fmt.Println(strings.Join(gids, " "))

		return
	}

	if *minusG {
		fmt.Println(u.Gid)
		return
	}

	fmt.Println(u.Uid)
	return
}
