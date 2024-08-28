package main

import (
	"fmt"
	"time"
)

//从goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Println("new Goroutine task", i)
		time.Sleep(1 * time.Second)

	}

}

//主goroutine
func main() {
	// 创建go线程
	go newTask()

	i := 0
	for {
		i++
		fmt.Println("main goroutine ", i)

		time.Sleep(1 * time.Second)
	}

}
