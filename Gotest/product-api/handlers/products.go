package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"../data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (product *Products) ServeHTTP(responseWriter http.ResponseWriter, reader *http.Request) {
	productList := data.GetProducts()
	data, err := json.Marshal(productList)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}

	responseWriter.Write(data)
}
