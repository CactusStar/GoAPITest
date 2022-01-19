package main

import (
	"fmt"
)

func main(){
	fmt.Println("aaa")
	var test []string
	// test := []string{"followers": "abc", "following": "def"} {"followers": "sss", "following": "ttt"}
	// test := []string{"aaa","bbb"}
	test = []string{"{'abc': 'eee'}", "{'fff':'ggg'}"}
	fmt.Println(test)
}