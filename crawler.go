package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://www.baidu.com")
	fmt.Println(resp)
}
