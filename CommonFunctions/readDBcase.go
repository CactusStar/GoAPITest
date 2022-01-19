package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
const (
	USERNAME = "root"
	PASSWORD = "xxxxx"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "cactus"
)

type testcase struct {
	TestCaseID string `db:"TestCaseID"`
	TestPurpose string `db:"TestPurpose"`
	URL sql.NullString `db:"URL"`
	Run string `db:"Run"`
	Response string `db:"Response"`
	Method string `db:"Method"`
}

func readCase() ([]string, []string) {
	var methodlist []string
	var responselist []string
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return nil, nil
	}
	user := new(testcase)
    rows,_ := DB.Query("select Response, Method from cactus.testcase where Run=\"Yes\"")

	for rows.Next() {
		if err := rows.Scan(&user.Response, &user.Method); err != nil{
			fmt.Printf("scan failed, err:%v",err)
			return nil, nil
		}
		// fmt.Println(user.Response, user.Method)
		methodlist = append(methodlist, user.Method)
		responselist = append(responselist, user.Response)
	}
	fmt.Println(methodlist)
	fmt.Println(responselist)
	return methodlist, responselist
}

// func main() {
// 	readCase()
// }
