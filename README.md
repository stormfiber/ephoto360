# Photo360 Go Package

[![Go Reference](https://pkg.go.dev/badge/github.com/stormfiber/ephoto360.svg)](https://pkg.go.dev/github.com/stormfiber/ephoto360)
[![Go Report Card](https://goreportcard.com/badge/github.com/stormfiber/ephoto360)](https://goreportcard.com/report/github.com/stormfiber/ephoto360)
[![GitHub release](https://img.shields.io/github/release/stormfiber/ephoto360.svg)](https://github.com/stormfiber/ephoto360/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://github.com/stormfiber/ephoto360/workflows/Go/badge.svg)](https://github.com/stormfiber/ephoto360/actions)

A powerful and easy-to-use Go package for generating stunning photo effects using EPhoto360 and similar websites. Create professional-looking text effects programmatically with just a few lines of code!

## 🌟 Features

- ✅ **Universal Support**: Works with any photo360.com effect URL
- ✅ **Flexible Input**: Single and multiple text input support
- ✅ **Smart Automation**: Automatic form data extraction and token handling
- ✅ **Robust Error Handling**: Comprehensive error management with helpful messages
- ✅ **Session Management**: Built-in cookie and session handling
- ✅ **Type Safety**: Full Go type safety with structured responses
- ✅ **Well Documented**: Complete documentation with examples
- ✅ **Production Ready**: Tested, reliable, and performant

## 🚀 Quick Start

### Installation

```bash
go get github.com/stormfiber/ephoto360
```

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/stormfiber/ephoto360"
)

func main() {
    // Create a new Photo360 instance
    p360, err := photo360.NewPhoto360("https://en.ephoto360.com/handwritten-text-on-foggy-glass-online-680.html")
    if err != nil {
        log.Fatal(err)
    }
    
    // Set your text
    p360.SetName("Hello World")
    
    // Generate the effect
    result, err := p360.Execute()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Status: %t\n", result.Status)
    fmt.Printf("Image URL: %s\n", result.ImageURL)
    fmt.Printf("Session ID: %s\n", result.SessionID)
}
```

**Output:**
```
Status: true
Image URL: https://e2.yotools.net/2025/01/680_12345.jpg
Session ID: 12345
```

## 📖 Complete Documentation

### Creating an Instance

#### With Default Effect
```go
// Uses default handwritten text on foggy glass effect
p360, err := photo360.NewPhoto360("")
if err != nil {
    log.Fatal(err)
}
```

#### With Custom Effect URL
```go
// Use any photo360.com effect URL
p360, err := photo360.NewPhoto360("https://en.ephoto360.com/write-text-on-wet-glass-online-589.html")
if err != nil {
    log.Fatal(err)
}
```

### Setting Text Input

#### Single Text Input
```go
// Perfect for simple effects that need one text
p360.SetName("Your Amazing Text")
```

#### Multiple Text Inputs
```go
// For effects that support multiple text fields
p360.SetNames([]string{"First Line", "Second Line", "Third Line"})
```

### Generating Effects

#### Basic Generation
```go
result, err := p360.Execute()
if err != nil {
    log.Fatal(err)
}

if result.Status {
    fmt.Printf("🎉 Success! Your image: %s\n", result.ImageURL)
} else {
    fmt.Println("❌ Generation failed")
}
```

#### With Error Handling
```go
result, err := p360.Execute()
if err != nil {
    if strings.Contains(err.Error(), "Please Try Using A Url That Requires 1 Input Field") {
        // Some effects only support single input
        fmt.Println("⚠️ This effect requires single input, switching...")
        p360.SetName("Single Text")
        result, err = p360.Execute()
    }
    
    if err != nil {
        log.Fatal(err)
    }
}
```

## 🎯 Advanced Examples

### Multiple Effects Processing
```go
package main

import (
    "fmt"
    "log"
    "time"
    
    "github.com/stormfiber/ephoto360"
)

func main() {
    effects := []string{
        "https://en.ephoto360.com/handwritten-text-on-foggy-glass-online-680.html",
        "https://en.ephoto360.com/write-text-on-wet-glass-online-589.html",
        "https://en.ephoto360.com/neon-light-text-effect-online-879.html",
    }
    
    text := "Amazing Effect"
    
    for i, effectURL := range effects {
        fmt.Printf("\n🎨 Processing Effect %d...\n", i+1)
        
        p360, err := photo360.NewPhoto360(effectURL)
        if err != nil {
            fmt.Printf("❌ Error: %v\n", err)
            continue
        }
        
        p360.SetName(text)
        result, err := p360.Execute()
        
        if err == nil && result.Status {
            fmt.Printf("✅ Success! Image: %s\n", result.ImageURL)
        } else {
            fmt.Printf("❌ Failed: %v\n", err)
        }
        
        // Be respectful - add delay between requests
        time.Sleep(2 * time.Second)
    }
}
```

### Batch Text Processing
```go
package main

import (
    "fmt"
    "log"
    
    "github.com/stormfiber/ephoto360"
)

func main() {
    texts := []string{"Hello", "World", "Photo360", "Amazing", "Effects"}
    effectURL := "https://en.ephoto360.com/handwritten-text-on-foggy-glass-online-680.html"
    
    p360, err := photo360.NewPhoto360(effectURL)
    if err != nil {
        log.Fatal(err)
    }
    
    var results []string
    
    for _, text := range texts {
        fmt.Printf("🔄 Processing: %s\n", text)
        
        p360.SetName(text)
        result, err := p360.Execute()
        
        if err == nil && result.Status {
            results = append(results, result.ImageURL)
            fmt.Printf("✅ Generated: %s\n", result.ImageURL)
        } else {
            fmt.Printf("❌ Failed for '%s': %v\n", text, err)
        }
    }
    
    fmt.Printf("\n🎉 Successfully generated %d images!\n", len(results))
    for i, url := range results {
        fmt.Printf("%d. %s\n", i+1, url)
    }
}
```

### Custom HTTP Client Configuration
```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"
    
    "github.com/stormfiber/ephoto360"
)

func main() {
    // For advanced users who need custom HTTP configuration
    // Note: This would require extending the package to accept custom clients
    
    p360, err := photo360.NewPhoto360("")
    if err != nil {
        log.Fatal(err)
    }
    
    p360.SetName("Custom Config")
    
    // Execute with context for timeout control
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    // This would be a future enhancement
    result, err := p360.Execute()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Result: %+v\n", result)
}
```

## 🔧 API Reference

### Types

#### Photo360
The main struct for photo360 operations.

```go
type Photo360 struct {
    EffectPageURL string                 // The effect page URL
    InputText     []string               // Text inputs for the effect
    FormData      map[string]interface{} // Internal form data
}
```

#### Photo360Result
Represents the result of photo360 execution.

```go
type Photo360Result struct {
    Status    bool   `json:"status"`    // Whether generation was successful
    ImageURL  string `json:"imageUrl"`  // URL of the generated image
    SessionID string `json:"sessionId"` // Session ID for tracking
}
```

### Functions

#### NewPhoto360(effectPageURL string) (*Photo360, error)

Creates a new Photo360 instance.

**Parameters:**
- `effectPageURL` (string): The URL of the photo360 effect page. Must contain "photo360.com". Use empty string for default effect.

**Returns:**
- `*Photo360`: New Photo360 instance
- `error`: Error if URL is invalid

**Example:**
```go
p360, err := photo360.NewPhoto360("https://en.ephoto360.com/handwritten-text-on-foggy-glass-online-680.html")
```

#### (p *Photo360) SetName(name string)

Sets a single text input for the effect.

**Parameters:**
- `name` (string): The text to be used in the effect

**Example:**
```go
p360.SetName("Hello World")
```

#### (p *Photo360) SetNames(names []string)

Sets multiple text inputs for the effect.

**Parameters:**
- `names` ([]string): Array of text strings to be used in the effect

**Example:**
```go
p360.SetNames([]string{"First Text", "Second Text"})
```

#### (p *Photo360) Execute() (*Photo360Result, error)

Executes the photo360 effect generation process.

This method performs the complete workflow:
1. Fetches the initial effect page
2. Extracts necessary form data and tokens
3. Submits the text inputs
4. Creates and retrieves the final image

**Returns:**
- `*Photo360Result`: Result containing status, image URL, and session ID
- `error`: Error if the process fails

**Example:**
```go
result, err := p360.Execute()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Generated: %s\n", result.ImageURL)
```

## ⚠️ Error Handling

The package provides comprehensive error handling for various scenarios:

### Common Errors and Solutions

#### Invalid URL Error
```go
p360, err := photo360.NewPhoto360("https://example.com/invalid")
// Error: invalid URL: Must be a photo360.com URL
```

**Solution:** Use a valid photo360.com URL.

#### Single Input Requirement Error
```go
result, err := p360.Execute()
if err != nil && strings.Contains(err.Error(), "Please Try Using A Url That Requires 1 Input Field") {
    // Some effects only support single text input
    p360.SetName("Single Text Only")
    result, err = p360.Execute()
}
```

#### Network Timeout Error
```go
result, err := p360.Execute()
if err != nil {
    if strings.Contains(err.Error(), "timeout") {
        fmt.Println("Request timed out, please try again")
        // Implement retry logic if needed
    }
}
```

#### Form Data Extraction Error
```go
result, err := p360.Execute()
if err != nil {
    if strings.Contains(err.Error(), "no generated form value found") {
        fmt.Println("Effect page structure may have changed")
        // Try with a different effect URL
    }
}
```

### Best Practices for Error Handling

```go
func generateEffectSafely(effectURL, text string) (*photo360.Photo360Result, error) {
    p360, err := photo360.NewPhoto360(effectURL)
    if err != nil {
        return nil, fmt.Errorf("failed to create instance: %w", err)
    }
    
    p360.SetName(text)
    result, err := p360.Execute()
    
    if err != nil {
        // Handle single input requirement
        if strings.Contains(err.Error(), "Please Try Using A Url That Requires 1 Input Field") {
            p360.SetName(text) // Ensure single input
            result, err = p360.Execute()
            if err != nil {
                return nil, fmt.Errorf("failed after single input retry: %w", err)
            }
        } else {
            return nil, fmt.Errorf("execution failed: %w", err)
        }
    }
    
    if !result.Status {
        return nil, fmt.Errorf("generation failed: status is false")
    }
    
    return result, nil
}
```

## 🎨 Supported Effects

This package works with hundreds of effects available on photo360.com. Here are some popular categories:

### Text Effects
- Handwritten text on foggy glass
- Neon light text effects
- 3D text effects
- Glowing text effects
- Metallic text effects

### Logo Creation
- Logo design effects
- Brand name effects
- Company logo effects

### Social Media
- Instagram story effects
- Facebook cover effects
- Profile picture effects

### Special Occasions
- Birthday effects
- Wedding effects
- Holiday effects
- Anniversary effects

### How to Find Effect URLs

1. Visit [ephoto360.com](https://ephoto360.com)
2. Browse or search for effects
3. Click on any effect you like
4. Copy the URL from your browser
5. Use it with this package!

## 🚦 Rate Limiting and Best Practices

### Respectful Usage

```go
import "time"

// Add delays between requests
func processMultipleEffects(effects []string, text string) {
    for _, effectURL := range effects {
        p360, err := photo360.NewPhoto360(effectURL)
        if err != nil {
            continue
        }
        
        p360.SetName(text)
        result, err := p360.Execute()
        
        // Process result...
        
        // Be respectful - wait between requests
        time.Sleep(2 * time.Second)
    }
}
```

### Production Recommendations

1. **Add Request Delays**: Wait 1-2 seconds between requests
2. **Implement Retry Logic**: Handle temporary failures gracefully
3. **Cache Results**: Store generated images to avoid repeated requests
4. **Monitor Usage**: Track your API usage and respect limits
5. **Error Logging**: Log errors for debugging and monitoring

## 🤝 Contributing

We welcome contributions! Here's how you can help:

### Reporting Issues

Found a bug? Please [open an issue](https://github.com/stormfiber/ephoto360/issues) with:

- Go version and OS
- Effect URL you were trying to use
- Complete error message
- Minimal code example

### Submitting Changes

1. Fork the repository
2. Create your feature branch: `git checkout -b feature/amazing-feature`
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass: `go test ./...`
6. Update documentation
7. Commit your changes: `git commit -m 'Add some amazing feature'`
8. Push to the branch: `git push origin feature/amazing-feature`
9. Open a Pull Request

### Development Setup

```bash
# Clone the repository
git clone https://github.com/stormfiber/ephoto360.git
cd photo360-go

# Install dependencies
go mod download

# Run tests
go test -v ./...

# Run with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 📦 Installation & Requirements

### Requirements
- Go 1.19 or later
- Internet connection (for API calls)

### Dependencies
- `github.com/PuerkitoBio/goquery` - HTML parsing and DOM manipulation

### Installation Methods

#### Go Modules (Recommended)
```bash
go get github.com/stormfiber/ephoto360
```

#### Specific Version
```bash
go get github.com/stormfiber/ephoto360@v1.0.0
```

#### Latest Development Version
```bash
go get github.com/stormfiber/ephoto360@main
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```
MIT License

Copyright (c) 2025 stormfiber

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## 🙏 Acknowledgments

- [EPhoto360](https://ephoto360.com) for providing amazing photo effect services
- [goquery](https://github.com/PuerkitoBio/goquery) for excellent HTML parsing capabilities
- Go community for their valuable feedback and contributions

## 📞 Support

- 📧 **Issues**: [GitHub Issues](https://github.com/stormfiber/ephoto360/issues)
- 📖 **Documentation**: [pkg.go.dev](https://pkg.go.dev/github.com/stormfiber/ephoto360)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/stormfiber/ephoto360/discussions)

## 🔗 Links

- **Repository**: https://github.com/stormfiber/ephoto360
- **Package**: https://pkg.go.dev/github.com/stormfiber/ephoto360
- **Issues**: https://github.com/stormfiber/ephoto360/issues
- **Releases**: https://github.com/stormfiber/ephoto360/releases

---

**Made with ❤️ for the Go community**

If this package helps you, please consider giving it a ⭐ on GitHub!