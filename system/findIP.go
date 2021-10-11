package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIP(input string) string {
	/*
		IP가 255가 최대이기 때문에 한 영역에 대한 정규표현식은 다음과 같다.
	*/
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammer := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammer)
	return matchMe.FindString(input)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s logFile\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer f.Close()

	ipMap := make(map[string]int)
	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()
		ip := findIP(line)

		// 문자열을 IP 타입으로 파싱하는 것
		trial := net.ParseIP(ip)
		if trial.To4() == nil {
			// IPv4 형태가 아니면 nil 리턴
			continue
		}

		if _, ok := ipMap[ip]; ok {
			ipMap[ip]++
		} else {
			ipMap[ip] = 1
		}
	}

	for key := range ipMap {
		fmt.Printf("%s %d\n", key, ipMap[key])
	}
}
