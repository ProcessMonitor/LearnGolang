package main

import "fmt"

func test(a string, b int) int {
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	c := 11
	return c
}

//返回多个返回值 匿名形参
func test1(a string, b int) (int, int) {
	fmt.Println("method name =", a)
	fmt.Println("arg int =", b)
	return 22, 33
}

//返回多个返回值 有形参名称
//或者可以写 func test2(a string, b int) (r1 , r2 int)
func test2(a string, b int) (r1 int, r2 int) {
	fmt.Println("method name =", a)
	fmt.Println("arg int =", b)
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)
	// 返回值赋值
	r1 = 44
	r2 = 55
	return
}
func main() {
	c := test("123", 512)
	fmt.Println("c= ", c)

	ret1, ret2 := test1("test1", 1024)
	fmt.Println("ret1 , ret2 :", ret1, ret2)

	ret3, ret4 := test2("test2", 2048)
	fmt.Println("ret3 , ret4 :", ret3, ret4)
}
