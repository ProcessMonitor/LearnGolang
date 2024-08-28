package main

import "fmt"

var (
	dd **int
)

func main() {
	var (
		tmp *int
		q   *int
		p   *int
	)
	var a int = 10
	var b int = 20
	p = &a
	q = &b

	fmt.Println(" p , q", *p, *q)

	tmp = p
	p = q
	q = tmp

	fmt.Println("p , q", *p, *q)

	//二级指针
	var qq *int
	qq = &a
	dd = &qq

	fmt.Println("qq:", *qq)
	fmt.Println("dd:", **dd)

}
