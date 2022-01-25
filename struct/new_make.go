package main

import "fmt"

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	// OK
	y := new(Bar)
	y.thingOne = "hello"
	y.thingTwo = 1

	y1 := &Bar{
		thingOne: "h",
		thingTwo: 2,
	}
	fmt.Println(y1)
	// 结构初始化可以使用new()或者&{}

	// NOT OK
	//z := make(Bar)
	//(*z).thingOne = "hello"
	//(*z).thingTwo = 1

	// 对于map，可以使用make进行初始化，如果使用new()方法，将会返回nil对象
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	u := new(Foo)
	(*u)["x"] = "goodbye"
	(*u)["y"] = "world"
}
