package main

import (
	"bufio"
	"fmt"
	"github.com/rip0532/golang-sample/utils"
	"os"
)

func main() {
	outputFile, outputError := os.OpenFile(*utils.HomeDir() + "output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}

	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()

}
