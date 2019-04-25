package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type Person struct {
	Name string
	Id   int
}

func main() {
	// gob.
	p := Person{
		Name: "Mohit",
		Id:   1,
	}
	var bytes bytes.Buffer
	encoder := gob.NewEncoder(&bytes)
	err := encoder.Encode(p)
	if err != nil {
		log.Fatalf("error occured while encoding")
	}
	dec := gob.NewDecoder(&bytes)
	newP := &Person{}
	err = dec.Decode(newP)
	if err != nil {
		log.Fatalf("error occured while decoding.")
	}
	fmt.Printf("name: %v\nid: %v\n", newP.Name, newP.Id)

	aErr := errors.New("error A")
	bErr := errors.New("error B")

	if aErr == bErr {
		fmt.Printf("errors are same")
	} else {
		fmt.Printf("errors are different")
	}

	// var buf bytes.Buffer
	// n, err := buf.Write([]byte("Apna Time Aayega"))
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// fmt.Printf("%v bytes written\n", n)
	// read(&buf)
}

func read(r io.Reader) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	str := string(bytes)
	fmt.Printf("str: %v\n", str)
}
