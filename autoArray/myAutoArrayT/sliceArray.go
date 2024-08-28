package main

import "fmt"

func main() {
	//声明slice 无对象 无元素空间
	var slice []int
	fmt.Println("about slice[]: ", len(slice), slice)
	//声明slice是切片 有对象 无元素空间
	slice1 := []int{}
	//判断一个slice 是否为空
	if slice == nil {
		fmt.Println("slice nil slice")
	} else {
		fmt.Println("slice not a nil slice")
	}
	if slice1 == nil {
		fmt.Println("slice1 nil slice")
	} else {
		fmt.Println("slice1 not a nil slice")
	}
	//打一下信息
	fmt.Println("about slice1[]: ", len(slice1), slice1)
	//开辟1个int大小空间
	slice1 = make([]int, 1)
	fmt.Println("new obj slice1[]: ", len(slice1), slice1)
	//第二种声明
	var slice2 []int = make([]int, 2)
	fmt.Println("about slice2[]: ", len(slice2), slice2)
	// 第三种 :=
	slice3 := make([]int, 3)
	fmt.Println("about slice3[]: ", len(slice3), slice3)

}
