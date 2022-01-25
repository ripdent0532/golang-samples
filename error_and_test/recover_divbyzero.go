package main

import "fmt"

func badCall2() {
	a, b := 10, 0
	n := a / b
	fmt.Println(n)
}

func test2() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall2()
	fmt.Printf("After bad call\r\n")
}

func main() {
	fmt.Printf("Call test\r\n")
	test2()
	fmt.Printf("Test completed\r\n")
}
