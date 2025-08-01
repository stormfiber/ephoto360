// Package photo360 provides functionality to generate photo effects using EPhoto360 and similar websites.
//
// This package allows you to programmatically create text-based photo effects by:
//   - Fetching effect pages from photo360 websites
//   - Extracting required form data
//   - Submitting text inputs
//   - Generating and retrieving the final image
//
// Basic usage:
//
//	p360, err := photo360.NewPhoto360("https://en.ephoto360.com/handwritten-text-on-foggy-glass-online-680.html")
//	if err != nil {
//		log.Fatal(err)
//	}
//	p360.SetName("Your Text")
//	result, err := p360.Execute()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Generated image: %s\n", result.ImageURL)
package photo360

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Photo360 represents the main struct for photo360 operations
type Photo360 struct {
	EffectPageURL string
	InputText     []string
	FormData      map[string]interface{}
}

// Photo360Result represents the result of photo360 execution
type Photo360Result struct {
	Status    bool   `json:"status"`
	ImageURL  string `json:"imageUrl"`
	SessionID string `json:"sessionId"`
}

// ImageCreationResponse represents the response from image creation API
type ImageCreationResponse struct {
	Success       bool        `json:"success"`
	Image         string      `json:"image"`
	FullsizeImage string      `json:"fullsize_image"`
	SessionID     interface{} `json:"session_id"`
}

// FormDataValues represents the parsed form data
type FormDataValues struct {
	ID            string `json:"id"`
	Token         string `json:"token"`
	BuildServer   string `json:"build_server"`
	BuildServerID string `json:"build_server_id"`
	Radio0        struct {
		Radio string `json:"radio"`
	} `json:"radio0"`
	Text []string `json:"text"`
}

// NewPhoto360 creates a new Photo360 instance with the given effect page URL
func NewPhoto360(effectPageURL string) (*Photo360, error) {
	if effectPageURL == "" {
		effectPageURL = "https://en.ephoto360.com/handwritten-text-on-foggy-glass-online-680.html"
	}
	if !strings.Contains(effectPageURL, "photo360.com") {
		return nil, errors.New("invalid URL: Must be a photo360.com URL")
	}
	return &Photo360{
		EffectPageURL: effectPageURL,
		InputText:     []string{"Faris"},
		FormData:      make(map[string]interface{}),
	}, nil
}

// SetName sets a single name for the photo effect
func (p *Photo360) SetName(name string) {
	p.InputText = []string{name}
}

// SetNames sets multiple names for the photo effect
func (p *Photo360) SetNames(names []string) {
	p.InputText = names
}

// Execute runs the complete photo360 process and returns the result
func (p *Photo360) Execute() (*Photo360Result, error) {
	cookies, err := p.fetchInitialPage()
	if err != nil {
		return nil, err
	}
	generatedFormValue, err := p.submitFormData(cookies)
	if err != nil {
		return nil, err
	}
	result, err := p.createImage(generatedFormValue, cookies)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// secureRandomInt generates a cryptographically secure random integer
func secureRandomInt(max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}

func (p *Photo360) fetchInitialPage() (string, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("GET", p.EffectPageURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.188")
	parsedURL, _ := url.Parse(p.EffectPageURL)
	req.Header.Set("Origin", fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host))
	req.Header.Set("Referer", p.EffectPageURL)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var cookies []string
	for _, cookie := range resp.Cookies() {
		cookies = append(cookies, cookie.String())
	}
	cookieString := strings.Join(cookies, "; ")

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	buildServer, _ := doc.Find("#build_server").Attr("value")
	buildServerId, _ := doc.Find("#build_server_id").Attr("value")
	token, _ := doc.Find("#token").Attr("value")
	submitValue, _ := doc.Find("#submit").Attr("value")

	var radioOptions []string
	doc.Find("input[name=\"radio0[radio]\"]").Each(func(i int, s *goquery.Selection) {
		if value, exists := s.Attr("value"); exists {
			radioOptions = append(radioOptions, value)
		}
	})

	formData, err := p.prepareFormData(buildServer, buildServerId, token, submitValue, radioOptions)
	if err != nil {
		return "", fmt.Errorf("failed to prepare form data: %w", err)
	}
	p.FormData = formData
	return cookieString, nil
}

func (p *Photo360) submitFormData(cookies string) (string, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	for key, value := range p.FormData {
		if err := writer.WriteField(key, fmt.Sprintf("%v", value)); err != nil {
			return "", fmt.Errorf("failed to write form field %s: %w", key, err)
		}
	}

	for _, text := range p.InputText {
		if err := writer.WriteField("text[]", text); err != nil {
			return "", fmt.Errorf("failed to write text field: %w", err)
		}
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %w", err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("POST", p.EffectPageURL, &buf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.188")
	req.Header.Set("Cookie", cookies)
	parsedURL, _ := url.Parse(p.EffectPageURL)
	req.Header.Set("Origin", fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host))
	req.Header.Set("Referer", p.EffectPageURL)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	generatedFormValue := p.extractGeneratedFormValue(string(body))
	if generatedFormValue == "" {
		return "", errors.New("no generated form value found")
	}
	return generatedFormValue, nil
}

