package main

import (
	"bufio"
	"flag"
	"net"
	"fmt"
	"os"
)

func read(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func write(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Fprintln(conn, scanner.Text())
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()

	conn, err := net.Dial("tcp", *addrPtr)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	go read(conn)

	write(conn)
}
