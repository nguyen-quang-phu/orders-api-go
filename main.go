package main

import (
	"context"
	"fmt"

	"github.com/dreamsofcode-io/orders-api/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("Error starting the server: %v", err)
	}
}
