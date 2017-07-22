package main

import (
	"C"
	"fmt"
	"io/ioutil"
	"net/http"
)

//export Fooooo
func Fooooo() int64 {
	resp, err := http.Get("http://www.yahoo.co.jp")
	if err != nil {
		return -1
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -2
	}
	fmt.Printf("%s\n", body)

	return 99990
}

func main() {
	fmt.Println("vim-go")
}
