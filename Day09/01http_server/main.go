package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	str := "AAA"
	w.Write([]byte(str))
}

func f2(w http.ResponseWriter, r *http.Request) {
	//GET 参数是在请求连接上的， query parm
	fmt.Println(r.URL.Query())
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/f1", f1)
	http.HandleFunc("/f2", f2)
	http.ListenAndServe("127.0.0.1:8080", nil)

}
