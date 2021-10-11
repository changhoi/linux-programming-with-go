package main

import (
	"fmt"
	"time"
)

func main() {
	var zeroTime time.Time
	fmt.Println("Zero:", zeroTime)
	fmt.Println()

	current := time.Now() // 현재 시각 기준으로 Time 타입 만들기

	fmt.Println("time.Now")
	fmt.Println("Default:", current)
	fmt.Println("Epoch time:", current.Unix()) // Epoch time seconds
	fmt.Println("Format:", current.Format("01 02 03 04 05 06 MST"))
	//	Jan 2 15:04:05 2006 MST
	//	1	2	3:4:5	6	7
	fmt.Println()

	date := time.Date(2021, time.October, 10, 10, 0, 0, 0, time.UTC)
	fmt.Println("time.Date")
	fmt.Println("Default:", date)
	fmt.Println("Local:", date.Local()) // 현재 지역 기준으로
	fmt.Println("Format:", date.Format("2006/Jan/02"))

	fmt.Println()
	fmt.Println("time.Parse")
	parsed, _ := time.Parse("02 January 2006", "10 October 2021")
	fmt.Println("Default:", parsed)
	fmt.Println("Unix epoch nsec:", parsed.UnixNano())
	fmt.Println("Preset format:", parsed.Format(time.RFC3339))

	fmt.Println()
	fmt.Println("time.Unix")
	fmt.Println("Default:", time.Unix(0, 0))
}
