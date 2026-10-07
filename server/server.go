package main

import (
	"bufio"
	"flag"
	"net"
	"fmt"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	for {
		conn, err := ln.Accept()

		handleError(err)

		conns <- conn
	}
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	scanner := bufio.NewScanner(client)

	for scanner.Scan() {

		msg:= Message{
			sender: clientid,
			message: scanner.Text(),
		}

		msgs <- msg
	}

	client.Close()
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	ln, err := net.Listen("tcp", *portPtr)

	handleError(err)

	defer ln.Close()

	fmt.Println("Server listening on", *portPtr)

	//TODO Create a Listener for TCP connections on the port given above.

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			clientid := len(clients)

			clients[clientid] = conn

			fmt.Println("Client connected:", clientid)

			go handleClient(conn, clientid, msgs)
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients map
			// - start to asynchronously handle messages from this client
		case msg := <-msgs:
			fmt.Println(
				"Client",
				msg.sender,
				":",
				msg.message,
			)
			for id, conn := range clients {

				if id != msg.sender {
					fmt.Fprintln(conn, msg.message)
				}
			}
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
		}
	}
}
