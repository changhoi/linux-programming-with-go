package main

import (
	"fmt"
)

type DBPerformer interface {
	create(string, string) (bool, string)
	read(string) (bool, string)
	update(string, string) (bool, string)
	delete(string) bool
}

type StringDB struct {
	db map[string]string
}

func (database StringDB) create(k string, v string) (bool, string) {
	if _, exist := database.db[k]; exist {
		return false, ""
	} else {
		database.db[k] = v
		return true, database.db[k]
	}
}

func (database StringDB) read(k string) (bool, string) {
	if v, exist := database.db[k]; exist {
		return true, v
	} else {

		return false, ""
	}
}

func (database StringDB) update(k string, v string) (bool, string) {
	if _, exist := database.db[k]; exist {
		database.db[k] = v
		return true, database.db[k]
	} else {
		return false, ""
	}
}

func (database StringDB) delete(k string) bool {
	if _, exist := database.db[k]; exist {
		delete(database.db, k)
		return true
	} else {
		return false
	}
}

func newStringDB() *StringDB {
	d := StringDB{}
	d.db = make(map[string]string)
	return &d
}

type Query struct {
	Op    string
	Key   string
	Value string
}

func (query *Query) printQuery() string {
	return query.Op + " " + query.Key + " " + query.Value
}

func main() {
	var db DBPerformer = newStringDB()
	// 쿼리가 들어온다고 가정
	Queries := make([]Query, 0)

	Queries = append(Queries, Query{"c", "Go", "is easy"})
	Queries = append(Queries, Query{"r", "Go", ""})
	Queries = append(Queries, Query{"c", "Hello", "World"})
	Queries = append(Queries, Query{"u", "Hello", "Go"})
	Queries = append(Queries, Query{"r", "Hello", ""})
	Queries = append(Queries, Query{"d", "Hello", ""})
	Queries = append(Queries, Query{"u", "Hello", "error"})

	for _, query := range Queries {
		fmt.Println("-------")
		fmt.Println(query.printQuery())
		switch query.Op {
		case "c":
			if ok, v := db.create(query.Key, query.Value); ok {
				fmt.Println(v + " 입력 성공")
			} else {
				fmt.Println("이미 존재하는 키 입니다")
			}

		case "r":
			if ok, v := db.read(query.Key); ok {
				fmt.Println(query.Key + " : " + v)
			} else {
				fmt.Println("존재하지 않는 키 입니다.")
			}
		case "u":
			if ok, v := db.update(query.Key, query.Value); ok {
				fmt.Println(query.Key + " : " + v + " updated")
			} else {
				fmt.Println("존재하지 않는 키 입니다.")
			}

		case "d":

			if ok := db.delete(query.Key); ok {
				fmt.Println(query.Key + " 삭제 성공")
			} else {
				fmt.Println("존재하지 않는 키 입니다.")
			}
		}
	}
}
