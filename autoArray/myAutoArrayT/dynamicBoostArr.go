package main

import (
	"fmt"
)

func main() {
	//slice是根据初始化的容量每次自动扩容一倍 比如numbers就是从 5 -》10 -》15
	var numbers = make([]int, 3, 5)
	fmt.Println("len cap slice ", len(numbers), cap(numbers), numbers)
	//fmt.Println(binary.Size(numbers))
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	fmt.Println("len cap slice ", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 4)
	fmt.Println("len cap slice ", len(numbers), cap(numbers), numbers)

	//切片 [3,6)
	slice := numbers[3:6]

	fmt.Println("slice : ", slice)

	slice2 := make([]int, 5)
	copy(slice2, slice)
	fmt.Println("slice 2", slice2)
}
