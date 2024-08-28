package main

import (
	"fmt"
)

type ii int
type Book struct {
	Name      string
	Author    string
	Price     ii
	Sellcount int
}

func (this Book) GetInfo() string {
	return this.Name + " " + this.Author
}

func PrintBook(book Book) {
	fmt.Println("Book:", book)
}

func (this *Book) setInfo(book Book) {
	//this是调用的book对象的copy
	//this *是取地址
	this.Name = book.Name
	this.Author = book.Author
}

func main() {
	var a ii = 10
	fmt.Println("a : ", a)
	var book1 Book
	book1.Price = 11
	book1.Name = "golang"
	book1.Author = "zwq"
	book1.Sellcount = 10000

	fmt.Println("book1 info:", book1.GetInfo())

	var bk2 Book
	bk2.Name = "bk2"
	bk2.Author = "qwe"
	book1.setInfo(bk2)
	PrintBook(book1)
}
