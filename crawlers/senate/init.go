package senate

import (
	"io"
	"log"
	"net/url"
	"os"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"

	"github.com/acheong08/politics/utilities"
)

const HOST string = "https://efdsearch.senate.gov"
const DOMAIN string = "efdsearch.senate.gov"

var (
	jar     = tls_client.NewCookieJar()
	options = []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(360),
		tls_client.WithClientProfile(tls_client.Safari_IOS_16_0),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
	}
	http_proxy       = os.Getenv("http_proxy")
	initialized bool = false
)

var client tls_client.HttpClient

var HEADERS map[string]string = map[string]string{
	"user-agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
	"referer":    "https://efdsearch.senate.gov/search/home/",
}

var middleware_token string

func Init() *tls_client.HttpClient {
	initialized = true
	// Reset client
	client, _ = tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	log.Println("Senate crawler initialized")

	// Set proxy if it exists
	if http_proxy != "" {
		client.SetProxy(http_proxy)
	}

	// Initial request to fetch cookies for csrf token
	req, _ := http.NewRequest("GET", HOST+"/search/", nil)
	utilities.SetHeaders(req, HEADERS)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 302 {
		log.Printf("Expected status code 302, got %d \n", resp.StatusCode)
		// Read body
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("Body: %s", string(body))
	}
	// Next request to fetch csrf token
	redirected_path := resp.Header.Get("Location")
	req, _ = http.NewRequest("GET", HOST+redirected_path, nil)
	utilities.SetHeaders(req, HEADERS)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	middleware_token, err = utilities.ExtractMiddlewareToken(string(body))
	if err != nil {
		log.Fatal(err)
	}

	// Remove messages cookie
	client.SetCookies(&url.URL{Host: DOMAIN}, []*http.Cookie{
		{
			Name:  "messages",
			Value: "",
		},
	})

	// POST to /search/home/
	payload := "prohibition_agreement=1&csrfmiddlewaretoken=" + middleware_token
	req, _ = http.NewRequest("POST", HOST+"/search/home/", strings.NewReader(payload))
	utilities.SetHeaders(req, HEADERS)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 302 {
		log.Printf("Expected status code 302, got %d \n", resp.StatusCode)
		// Read body
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("Body: %s", string(body))
	}
	// Log session id cookie
	log.Printf("Session id cookie: %s\n", resp.Header.Get("Set-Cookie"))

	// Go to redirected page
	redirected_path = resp.Header.Get("Location")
	log.Printf("Redirected path: %s\n", redirected_path)
	req, _ = http.NewRequest("GET", HOST+redirected_path, nil)
	utilities.SetHeaders(req, HEADERS)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}
	// Get middleware token
	body, _ = io.ReadAll(resp.Body)
	middleware_token, err = utilities.ExtractMiddlewareToken(string(body))
	if err != nil {
		log.Fatal(err)
	}
	return &client
}
