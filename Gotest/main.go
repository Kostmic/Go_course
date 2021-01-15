package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logger)
	byeHandler := handlers.NewGoodbye(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", byeHandler)

	server := &http.Server{
		Addr:              ":9090",
		Handler:           serveMux,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	signal := <-signalChannel
	logger.Println("Recieved terminate, graceful shutdown", signal)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
