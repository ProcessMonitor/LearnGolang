package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 用go创建一个 形参为空，返回值为空的匿名函数 并执行
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//退出当前方法
			//return
			//退出当前goroutine
			runtime.Goexit()
			fmt.Println("B func")
		}()
		fmt.Println("A func")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