func (p *Photo360) createImage(generatedFormValue, cookies string) (*Photo360Result, error) {
	parsedURL, _ := url.Parse(p.EffectPageURL)
	createImageURL := fmt.Sprintf("%s://%s/effect/create-image", parsedURL.Scheme, parsedURL.Host)

	var imageCreationData FormDataValues
	if err := json.Unmarshal([]byte(generatedFormValue), &imageCreationData); err != nil {
		return nil, errors.New("please try using a URL that requires 1 input field")
	}

	formData := url.Values{}
	formData.Set("id", imageCreationData.ID)
	formData.Set("token", imageCreationData.Token)
	formData.Set("build_server", imageCreationData.BuildServer)
	formData.Set("build_server_id", imageCreationData.BuildServerID)
	if imageCreationData.Radio0.Radio != "" {
		formData.Set("radio0[radio]", imageCreationData.Radio0.Radio)
	}
	for _, text := range imageCreationData.Text {
		formData.Add("text[]", text)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("POST", createImageURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.188")
	req.Header.Set("Cookie", cookies)
	req.Header.Set("Origin", fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host))
	req.Header.Set("Referer", p.EffectPageURL)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var imageResponse ImageCreationResponse
	if err := json.Unmarshal(body, &imageResponse); err != nil {
		return nil, err
	}

	imageURL := imageCreationData.BuildServer
	if imageResponse.Image != "" {
		imageURL += imageResponse.Image
	} else if imageResponse.FullsizeImage != "" {
		imageURL += imageResponse.FullsizeImage
	}

	var sessionID string
	switch v := imageResponse.SessionID.(type) {
	case string:
		sessionID = v
	case float64:
		sessionID = fmt.Sprintf("%.0f", v)
	case int:
		sessionID = fmt.Sprintf("%d", v)
	default:
		sessionID = fmt.Sprintf("%v", v)
	}

	return &Photo360Result{
		Status:    imageResponse.Success,
		ImageURL:  imageURL,
		SessionID: sessionID,
	}, nil
}

func (p *Photo360) prepareFormData(buildServer, buildServerId, token, submitValue string, radioOptions []string) (map[string]interface{}, error) {
	formData := map[string]interface{}{
		"submit":       submitValue,
		"token":        token,
		"build_server": buildServer,
	}
	if buildServerId != "" {
		if id, err := strconv.Atoi(buildServerId); err == nil {
			formData["build_server_id"] = id
		}
	}
	if len(radioOptions) > 0 {
		randomIndex, err := secureRandomInt(len(radioOptions))
		if err != nil {
			return nil, fmt.Errorf("failed to generate secure random number: %w", err)
		}
		formData["radio0[radio]"] = radioOptions[randomIndex]
	}
	return formData, nil
}

func (p *Photo360) extractGeneratedFormValue(responseData string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(responseData))
	if err != nil {
		return ""
	}

	formValue := strings.TrimSpace(doc.Find("#form_value").Text())
	if formValue == "" {
		formValue = strings.TrimSpace(doc.Find("#form_value_input").Text())
	}
	if formValue == "" {
		if val, exists := doc.Find("#form_value").Attr("value"); exists {
			formValue = val
		} else if val, exists := doc.Find("#form_value_input").Attr("value"); exists {
			formValue = val
		}
	}
	return formValue
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

