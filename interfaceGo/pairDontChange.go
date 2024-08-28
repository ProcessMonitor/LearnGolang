package main

import (
	"fmt"
	"io"
	os "os"
)

func main() {
	var allType interface{}

	str := "abc"

	allType = str
	//类型断言
	s, _ := allType.(string)
	fmt.Println("str :", s)
	//pair <type: os.File , value: /dev/tty*文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("nil :", nil)
		return
	}
	//pair <type: ,value: >
	var r io.Reader
	// r:pair <type: os.File, value: /dev/tty*文件描述符>
	r = tty
	//pair <type: , value: >
	var w io.Writer
	//pair <type: os.File , value: /dev/tty*文件描述符 //总之： 无论怎么赋值pair 结构不会改变
	w = r.(io.Writer)
	w.Write([]byte("hello os.writer"))
}
