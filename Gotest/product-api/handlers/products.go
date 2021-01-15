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

	//handle an update
	if reader.Method == http.MethodPost {
		products.addProduct(responseWriter, reader)
		return
	}
	//catch all
	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (products *Products) getProducts(responseWriter http.ResponseWriter, reader *http.Request) {
	products.logger.Println("Handle GET Products")
	// fetch the products from the datastore
	productList := data.GetProducts()

	// Serialize the list to JSON
	err := productList.ToJSON(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (products *Products) addProduct(responseWriter http.ResponseWriter, reader *http.Request) {
	products.logger.Println("Handle POST Products")

	product := &data.Product{}

	err := product.FromJSON(reader.Body)
	if err != nil {
		http.Error(responseWriter, "Unabløe to unmarshal JSON", http.StatusBadRequest)
	}

	products.logger.Printf("Prod: %#v", product)
}
