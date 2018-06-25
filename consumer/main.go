package main

import (
	"fmt"
	"os"

	configEnv "github.com/joho/godotenv"
	"github.com/wuriyanto48/go-rabbitmq-example/consumer/internal/sub"
	"github.com/wuriyanto48/go-rabbitmq-example/consumer/internal/handler"
)

func main() {
	fmt.Println("consumer")

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

	subscriber, err := sub.NewSubscriber(rabbitAddress)

	if err != nil {
		fmt.Println("error create subscriber")
		os.Exit(2)
	}

	workerHandler := handler.NewWorkerHandler(rabbitKey, subscriber)

	workerHandler.Pool()
}