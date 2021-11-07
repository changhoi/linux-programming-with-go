package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Request struct {
	method string
	url    string
}

type RequestCountable struct {
	Request
	count int
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

// func countIP(ip string)

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

	// ipMap := make(map[string]Request)
	ipMap := make(map[string]int)
	reqMap := make(map[Request]int)
	requests := []RequestCountable{}
	scanner := bufio.NewScanner(logFile)

	for scanner.Scan() {
		line := scanner.Text()
		slice := strings.Split(line, "|")
		ip := findIP(slice[0])
		method, url := findAPI(slice[1])
		request := Request{method, url}

		// 문자열을 IP 타입으로 파싱하는 것
		trial := net.ParseIP(ip)
		if trial.To4() == nil {
			// IPv4 형태가 아니면 nil 리턴
			continue
		}

		// 1. 어떤 IP에서 몇 번의 요청을 보냈는지
		if _, ok := ipMap[ip]; ok {
			ipMap[ip] += 1
		} else {
			ipMap[ip] = 1
		}

		// 2. 어떤 메서드 + 엔드포인트 조합이 가장 많은 요청을 받았는지
		if _, ok := reqMap[request]; ok {
			reqMap[request] += 1
		} else {
			reqMap[request] = 1
		}
		for key, value := range reqMap {
			requests = append(requests, RequestCountable{Request{key.method, key.url}, value})
		}
		sort.Slice(requests, func(i, j int) bool { return requests[i].count > requests[j].count })

	}

	if *ipAddr != "" {
		fmt.Println(*ipAddr, ipMap[*ipAddr])
	}
	fmt.Println(requests[0].method, requests[0].url, requests[0].count)
}
