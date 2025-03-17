package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Println("***** connection accepted *****")

		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Printf("%v\n", line)
		}

		fmt.Println("***** connection closed *****")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	currentLine := ""

	go func() {
		defer close(ch)
		defer f.Close()

		for {
			buffer := make([]byte, 8)
			n, err := f.Read(buffer)
			if err != nil {
				if currentLine != "" {
					ch <- currentLine
					currentLine = ""
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %v\n", err.Error())
				break
			}
	
			str := string(buffer[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts) - 1; i++ {
				ch <- currentLine + parts[i]
				currentLine = ""
			}
	
			currentLine += parts[len(parts) - 1]
		}
	}()

	return ch
}