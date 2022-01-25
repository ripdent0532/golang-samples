package main

import (
	"encoding/gob"
	"fmt"
	"github.com/rip0532/golang-sample/utils"
	"os"
)

func main() {

	type Address struct {
		Type    string
		City    string
		Country string
	}

	type VCard struct {
		FirstName string
		LastName  string
		Addresses []*Address
		Remark    string
	}

	vc := &VCard{}

	file, _ := os.Open(*utils.HomeDir() + "vcard.gob")
	defer file.Close()
	dec := gob.NewDecoder(file)
	dec.Decode(vc)

	fmt.Println(vc)
}