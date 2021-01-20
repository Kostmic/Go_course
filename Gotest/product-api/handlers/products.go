package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"../data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (products *Products) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		products.getProducts(responseWriter, request)
		return
	}

	//handle an update
	if request.Method == http.MethodPost {
		products.addProduct(responseWriter, request)
		return
	}

	if request.Method == http.MethodPut {
		products.logger.Println("PUT", request.URL.Path)
		// expect the id in the URI
		regex := regexp.MustCompile(`/([0-9]+)`)
		group := regex.FindAllStringSubmatch(request.URL.Path, -1)

		if len(group) != 1 {
			products.logger.Println("Invalid URI more than one id")
			http.Error(responseWriter, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			products.logger.Println("Invalid URI more than one capture group")
			http.Error(responseWriter, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := group[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			products.logger.Println("Invalid URI unable to convert to number", idString)
			http.Error(responseWriter, "Invalid URI", http.StatusBadRequest)
			return
		}

		products.updateProducts(id, responseWriter, request)
		return

		products.logger.Println("got id", id)
	}

	//catch all
	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (products *Products) getProducts(responseWriter http.ResponseWriter, request *http.Request) {
	products.logger.Println("Handle GET Products")
	// fetch the products from the datastore
	productList := data.GetProducts()

	// Serialize the list to JSON
	err := productList.ToJSON(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (products *Products) addProduct(responseWriter http.ResponseWriter, request *http.Request) {
	products.logger.Println("Handle POST Product")

	prod := &data.Product{}

	err := prod.FromJSON(request.Body)
	if err != nil {
		http.Error(responseWriter, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (products Products) updateProducts(id int, responseWriter http.ResponseWriter, request *http.Request) {
	products.logger.Println("Handle PUT Product")

	prod := &data.Product{}

	err := prod.FromJSON(request.Body)
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
