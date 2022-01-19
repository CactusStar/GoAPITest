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
	body, err := ioutil.ReadAll(resp.Body) 
	newbody := string(body)
	return str2map(newbody)
}

func getheader(url string){
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}

func verifyGetRequest(url string, expectResult map[string]string) {
	response := getRequest(url)
	for key, value := range expectResult {
		
		if _, ok := response[key]; ok {
			// should add assert here
			fmt.Println(value)
		}
	}
}

func postRequest() {

}

func verifyPostRequest() {

}

func str2map(cmd string) map[string]interface{}{
    str := strings.Replace(string(cmd), "'", "\"", -1)
    str = strings.Replace(str, "\n", "", -1)


    var dat map[string]interface{}
    if err := json.Unmarshal([]byte(str), &dat); err == nil {
        fmt.Println(dat)
		fmt.Printf("%T\n", dat)
		return dat
    } else {
        fmt.Println(err)
		return nil
    }
}

func analyzeResponse(string) {

}
// loop the method list and response list
// in a single list first get the key from the response
// Based on the method call the verify get or verify post
// the verify get or post contains searchkeys ==> this will generate a slice as the actual response
// the verify get or post will also contain a compare which compare the actual response with expected response
func testcaseExecution() {
	methodlist, responselist := readCase()
	for i, method := range methodlist {
		analyzeResponse(responselist[i])
		if method == "GET"{
			verifyGetRequest()
		} else if method == "POST" {
			verifyPostRequest()
		} else {
			fmt.Println("method incorrect")
		}
	}
}


// func main(){
// 	getheader("https://api.github.com/users/forAPItest")
// }