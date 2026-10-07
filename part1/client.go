package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8030")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("Enter message: ")

	message, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	_, err =conn.Write([]byte(message))

	if err != nil{
		panic(err)
	}

	fmt.Println("message sent!")

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		panic(err)
	}

	fmt.Println("Server replied:", string(buffer[:n]))

	}

	
}