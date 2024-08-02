package server

import (
	"net"
	"os"

	"github.com/Gabriel-Nugent/go-server/http"
	"github.com/Gabriel-Nugent/go-server/util"
)

const TYPE = "tcp"

// Run starts the tcp server on the specified host and port
func Run(host string, port string)  {
	util.Log("Starting Server...")
	server, err := net.Listen(TYPE, host + ":" + port )

	if err != nil {
		util.Error("Error listening: " + err.Error())
		os.Exit(1)
	}

	defer server.Close()

	util.Success("Listening on " + host + ":" + port)
	util.Log("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			util.Error("Error accepting: " + err.Error())
			os.Exit(1)
		}
		go processClient(connection)
	}
}

// proccessClient receives a tcp message from a connected client
func processClient(connection net.Conn)  {
	buffer := make([]byte, 1024)
	defer connection.Close()

	messageLength, err := connection.Read(buffer)
	if messageLength < 1 {
		return
	}
	if err != nil {
		util.Error("Error reading: " + err.Error())
	}

	req := http.ProcessRequest(buffer[:messageLength])
	util.Log("Received: " + req.Method + " " + req.Target)

	handleRequest(connection, req)
}
