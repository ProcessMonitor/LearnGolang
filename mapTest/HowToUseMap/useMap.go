package main

import "fmt"

//引用传递
func printMap(myMap map[string]int) {
	for key, val := range myMap {
		fmt.Println("key :", key, "value ", val)
	}
}

func main() {
	myMap := make(map[string]int)
	//add
	myMap["china"] = 1
	myMap["usa"] = 2
	myMap["uk"] = 3
	//遍历
	printMap(myMap)
	//delete
	delete(myMap, "usa")
	//update
	myMap["china"] = 0
	fmt.Println("-----------")
	//遍历
	printMap(myMap)

}
