package main

import "fmt"

func main() {
	//不初始化的channel 无法收发数据
	c := make(chan int, 5)

	go func() {
		for i := 0; i < 3; i++ {
			c <- i
		}
		close(c)
		defer fmt.Println(" son over")
	}()

	//for {
	//	//ok == ture 则没有关闭
	//	if data, ok := <-c; ok {
	//		fmt.Println("data:", data)
	//	} else {
	//		break
	//	}
	//}

	//用range自动迭代访问channel
	for data := range c {
		fmt.Println("data", data)
	}

	fmt.Println("main over")
}
