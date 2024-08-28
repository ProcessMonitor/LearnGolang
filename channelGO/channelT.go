package main

import "fmt"

//main go
func main() {
	//定义一个channel
	c := make(chan int)
	i := 0
	//sub go
	go func() {
		defer fmt.Println("goroutine1 over")

		fmt.Println("goroutine1 running ")
		i = i + 10
		c <- i //666写入管道中

	}()
	go func() {
		defer fmt.Println("goroutine2 over")

		fmt.Println("goroutine2 running ")
		i = i + 2
		c <- i //666写入管道中

	}()

	num := <-c
	fmt.Println("num :", num)

	defer fmt.Println("main goroutine over")
}
