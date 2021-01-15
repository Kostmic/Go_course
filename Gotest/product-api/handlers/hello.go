package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
	logger *log.Logger
}

// Newhello creates a new hello handler with the given logger
func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

// ServeHTTP implements the go http.Handler interface
func (hello *Hello) ServeHTTP(responseWriter http.ResponseWriter, reader *http.Request) {

	// read the body
	hello.logger.Println("Hello World")
	output, err := ioutil.ReadAll(reader.Body)
	if err != nil {
		http.Error(responseWriter, "Oops", http.StatusBadRequest)
		return
	}
	// Write the response
	fmt.Fprintf(responseWriter, "Hello %s\n", output)
}
