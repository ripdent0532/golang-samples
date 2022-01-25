package main

import (
	"fmt"
	"github.com/rip0532/golang-sample/utils"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := *utils.HomeDir() + "products.txt"
	outputFile := *utils.HomeDir() + "products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	//fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
