package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("a.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//writing byte slice to file using package os
	bs1 := []byte("I learn Golang! 传")
	bw1, err := file.Write(bs1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bw1)

	//writing byte slice to file using ioutil package
	// ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
	// if the file doesn't exist WriteFile() creates it
	// and if it already exists the function will truncate it before writing to file.
	bs2 := []byte("Go programming is cool!")
	err = ioutil.WriteFile("a.txt", bs2, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
