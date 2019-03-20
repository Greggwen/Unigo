package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {

	var requestURL string
	requestURL = "https://www.evente.cn"

	//http.SetCookie("adsfasdf", "")

	res, _ := http.Get(requestURL)
	defer res.Body.Close()

	//res.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//res.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
	//res.Header.Add("Connection", "keep-alive")
	//res.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
