package main

import (
	"fmt"
	"github.com/rip0532/golang-sample/utils"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.Open(*utils.HomeDir() + "products2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		// scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)

	fmt.Println(filepath.Base(*utils.HomeDir() + "products2.txt"))
}
