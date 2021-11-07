package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

type Request struct {
	count  int
	method string
	url    string
}

func findIP(input string) string {
	/*
		IP가 255가 최대이기 때문에 한 영역에 대한 정규표현식은 다음과 같다.
	*/
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammer := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammer)
	return matchMe.FindString(input)
}

func findAPI(input string) (string, string) {
	method := "(GET|POST|HEAD|OPTIONS|DELETE)"
	url := `((\/([.a-zA-Z0-9])*)+)`
	params := `(\?([a-zA-Z0-9])*)*`

	methodMatch := regexp.MustCompile(method)
	urlMatch := regexp.MustCompile(url + params)

	return methodMatch.FindString(input), urlMatch.FindString(input)
}

// go run regex.go -address="172.104.131.24" access.log
func main() {
	ipAddr := flag.String("address", "", "IP Address")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		os.Exit(1)
	}
	path := args[0]

	logFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	ipMap := make(map[string]Request)
	scanner := bufio.NewScanner(logFile)

	for scanner.Scan() {
		line := scanner.Text()
		slice := strings.Split(line, "|")
		ip := findIP(slice[0])
		method, url := findAPI(slice[1])

		// 문자열을 IP 타입으로 파싱하는 것
		trial := net.ParseIP(ip)
		if trial.To4() == nil {
			// IPv4 형태가 아니면 nil 리턴
			continue
		}

		if entry, ok := ipMap[ip]; ok {
			entry.count += 1
			ipMap[ip] = entry
		} else {
			newEntry := Request{1, method, url}
			ipMap[ip] = newEntry
		}
	}

	if *ipAddr != "" {
		fmt.Println(*ipAddr, ipMap[*ipAddr])
	}
}
