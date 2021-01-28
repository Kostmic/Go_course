package handlers

import (
	"log"
	"net/http"
	"strconv"

	"../data"
	"github.com/gorilla/mux"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (products *Products) GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
	products.logger.Println("Handle GET Products")
	// fetch the products from the datastore
	productList := data.GetProducts()

	// Serialize the list to JSON
	err := productList.ToJSON(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (products *Products) AddProduct(responseWriter http.ResponseWriter, request *http.Request) {
	products.logger.Println("Handle POST Product")

	prod := &data.Product{}

	err := prod.FromJSON(request.Body)
	if err != nil {
		http.Error(responseWriter, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (products Products) UpdateProducts(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(responseWriter, "Unable to convert id", http.StatusBadRequest)
		return
	}
	products.logger.Println("Handle PUT Product", id)

	prod := &data.Product{}

	err = prod.FromJSON(request.Body)
	if err != nil {
		http.Error(responseWriter, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(responseWriter, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(responseWriter, "Product not found", http.StatusInternalServerError)
		return
	}
}
