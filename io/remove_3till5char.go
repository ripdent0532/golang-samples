package main

import (
	"bufio"
	"fmt"
	"github.com/rip0532/golang-sample/utils"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open(*utils.HomeDir() + "goprogram")
	outputFile, _ := os.OpenFile(*utils.HomeDir() + "goprogramT", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)

	for {
		inputStirng, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}

		outputString := string(inputStirng[2:5]) + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
		outputWriter.Flush()
	}
	fmt.Println("Conversion done")
}
