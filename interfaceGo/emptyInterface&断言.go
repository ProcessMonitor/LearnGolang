package main

import "fmt"

//万能类型 interface 类似 T object
func print(arg interface{}) {
	fmt.Println("arg :", arg)
	// arg 如何区别底层的类型？ "类型断言":
	value, ok := arg.(string)
	if !ok {
		fmt.Println("not string")
	} else {
		fmt.Println("is str : ", value)
	}
}

func main() {
	a := 1
	b := "str"
	c := 0.5
	print(a)
	print(b)
	print(c)
}
