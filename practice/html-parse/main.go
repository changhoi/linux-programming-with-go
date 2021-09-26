// https://m.blog.naver.com/PostView.naver?isHttpsRedirect=true&blogId=kwonsukmin&logNo=221238775732

package main

import (
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/xuri/excelize/v2"
)

// a 태그 통하여 사용된 것들을 쭉 모아서
// role, class 등을 표시하고
// 하위의 태그들을 정리
// 정리하여 엑셀로 뽑기 -> github.com/xuri/excelize/v2
type Result struct {
	Url                string
	Class              string
	Role               string
	HasChildren        bool
	ImmediateChildType string
}

var ResultList []Result

func main() {
	resp, err := http.Get("https://naver.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	// Excel 파일 초기화
	file := excelize.NewFile()
	sheetName := "Sheet1"
	file.SetCellValue(sheetName, "A1", "URL")
	file.SetCellValue(sheetName, "B1", "Class")
	file.SetCellValue(sheetName, "C1", "Role")
	file.SetCellValue(sheetName, "D1", "하위 요소")
	file.SetCellValue(sheetName, "E1", "텍스트")

	// 머리행 스타일 설정
	style, err := file.NewStyle(`{
		"fill": {
			"type": "pattern",
			"color": ["#E0EBF5"],
			"pattern": 1
		}
	}`)
	if err != nil {
		file.SetCellStyle(sheetName, "A1", "A1", style)
		file.SetCellStyle(sheetName, "B1", "B1", style)
		file.SetCellStyle(sheetName, "C1", "C1", style)
		file.SetCellStyle(sheetName, "D1", "D1", style)
		file.SetCellStyle(sheetName, "E1", "E1", style)
	}
	file.SetPanes(sheetName, `{
		"freeze": true,
		"split": false,
		"top_left_cell": "A2"
	}`)

	// title element
	// title := doc.Find("title").Text()
	// doc.Find("a").Each(func(i int, s *goquery.Selection) {
	// 	if i < 150 && i >= 70 {
	// 		child := s.Children()
	// 		if child.Length() > 0 {
	// 			fmt.Println("[CHILD]", goquery.NodeName(child), child.Text())
	// 		} else {
	// 			fmt.Println(s.Text())
	// 		}
	// 	}
	// })

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		index := strconv.Itoa(i + 2)
		element := s
		immediateChildType := ""

		url := element.AttrOr("href", "")
		className := element.AttrOr("class", "")
		role := element.AttrOr("role", "")
		text := element.Text()

		if element.Children().Length() > 0 {
			element = element.Children()
			immediateChildType = goquery.NodeName(element)
		}

		// excel로 뽑아내기
		file.SetCellValue(sheetName, "A"+index, url)
		file.SetCellValue(sheetName, "B"+index, className)
		file.SetCellValue(sheetName, "C"+index, role)
		file.SetCellValue(sheetName, "D"+index, immediateChildType)
		file.SetCellValue(sheetName, "E"+index, text)
	})

	// excel로 저장
	file.SaveAs("parse-result.xlsx")
}
