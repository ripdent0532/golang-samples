package main

import "runtime"

const NCPU = 4

func DoAll() {
	sem := make(chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		// 合理的缓冲区选项（个人理解就是和CPU的核心数相同）
		go DoPart(sem)
	}

	// 等待NCPU任务完成，释放通道sem的缓冲区
	for i := 0; i < NCPU; i++ {
		<-sem // 等待一个任务完成
	}

	// 全部完成
}

func DoPart(sem chan int) {
	// 进行计算的部分
	// ...
	sem <- 1 // 发送一个这部分已经完成的信号，用来释放sem的缓冲区
}

func main() {
	runtime.GOMAXPROCS(NCPU)

	DoAll()
}