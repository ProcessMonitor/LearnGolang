package main

import (
	"fmt"
)

func main() {
	m := map[string]int{}
	var mm map[string]int
	if m == nil {
		fmt.Println(" m nil map ")
	} else {
		fmt.Println("m not a nil map")
	}
	if mm == nil {
		fmt.Println("mm nil map ")
	} else {
		fmt.Println("mm not a nil map")
	}

	fmt.Println(" m map :", len(m), m)
	fmt.Println(" mm map size:", len(mm), "data: ", mm)
	m = make(map[string]int, 5)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	//底层是hash table 所以是乱序的
	fmt.Println(" m map size:", len(m), "data: ", m)

	//hash := maphash.Hash{}
	//byte := byte(111)
	//fmt.Println(hash.Size(), hash.BlockSize(), hash.WriteByte(byte))
	//fmt.Println(hash)
	fmt.Println("map[three]: ", m["three"])

	//1直接声明
	var myMap1 map[string]int
	myMap1 = make(map[string]int, 10)
	//2直接开辟map空间
	myMap2 := make(map[string]int, 10)
	myMap2["number1"] = 1
	myMap2["number2"] = 2
	//3打括号初始化
	myMap3 := map[string]int{
		"name": 10010,
		"age":  3,
		"tag":  4,
	}
	fmt.Println("my map1:", myMap1, " my map2:", myMap2, "my map3:", myMap3)

}
