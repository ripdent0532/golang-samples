package main

import (
	"fmt"
	"time"
)

// 通道工厂模式
// 不将通道作为参数传递给协程，而用函数来生成一个通道并返回（工厂角色）
// 函数内有个匿名函数被协程调用
func main() {
	stream := pump3()
	go suck3(stream)
	time.Sleep(1e9)
}

func pump3() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck3(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}