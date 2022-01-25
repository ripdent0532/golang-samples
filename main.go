// @Title main
// @Description 主程序包
// @Author Jianghui Guo ${DATE} ${TIME}
package main

import (
	"fmt"
)

type Task struct {
	id int
}

func main() {
	pending, done := make(chan *Task), make(chan *Task)
	go sendWork(pending) // put tasks with work on the channel
	for i := 0; i < 2; i++ {
		go Worker(pending, done)
	}
	consumeWork(done)
	close(done)
}

func consumeWork(ch chan *Task) {
	for {
		fmt.Println(<-ch)
	}
}

func sendWork(ch chan *Task) {
	for i := 0; i < 10; i++ {
		ch <- &Task {id: i}
	}
}

func Worker(in, out chan *Task) {
	for {
		t := <-in
		process(t)
		out <- t
	}
}

func process(t *Task) {
	t.id = t.id + 1
}