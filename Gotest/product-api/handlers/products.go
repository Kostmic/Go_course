package handlers

import (
	"log"
	"net/http"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{logger}
}

func (product *Products) ServeHTTP(responseWriter http.ResponseWriter, reader *http.Request) {

}
