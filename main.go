package main

import (
	"net"
	"os"

	"github.com/Gabriel-Nugent/go-server/http"
	"github.com/Gabriel-Nugent/go-server/server"
	"github.com/Gabriel-Nugent/go-server/util"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {

	server.Get("/", home)
	server.Run(HOST,PORT)
}

func home(conn net.Conn, req http.HttpRequest, res http.HttpResponse) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		util.Error(err.Error())
	}
	res.Body = string(data)
	res.Send(conn, "200", "OK")
}
