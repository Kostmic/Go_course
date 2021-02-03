package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"./env"

	"./product-api/handlers"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handlers
	productHandler := handlers.NewProducts(logger)

	// create a new serve mux and register the handlers
	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productHandler.GetProducts)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProduct)
	putRouter.Use(productHandler.MiddlewareValidateProduct)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProduct)
	postRouter.Use(productHandler.MiddlewareValidateProduct)

	// create a new server
	server := &http.Server{ // configure the bind address
		Addr:         *bindAddress,      // set the default handler
		Handler:      serveMux,          // set the logger for the server
		ErrorLog:     logger,            // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to write response to the client
		WriteTimeout: 10 * time.Second,  // max time for the connections using TCP Keep-Alive
		IdleTimeout:  120 * time.Second, // max time to read request from the client
	}

	// start the server
	go func() {
		logger.Println("Starting server port 9090")

		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
	// trap sigterm or interrupt and gracefully shutdown the server
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	signal := <-signalChannel
	logger.Println("Recieved terminate, graceful shutdown", signal)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
