package handlers

import (
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

func (products *Products) ServeHTTP(responseWriter http.ResponseWriter, reader *http.Request) {
	if reader.Method == http.MethodGet {
		products.getProducts(responseWriter, reader)
		return
	}

	//catch all
	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (products *Products) getProducts(responseWriter http.ResponseWriter, reader *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJSON(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}
