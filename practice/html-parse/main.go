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

//

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
	SetRowValues(file, sheetName, 1, []string{"URL", "Class", "Role", "하위 요소", "텍스트"})

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
		"x_split": 0,
		"y_split": 1,
		"top_left_cell": "A2"
	}`)

	doc.Find("a").Each(func(i int, element *goquery.Selection) {
		index := i + 2
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
		SetRowValues(file, sheetName, index, []string{url, className, role, immediateChildType, text})
	})

	// excel로 저장
	file.SaveAs("parse-result.xlsx")
}
