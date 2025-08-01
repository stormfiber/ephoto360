package main

import (
	"fmt"
	"log"

	"github.com/stormfiber/ephoto360"
)

func main() {
	// Create a new Photo360 instance with default effect
	p360, err := photo360.NewPhoto360("")
	if err != nil {
		log.Fatal(err)
	}

	// Set custom text
	p360.SetName("Hello World")

	// Execute and get result
	result, err := p360.Execute()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %t\n", result.Status)
	fmt.Printf("Image URL: %s\n", result.ImageURL)
	fmt.Printf("Session ID: %s\n", result.SessionID)
}

