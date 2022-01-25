package main

import (
	"bufio"
	"fmt"
	"github.com/rip0532/golang-sample/utils"
	"io/ioutil"
	"os"
)

type Page struct {
	Title string
	Body []byte
}

func (page *Page) save() {
	file, err := os.OpenFile(*utils.HomeDir() + page.Title, os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	writer.Write(page.Body)
	writer.Flush()

}

func load(title string) {
	file, err := ioutil.ReadFile(*utils.HomeDir() + title)
	if err != nil {
		fmt.Printf("打开文件%v失败, 错误信息：%v", title, err)
	}
	fmt.Println(file)
}