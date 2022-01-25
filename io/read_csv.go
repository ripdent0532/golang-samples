package main

import (
	"encoding/csv"
	"fmt"
	"github.com/rip0532/golang-sample/utils"
	"io"
	"os"
)

type data struct {
	title string
	price string
	quantity string
}

type lines []data

func main() {
	file, _ := os.Open(*utils.HomeDir() + "products.txt")
	reader := csv.NewReader(file)
	reader.Comma = ';'

	records := make(lines, 0)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		records = append(records, data{title: record[0], price: record[1], quantity: record[2]})
	}

	fmt.Println(records)




}