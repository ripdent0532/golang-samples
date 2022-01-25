package main

import (
	"encoding/gob"
	"github.com/rip0532/golang-sample/utils"
	"log"
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


	pa := &Address{"private", "Aartselaar","Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa,wa}, "none"}

	file, _ := os.OpenFile(*utils.HomeDir() + "vcard.gob", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
}