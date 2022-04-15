/*
 * @title:http协议 server端
 * @Descripttion:
 * @version:
 * @Author: Guoyh
 * @Date: 2022-03-29 16:18:19
 * @LastEditors: Guoyh
 * @LastEditTime: 2022-03-30 18:15:15
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f(res http.ResponseWriter, req *http.Request) {
	str, err := ioutil.ReadFile("hello.html")
	if err != nil {
		// open xxx.txt: The system cannot find the file specified.
		res.Write([]byte(fmt.Sprintln(err)))
	}
	res.Write([]byte(str))
}

func main() {
	http.HandleFunc("/opt/lululu/Go", f)
	http.ListenAndServe("127.0.0.1:8989", nil)
}
