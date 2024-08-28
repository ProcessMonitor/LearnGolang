package main

import (
	"fmt"
	"time"
)

func main() {
	//buffer channel cap = 3
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("channel data:", len(c), cap(c), c)
		}
	}()

	time.Sleep(1 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println(" main get cur_num:", num)
	}
	defer fmt.Println("main go程结束")

	close(c)
}
