package main

import "fmt"

func fib(c, quit chan int) {
	x, y := 1, 1
	for {
		// select 多路监控channel状态
		select {
		case c <- x:
			//c可写就走该case
			x = y
			y = x + y
		case <-quit:
			fmt.Println("return for quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// sub go
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		//用叫做quit的channel 控制程序退出
		quit <- 0
	}()
	fib(c, quit)
}
