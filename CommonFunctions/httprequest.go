package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"encoding/json"
)


func getRequest(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //读取body的内容

	newbody := string(body)
	return str2map(newbody)
}

func verifyGetRequest(url string, expectResult map[string]string) {
	
}

func postRequest() {

}

func str2map(cmd string)map[string]interface{}{
    str := strings.Replace(string(cmd), "'", "\"", -1)
    str = strings.Replace(str, "\n", "", -1)


    var dat map[string]interface{}
    if err := json.Unmarshal([]byte(str), &dat); err == nil {
        fmt.Println(dat)
		fmt.Printf("%T\n", dat)
		return dat
        //fmt.Println(dat["status"])
    } else {
        fmt.Println(err)
		return nil
    }
}


func main(){
	getRequest("https://api.github.com/users/forAPItest")

}

// 