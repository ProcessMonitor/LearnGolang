package main

import "fmt"

func f1() {
	fmt.Println("f1")
}
func f2() {
	fmt.Println("f2")
}
func f3() {
	fmt.Println("f3")
}

func return_defer_T() int {
	defer defer_call()
	return return_call()
}

func defer_call() int {
	fmt.Println("defer func call")
	return 0
}
func return_call() int {
	fmt.Println("return func call")
	return 1
}

func main() {
	//defer 关键字修饰的语句 最后执行 （有多条语句入栈就先进后出）
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")
	//defer 函数也一样后入先出
	defer f1()
	defer f2()
	defer f3()

	fmt.Println("main body output1 ")
	fmt.Println("main body output2 ")
	//defer and return 先后顺序问题
	fmt.Println("Who is first?",
		return_defer_T())
	//验证后是 return先执行  defer后执行
}
