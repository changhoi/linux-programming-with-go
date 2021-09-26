// https://m.blog.naver.com/PostView.naver?isHttpsRedirect=true&blogId=kwonsukmin&logNo=221238775732

package main

import (
	"flag"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/xuri/excelize/v2"
)

// 파싱 결과를 담는 구조체
type Result struct {
	Url                string // 링크 주소
	Class              string // HTML Class Attribute value
	Role               string // HTML Role value
	ImmediateChildType string // 자식 요소가 있을 경우, 해당 요소의 태그
	Text               string // a 태그의 텍스트 값
}

// 실행 시 인자로부터 파싱 대상 URl을 획득
// 미제공시 프로그램 종료
func InitTargetURLFromArg() string {
	targetURL := flag.String("target", "", "Target URL to parse information")
	flag.Parse()
	if flag.NFlag() == 0 {
		return ""
	}
	if !(strings.Contains(*targetURL, "http://") && strings.Contains(*targetURL, "https://")) {
		*targetURL = "http://" + *targetURL
	}

	return *targetURL
}

// 인자로 받은 행/열 정보 바탕으로
// 엑셀 형식에 맞게 새로운 셀을 생성하여 시트에 추가
func SetRowValues(file *excelize.File, sheet string, rowNumBase int, values []string) {
	colNumBase := 65 // 열 인덱스는 'A' 부터 시작
	rowNum := strconv.Itoa(rowNumBase)
	for index, element := range values {
		colNum := string(rune(colNumBase + index))
		file.SetCellValue(sheet, colNum+rowNum, element)
	}
}

// 머리행의 스타일 설정 - 배경 설정은 작동하지 않음
func SetHeaderRowStyle(file *excelize.File, sheet string) {
	style, err := file.NewStyle(`{
		"fill": {
			"type": "pattern",
			"color": ["#E0EBF5"],
			"pattern": 1
		}
	}`)
	if err != nil {
		file.SetCellStyle(sheet, "A1", "E1", style)
	}
	file.SetPanes(sheet, `{
		"freeze": true,
		"split": false,
		"x_split": 0,
		"y_split": 1,
		"top_left_cell": "A2"
	}`)
}

// 'a' 태그에 해당하는 요소를 파싱하고, URL 관련 정보를 파싱하여 반환
func Parse(document *goquery.Document) []Result {
	var results []Result

	document.Find("a").Each(func(i int, element *goquery.Selection) {
		immediateChildType := ""

		url := element.AttrOr("href", "")
		className := element.AttrOr("class", "")
		role := element.AttrOr("role", "")
		text := element.Text()

		if element.Children().Length() > 0 {
			element = element.Children()
			immediateChildType = goquery.NodeName(element)
		}

		results = append(results, Result{
			Url:                url,
			Class:              className,
			Role:               role,
			ImmediateChildType: immediateChildType,
			Text:               text,
		})
	})

	return results
}

func main() {
	targetURL := InitTargetURLFromArg()
	if targetURL == "" {
		flag.Usage()
	}

	resp, err := http.Get(targetURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	// Excel 파일 초기화
	file := excelize.NewFile()
	sheetName := "Sheet1"
	SetRowValues(file, sheetName, 1, []string{"URL", "Class", "Role", "하위 요소", "텍스트"})

	// 머리행 스타일 설정
	SetHeaderRowStyle(file, sheetName)

	// excel로 저장
	parsedData := Parse(doc)
	for i, result := range parsedData {
		index := i + 2
		SetRowValues(file, sheetName, index, []string{result.Url, result.Class, result.Role, result.ImmediateChildType, result.Text})
	}
	file.SaveAs("parse-result.xlsx")
}
