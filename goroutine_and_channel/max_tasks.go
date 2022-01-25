package main

import "fmt"

const (
	AvailableMemory = 10 << 20 // 10M
	AverageMemoryPerRequest = 10 << 10 // 10K
	MAXREQS = AvailableMemory / AverageMemoryPerRequest
)

var sem = make(chan int, MAXREQS)

type Request struct {
	a, b int
	replyc chan int
}

func process(r *Request) {
	// Do something
	fmt.Println(r.a)
}

func handler(r *Request) {
	process(r)
	// 信号完成：开始启用下一个请求
	// 将sem的缓冲区释放一个位置
	<-sem
}

func Server(queue chan *Request) {
	for {
		sem <- 1
		// 当通道已满（1000个请求被激活）的时候将被阻塞
		// 所以停在这里等待，直到sem有容量（被释放），才能继续去处理请求
		request := <-queue
		go handler(request)
	}
}

func main() {
	queue := make(chan *Request)
	go Server(queue)
}


