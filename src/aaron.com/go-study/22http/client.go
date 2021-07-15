package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	// get

	//resp, err := http.Get("127.0.0.1:600/get/")
	//if err != nil {
	//	fmt.Println("get url failed,err:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("read from resp.Body failed,err:", err)
	//	return
	//}
	//fmt.Print(string(body))
	getUrl := "http://127.0.0.1:600/get"
	//url parameter
	getData := url.Values{}
	getData.Set("name", "张三")
	getData.Set("age", "25")
	u, err := url.ParseRequestURI(getUrl)
	if err != nil {
		fmt.Println("parse url failed,err:", err)
		return
	}
	u.RawQuery = getData.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:", err)
		return
	}
	fmt.Println(string(b))

	// post
	// resp, err := http.PostForm("http://127.0.0.1:600/post", url.Values{"key": {"Value"}, "id": {"123"}})
	postUrl := "http://127.0.0.1:600/post"
	contentType := "application/json"
	data := `{"name":"枯藤","age":18}`
	resp1, err := http.Post(postUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err:", err)
		return
	}
	defer resp1.Body.Close()
	b1, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		fmt.Println("get resp failed,err:", err)
		return
	}
	fmt.Println(string(b1))
}
