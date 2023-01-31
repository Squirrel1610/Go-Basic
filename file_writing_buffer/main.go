package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("buffer.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Creating a buffered writer from the file using bufio.NewWriter()
	bufferedWriter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99}

	// writing the byte slice to the buffer in memory
	bw1, err := bufferedWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file): %d\n", bw1)

	//check the available buffer
	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable)

	// writing a string (not a byte slice) to the buffer in memory
	bw2, err := bufferedWriter.WriteString("\nJust a random string")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file): %d\n", bw2)

	//check how much data is stored in buffer
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
	// -> 24 (3 bytes in the byte slice + 21 runes in the string, each rune is 1 byte)

	//push all data from buffer to file
	bufferedWriter.Flush()
}
