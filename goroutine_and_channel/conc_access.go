package main

import (
	"fmt"
	"strconv"
)

// 为了保护一个对象的并发修改，可以使用一个后台的协程来顺序执行一个匿名函数，而不是通过同步互斥锁（Mutex）进行锁定

type Person struct {
	Name   string
	salary float64
	chF    chan func()
}

func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

func (p *Person) backend() {
	for f := range p.chF {
		f()
	}
}

// SetSalary 设置salary
func (p *Person) SetSalary(sal float64) {
	p.chF <- func() {
		p.salary = sal
	}
}

// Salary 取回salary
func (p *Person) Salary() float64 {
	fChan := make(chan float64)
	p.chF <- func() {
		fChan <- p.salary
	}
	return <-fChan
}

func (p *Person) String() string {
	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

func main() {
	bs := NewPerson("Smith Bill", 2500.5)
	//fmt.Println(bs)
	bs.SetSalary(4000.25)
	fmt.Println("Salary changed:")
	fmt.Println(bs)
}
