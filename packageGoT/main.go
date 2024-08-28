package main

import (
	"fmt"
	// 匿名方式导包 （只调用lib2的init方法）
	_ "LearnGoProject/packageGoT/lib2"
	// 起别名
	l3 "LearnGoProject/packageGoT/lib3"
	// . 相当于把文件所有方法都移到当前文件中 可以直接函数名调用
	. "LearnGoProject/packageGoT/lib4"
)

func main() {
	fmt.Println("main running ..")
	l3.Test()
	L4test()
}
