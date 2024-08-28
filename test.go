package main

import "fmt"

// 求元素和
func sumArr(a []int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func main() {
	//fmt.Println("hello GO! ")
	//var a interface{}
	//
	//value, ok := a.(string)
	//if !ok {
	//	fmt.Println("It's not ok for type string")
	//	return
	//}
	//
	//fmt.Println("The value is ", value)

	// 若想做一个真正的随机数，要种子
	// seed()种子默认是1
	//rand.Seed(1)
	//rand.Seed(time.Now().Unix())
	//
	//var b [10]int
	//for i := 0; i < len(b); i++ {
	//	// 产生一个0到1000随机数
	//	b[i] = rand.Intn(5)
	//	println(b[i])
	//}
	//sum := sumArr(b)
	//fmt.Printf("sum=%d\n", sum)
	mymap := map[int]string{
		11: "11",
		22: "22",
	}
	for i, s := range mymap {
		fmt.Println("key : ", i, "val : ", s)
	}

	numbers := make([]int, 3, 10)
	fmt.Printf(" value : %v\n , cap = %d \n", (numbers), cap(numbers))
	//apArray := []int{1, 2, 3, 4}
	//value := 0
	numbers = append(numbers, 4)

	fmt.Printf(" value : %v\n , cap = %d \n", (numbers), cap(numbers))

}
