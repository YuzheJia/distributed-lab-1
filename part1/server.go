package main

import (
	"fmt"
	"net"
)

func  main() {
	ln, err := net.Listen("tcp", ":8030")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server is waiting...")

	for{
		conn, err := ln.Accept()

		if err != nil {
			panic(err)
		}

		fmt.Println("Client connected!")

		for {
			buffer := make([]byte, 1024)

			n, err := conn.Read(buffer)

			if err != nil {
				conn.Close()
				break
			}

			fmt.Println("Received:", string(buffer[:n]))

			_, err = conn.Write([]byte("Message received!"))

			if err != nil {
				conn.Close()
				break
			}
		}
	}
}