package main

import (
	"fmt"
	"strings"

	"github.com/stormfiber/ephoto360"
)

func main() {
	// List of different effects to try
	effects := []string{
		"https://en.ephoto360.com/handwritten-text-on-foggy-glass-online-680.html",
		"https://en.ephoto360.com/write-text-on-wet-glass-online-589.html",
	}

	textInputs := []string{"Hello", "World", "Photo360"}

	for i, effectURL := range effects {
		fmt.Printf("\n=== Effect %d: %s ===\n", i+1, effectURL)

		p360, err := photo360.NewPhoto360(effectURL)
		if err != nil {
			fmt.Printf("Error creating instance: %v\n", err)
			continue
		}

		// Try with single text first
		p360.SetNames(textInputs)
		result, err := p360.Execute()

		if err != nil {
			if strings.Contains(err.Error(), "Please Try Using A Url That Requires 1 Input Field") {
				fmt.Println("Effect requires single input, trying with first text only...")
				p360.SetName(textInputs[0])
				result, err = p360.Execute()
			}

			if err != nil {
				fmt.Printf("Error executing: %v\n", err)
				continue
			}
		}

		if result.Status {
			fmt.Printf("✅ Success! Image URL: %s\n", result.ImageURL)
			fmt.Printf("Session ID: %s\n", result.SessionID)
		} else {
			fmt.Println("❌ Generation failed")
		}
	}
}
