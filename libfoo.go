package main

// #include <Python.h>
// static PyObject* hoge(PyObject* )
import (
	"C"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Hoge struct {
	Foo  int64
	Bar  float64
	Baz  uint8
	Boo  uint32
	Data [64]byte
}

type Foo interface{}

type XXX func() string

//export Asdf
func Asdf() XXX {
	return Baaaaar
}

//export Qwer
func Qwer(x XXX) {
	fmt.Println(x())
}

//export FugaFuga
func FugaFuga() Foo {
	h := Hoge{Foo: 33, Bar: 44.333, Baz: 0xfe, Boo: 0x80000000}
	copy(h.Data[:], "Hello World!!!!!!!!!!!!")
	return h
}

//export HogeHoge
func HogeHoge(h Foo) {
	switch h := h.(type) {
	case Hoge:
		fmt.Println("is Hoge", h)
	default:
		fmt.Println("is not Hoge", h)
		return
	}
	// h := Hoge{Foo: 33, Bar: 44.333}
	// copy(h.Data[:], "Hello World!!!!!!!!!!!!")

	// pack
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, h)
	if err != nil {
		panic(err)
	}

	// print hex
	hextxt := hex.EncodeToString(buf.Bytes())
	fmt.Println(hextxt)
	border := []int{8, 8, 1, 4, -1}
	startAt := 0
	for _, v := range border {
		if v >= 0 {
			fmt.Println(hextxt[startAt*2 : (startAt+v)*2])
		} else {
			fmt.Println(hextxt[startAt*2:])
		}

		startAt += v
	}

	// unpack
	h2 := Hoge{}
	err = binary.Read(buf, binary.BigEndian, &h2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d %f %v %v %s\n", h2.Foo, h2.Bar, h2.Baz, h2.Boo, h2.Data)
	return
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

//export Baaaaar
func Baaaaar() string {
	return "Hello, Bar"
}

//export Bar2
func Bar2() (string, string) {
	return "Hello, Bar2", "bra bra bra"
}

//export Boo
func Boo(v string) {
	fmt.Println(v)
}

func main() {
	startAt := time.Now()
	HogeHoge(FugaFuga())
	fmt.Println(time.Now().Sub(startAt))
}
