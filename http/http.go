package http

import (
	"net"
	"strings"

	"github.com/Gabriel-Nugent/go-server/util"
)

// HttpRequest contains the information about a given request
type HttpRequest struct {
	Method  string
	Target  string
	Version string
	Headers map[string]string
}

// toString converts a HttpRequest to a single string
func (req HttpRequest) String() string {
	str := ""

	str += req.Method + " "
	str += req.Target + "\r\n"
	for key, value := range req.Headers {
		str += key + ": " + value + "\r\n"
	} 

	return str
}

// ProcessRequest takes in request data
// and fills in a HttpRequest struct
func ProcessRequest(buffer []byte) HttpRequest {
	requestLines := strings.Split(string(buffer), "\r\n")
	var request HttpRequest

	// grab request method, target, and version from request line
	requestLine := strings.Split(requestLines[0], " ")
	request.Method = requestLine[0]
	request.Target = requestLine[1]
	request.Version = requestLine[2]

	request.Headers = make(map[string]string)

	for i := 1; i < len(requestLines); i++ {
		line := strings.Split(requestLines[i], " ")
		// check for empty line
		if len(line) > 1 {
			request.Headers[line[0][:len(line[0])-1]] = line[1]
		}
	}

	return request
}

// HttpResponse contains information for an http response
type HttpResponse struct {
	Headers map[string]string
	Body string
}

// NewReponse creates and returns a new http response
func NewResponse() HttpResponse {
	var res HttpResponse
	res.Headers = make(map[string]string)
	return res
}

// Send sends an http response to the client connection with the provided status and message
func (res HttpResponse) Send(connection net.Conn, status string, message string)  {
	// construct status line
	responseStr := "HTTP/1.1 " + status + " " + message + "\r\n"

	// add headers
	for key, value := range res.Headers {
		responseStr += key + ": " + value + "\r\n"
	}

	// attach body
	responseStr += "\r\n" + res.Body

	// send response
	_, err := connection.Write([]byte(responseStr))
	if err != nil {
		util.Error("Error reading: " + err.Error())
	} else {
		util.Success("Sent response: " + responseStr)
	}
}
