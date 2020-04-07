package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8080/f2?page=15&user=1&name=比比大神")
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	//
	//var data []byte
	//resp.Body.Read(data)
	//resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("ERROR : ", err)
		return
	}
	//字节转字符串
	fmt.Println(string(b))

	//使用 encode
	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:8080/f1")
	data.Set("name", "比比大神")
	data.Set("age", "29")
	queryStr := data.Encode()
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, _ := http.NewRequest("GET", urlObj.String(), nil)
	resq, _ := http.DefaultClient.Do(req)
	defer resq.Body.Close()
	fmt.Println(resq.Body)
	//禁用长连接
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Print("ERROR : ", err)
		return
	}
}
