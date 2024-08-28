package main

import "fmt"

type My_Reader interface {
	ReadBook()
}

type My_Writter interface {
	WriteBook()
}

//具体类型
type Book struct {
	BookName string
}

func (this *Book) ReadBook() {
	fmt.Println("read a book")
}
func (this *Book) WriteBook() {
	fmt.Println("Write a book")
}

func main() {
	// b:pair : <type:book,value: Book{}指针>
	b := &Book{}
	// r:pair : <type:,value: >
	var r My_Reader
	// r:pair : <type:book,value: Book{}指针>
	r = b
	r.ReadBook()
	var w My_Writter
	w = r.(My_Writter) //此处的断言为什么会成功呢 ？ 以为w/r 具体的type是一致的
	w.WriteBook()
}
