package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
const (
	USERNAME = "root"
	PASSWORD = "StayReal1988@"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "cactus"
)

type testcase struct {
	TestCaseID string `db:"TestCaseID"`
	TestPurpose string `db:"TestPurpose"`
	URL string `db:"URL"`
	Run string `db:"Run"`
	Response string `db:"Response"`
	Method string `db:"Method"`
}
// var DB *sql.DB
// var err error
func readCase() ([]string, []string, []string) {
	fmt.Println("=====================start DB operation=======================")
	// var DB *sql.DB
	// var err error
	var methodlist []string
	var responselist []string
	var urllist []string
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return nil, nil, nil
	}
	user := new(testcase)
	// var rows *sql.Rows
    rows, _ := DB.Query("select Response, Method, URL from cactus.testcase where Run=\"Test\"")

	fmt.Println("==query finished prepare to loop==")
	for rows.Next() {
		fmt.Println("==start loop==")
		if err := rows.Scan(&user.Response, &user.Method, &user.URL); err != nil{
			fmt.Printf("scan failed, err:%v",err)
			return nil, nil, nil
		}
		// fmt.Println(user.Response, user.Method)
		methodlist = append(methodlist, user.Method)
		responselist = append(responselist, user.Response)
		urllist = append(urllist, user.URL)
	}
	fmt.Println(methodlist)
	fmt.Println(responselist)
	return methodlist, responselist, urllist
}

// func main() {
// 	readCase()
// }
