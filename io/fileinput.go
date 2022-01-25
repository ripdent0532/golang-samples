package main

import (
	"bufio"
	"fmt"
	"github.com/rip0532/golang-sample/utils"
	"os"
)

func main() {
	//inputFile, inputError := os.Open("/Users/anthony/Documents/projects/golang-sample/io/input.dat")
	//if inputError != nil {
	//	fmt.Printf("An error occurred on opening the inputfile\n" +
	//		"Does the file exist?\n" +
	//		"Have you got acces to it?\n")
	//	return
	//}
	//defer inputFile.Close()
	//
	//inputReader := bufio.NewReader(inputFile)
	//for {
	//	inputString, readerError := inputReader.ReadString('\n')
	//	fmt.Printf("The input was: %s", inputString)
	//	if readerError == io.EOF {
	//		return
	//	}
	//}
	bufferRead()
}

func bufferRead() {
	inputFile, inputError := os.Open(*utils.HomeDir() + "input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	buf := make([]byte, 1024)
	for {
		n, _ := inputReader.Read(buf)
		if n == 0 {
			break
		}
	}
	fmt.Println(string(buf))
}
