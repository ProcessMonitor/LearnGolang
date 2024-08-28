package main

import "fmt"

func main() {
	a := 10
	fmt.Println(&a)
	var p *int
	p = &a
	*p += 1
	fmt.Println(p)
}
