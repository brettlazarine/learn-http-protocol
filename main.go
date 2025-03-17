package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const filePath = "./httpfromtcp/messages.txt"

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", filePath, err)
	}
	defer file.Close() // may need to remove this call?

	fmt.Printf("Reading data from %v\n", filePath)
	fmt.Println("==============================")

	for {
		buffer := make([]byte, 8, 8)
		n, err := file.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %v\n", err.Error())
			break
		}

		str := string(buffer[:n])
		fmt.Printf("read: %v\n", str)
	}
}