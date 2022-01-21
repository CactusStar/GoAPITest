package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"encoding/json"
	"strconv"
)


func getRequest(url string) (map[string]interface{}, *http.Response) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) 
	newbody := string(body)
	return str2map(newbody), resp
}

func getheader(url string){
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Header)
}

func postRequest() {
	resp, err := http.Post()
}
func verifyRequest() {
	
	_, erresponses, urls := readCase()
	for i, erresponse := range erresponses {
		var actualR string
		response, resp := getRequest(urls[i])
		keys, finalERresp := analyzeResponse(erresponse)
		for _, key := range keys{
			var stringvalue string
			value := response[key]
			switch value.(type) {
				case float64:
					ft := value.(float64)
					stringvalue = strconv.FormatFloat(ft, 'f', -1, 64)
				case float32:
					ft := value.(float32)
					stringvalue = strconv.FormatFloat(float64(ft), 'f', -1, 64)
				case int:
					it := value.(int)
					stringvalue = strconv.Itoa(it)
				case uint:
					it := value.(uint)
					stringvalue = strconv.Itoa(int(it))
				case int8:
					it := value.(int8)
					stringvalue = strconv.Itoa(int(it))
				case uint8:
					it := value.(uint8)
					stringvalue = strconv.Itoa(int(it))
				case int16:
					it := value.(int16)
					stringvalue = strconv.Itoa(int(it))
				case uint16:
					it := value.(uint16)
					stringvalue = strconv.Itoa(int(it))
				case int32:
					it := value.(int32)
					stringvalue = strconv.Itoa(int(it))
				case uint32:
					it := value.(uint32)
					stringvalue = strconv.Itoa(int(it))
				case int64:
					it := value.(int64)
					stringvalue = strconv.FormatInt(it, 10)
				case uint64:
					it := value.(uint64)
					stringvalue = strconv.FormatUint(it, 10)
				case string:
					stringvalue = value.(string)
				case []byte:
					stringvalue = string(value.([]byte))
				default:
					newValue, _ := json.Marshal(value)
					stringvalue = string(newValue)
			}
			if key == "Status" {
				stringvalue = resp.Status
			} else if key == "StatusCode" {
				stringvalue = strconv.Itoa(resp.StatusCode)
				fmt.Println("=-==statuscode====")
				fmt.Println(stringvalue)
			} else if key == "Proto" {
				stringvalue = resp.Proto
			} else if key == "ProtoMajor" {
				stringvalue = strconv.Itoa(resp.ProtoMajor)
			} else if key == "ProtoMinor" {
				stringvalue = strconv.Itoa(resp.ProtoMinor)
			}
			actualR = actualR + key + ":" + stringvalue + ", "
		}
		if finalERresp == actualR {
			fmt.Println("passed")
		} else {
			fmt.Println("===ER===")
			fmt.Println(finalERresp)
			fmt.Println("===AR===")
			fmt.Println(actualR)
			fmt.Println("failed")
		}
	}
	
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

func analyzeResponse(erresp string) ([]string, string){
	var key []string
	temp_resp := erresp[1:]
	final_resp := temp_resp[0:len(temp_resp)-1]
	test := strings.Split(final_resp, ", ")
	for _, item := range test {
		key_temp := strings.Split(item, ":")
		if key_temp[0] != "" {
			key = append(key, key_temp[0])
		}
	}
	fmt.Println(key)
	fmt.Println(final_resp)
	return key, final_resp
}
// loop the method list and response list
// in a single list first get the key from the response
// Based on the method call the verify get or verify post
// the verify get or post contains searchkeys ==> this will generate a slice as the actual response
// the verify get or post will also contain a compare which compare the actual response with expected response
// func testcaseExecution() {
// 	methodlist, responselist := readCase()
// 	for i, method := range methodlist {
// 		currentResponse := analyzeResponse(responselist[i])
// 		if method == "GET"{
// 			verifyGetRequest()
// 		} else if method == "POST" {
			
// 		} else {
// 			fmt.Println("method incorrect")
// 		}
// 	}
// }

func main() {
	// verifyGetRequest()
	getheader("https://api.github.com/users/forAPItest")
}


// func main(){
// 	getheader("https://api.github.com/users/forAPItest")
// }