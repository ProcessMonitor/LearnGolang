package main

import "fmt"

//动态数组在传递参数上 是引用传递
func printArr(myArray []int) {
	//_ 表示匿名变量
	for _, value := range myArray {
		fmt.Print(" ", value)
	}
	myArray[0] = 1000
}

func main() {
	myArray := []int{1, 2, 3, 4} //动态数组 == 切片slice
	fmt.Println(myArray)
	//myArray2 := []int{1, 2, 3, 4}
	printArr(myArray)
	fmt.Print("\n",
		myArray)
}
