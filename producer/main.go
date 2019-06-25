package main

import (
	"fmt"
	"net/http"
	"os"

	configEnv "github.com/joho/godotenv"
	"github.com/wuriyanto48/go-rabbitmq-example/producer/internal/handler"
	"github.com/wuriyanto48/go-rabbitmq-example/producer/internal/pub"
)

func main() {
	fmt.Println("Hello Rabbit")

	err := configEnv.Load(".env")

	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	rabbitAddress, ok := os.LookupEnv("RABBITMQ_ADDRESS")

	if !ok {
		fmt.Println("cannot load RABBITMQ_ADDRESS from environment")
		os.Exit(2)
	}

	rabbitKey, ok := os.LookupEnv("RABBITMQ_KEY")

	if !ok {
		fmt.Println("cannot load RABBITMQ_KEY from environment")
		os.Exit(2)
	}

	publisher, err := pub.NewPublisher(rabbitAddress)
	if err != nil {
		fmt.Printf("Error create publisher %v", err.Error())
		os.Exit(2)
	}

	publisherHandler := handler.NewHTTPHandler(rabbitKey, publisher)

	http.Handle("/api/send", publisherHandler.PublishMessages())
	http.ListenAndServe(":3000", nil)

}
