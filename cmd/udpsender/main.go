package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	r , err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, r) 
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}
}