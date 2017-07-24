package main

import (
	"C"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Hoge struct {
	Foo  int64
	Bar  float64
	Data [256]byte
}

//export HogeHoge
func HogeHoge() {
	h := Hoge{Foo: 33, Bar: 44.333}
	copy(h.Data[:], "Hello World")

	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, h)
	if err != nil {
		panic(err)
	}

	h2 := Hoge{}
	err = binary.Read(buf, binary.BigEndian, &h2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d %f %s\n", h2.Foo, h2.Bar, h2.Data)
}

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
	HogeHoge()
}
