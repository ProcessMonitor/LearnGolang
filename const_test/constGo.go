package main

import "fmt"

const (
	//iota  是一个表达式
	BEIJING = 10 * iota
	SHANGHAI
	SHENZHEN
)
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, k
	l, m
)

func main() {
	const length int = 123

	fmt.Println("beijing = ", BEIJING)
	fmt.Println("shanghai = ", SHANGHAI)
	fmt.Println("shenzhen = ", SHENZHEN)

	fmt.Println("ab cd ef = ", a, b, c, d, e, f)
	fmt.Println("gh ik lm = ", g, h, i, k, l, m)

}
