package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//open the file
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
	scanner := bufio.NewScanner(file)

	// the default scanner is bufio.ScanLines and that means it will scan a file line by line.
	// there are also bufio.ScanWords and bufio.ScanRunes.
	// scanner.Split(bufio.ScanLines)

	success := scanner.Scan() //read a line
	if success == false {
		err := scanner.Err()
		if err == nil {
			log.Println("Scan was completed and it reached EOF")
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(scanner.Text())

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
