/*
Package photo360 provides a Go client for generating photo effects using EPhoto360 and similar websites.

# Overview

This package allows you to programmatically create text-based photo effects by interacting with photo360 websites.
It handles the entire process of form submission, token extraction, and image generation.

# Features

  - Support for any photo360.com effect URL
  - Single and multiple text input support
  - Automatic form data extraction
  - Error handling for common issues
  - Session management and cookie handling

# Quick Start

	package main

	import (
		"fmt"
		"log"

		"github.com/stormfiber/photo360"
	)

	func main() {
		// Create a new instance
		p360, err := photo360.NewPhoto360("https://en.ephoto360.com/handwritten-text-on-foggy-glass-online-680.html")
		if err != nil {
			log.Fatal(err)
		}

		// Set your text
		p360.SetName("Your Name")

		// Generate the effect
		result, err := p360.Execute()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Generated image: %s\n", result.ImageURL)
	}

# Error Handling

The package handles various error scenarios:

  - Invalid URLs (must contain "photo360.com")
  - Network timeouts and connection issues
  - HTML and JSON parsing failures
  - Single input requirement violations

# Rate Limiting

Be respectful when using this package. The underlying websites may have rate limits.
Consider adding delays between requests in production applications.
*/
package photo360
