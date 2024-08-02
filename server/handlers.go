package server

import (
	"net"
	"github.com/Gabriel-Nugent/go-server/http"
)

// maps for storing target handlers
var getHandlers = make(map[string]func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))
var headHandlers = make(map[string]func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))
var postHandlers = make(map[string]func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))
var putHandlers = make(map[string]func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))
var deleteHandlers = make(map[string]func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))
var connectHandlers = make(map[string]func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))
var optionsHandlers = make(map[string]func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))
var traceHandlers = make(map[string]func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))
var patchHandlers = make(map[string]func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))

// Get adds a new handler to the server for GET requests
func Get(route string, handler func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))  {
	getHandlers[route] = handler
}

// Head adds a new handler to the server for HEAD requests
func Head(route string, handler func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))  {
	headHandlers[route] = handler
}

// Post adds a new handler to the server for POST requests
func Post(route string, handler func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))  {
	postHandlers[route] = handler
}

// Put adds a new handler to the server for PUT requests
func Put(route string, handler func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))  {
	putHandlers[route] = handler
}

// Delete adds a new handler to the server for DELETE requests
func Delete(route string, handler func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))  {
	deleteHandlers[route] = handler
}

// Connect adds a new handler to the server for CONNECT requests
func Connect(route string, handler func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))  {
	connectHandlers[route] = handler
}

// Options adds a new handler to the server for OPTIONS requests
func Options(route string, handler func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))  {
	optionsHandlers[route] = handler
}

// Trace adds a new handler to the server for TRACE requests
func Trace(route string, handler func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))  {
	traceHandlers[route] = handler
}

// Patch adds a new handler to the server for PATCH requests
func Patch(route string, handler func(conn net.Conn, req http.HttpRequest, res http.HttpResponse))  {
	patchHandlers[route] = handler
}

// handleRequest takes an http req and forwards it to the appropriate handler
func handleRequest(connection net.Conn, req http.HttpRequest)  {
	res := http.NewResponse()
	notFound := false

	switch req.Method {
		case "GET":
			if handler, exists := getHandlers[req.Target]; exists {
				handler(connection, req, res)
			} else {
				notFound = true
			}
		case "HEAD":
			if handler, exists := headHandlers[req.Target]; exists {
				handler(connection, req, res)
			} else {
				notFound = true
			}
		case "POST":
			if handler, exists := postHandlers[req.Target]; exists {
				handler(connection, req, res)
			} else {
				notFound = true
			}
		case "PUT":
			if handler, exists := putHandlers[req.Target]; exists {
				handler(connection, req, res)
			} else {
				notFound = true
			}
		case "DELETE":
			if handler, exists := deleteHandlers[req.Target]; exists {
				handler(connection, req, res)
			} else {
				notFound = true
			}
		case "CONNECT":
			if handler, exists := connectHandlers[req.Target]; exists {
				handler(connection, req, res)
			} else {
				notFound = true
			}
		case "OPTIONS":
			if handler, exists := optionsHandlers[req.Target]; exists {
				handler(connection, req, res)
			} else {
				notFound = true
			}
		case "TRACE":
			if handler, exists := traceHandlers[req.Target]; exists {
				handler(connection, req, res)
			} else {
				notFound = true
			}
		case "PATCH":
			if handler, exists := patchHandlers[req.Target]; exists {
				handler(connection, req, res)
			} else {
				notFound = true
			}
		default:
			notFound = true
	}

	if notFound {
		res.Send(connection, "404", "Not Found")
	}
}
