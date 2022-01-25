package main

import (
	"github.com/rip0532/golang-sample/utils"
	"os"
)

func main() {
	os.Stdout.WriteString("hello, world\n")
	f, _ := os.OpenFile(*utils.HomeDir() + "test", os.O_CREATE|os.O_WRONLY, 0)
	defer f.Close()
	f.WriteString("hello, world in a file\n")
}
