package main

import "fmt"

func printArr(myArr [10]int) {
	fmt.Println("arr", myArr)
	//固定长度数组只有值拷贝不修改外面的myArr[0]
	myArr[0] = 1000

}
func main() {
	//定长数组
	array := [10]int{1, 2, 3, 4}

	printArr(array)
	// 不修改其myArr[0]

	for i := 0; i < 10; i++ {
		fmt.Print(array[i])
	}

	//range

	//for index, value := range array {
	//	fmt.Println("index :", index)
	//	fmt.Println("val :", value)
	//}
	//
	//for index, value := range array {
	//	fmt.Println("index : ", index, "val:", value)
	//}

}
