package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Status 这个结构体会保存解析后的返回数据
// 他们会形成有层级的XML，可以忽略一些无用的数据
type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status Status
}

func main() {
	response, _ := http.Get("http://twitter.com/users/Googland.xml")
	user := User{xml.Name{"", "user"}, Status{""}}
	body, _ := ioutil.ReadAll(response.Body)
	xml.Unmarshal(body, &user)
	fmt.Printf("status: %s", user.Status.Text)
}
