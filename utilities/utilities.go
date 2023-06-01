package utilities

import (
	"fmt"
	"log"
	"regexp"

	http "github.com/bogdanfinn/fhttp"
)

func SetHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

func ExtractPtrFromHref(href string) (string, error) {
	log.Println("Extracting ptr from href")
	// Look for href="/search/view/ptr/<UUID>/"
	re := regexp.MustCompile(`href="/search/view/ptr/(.*)/"`)
	matches := re.FindStringSubmatch(href)
	if len(matches) != 2 {
		return "", fmt.Errorf("expected 2 matches, got %d", len(matches))
	}
	return matches[1], nil
}

func ExtractMiddlewareToken(html string) (string, error) {
	re := regexp.MustCompile(`<input type="hidden" name="csrfmiddlewaretoken" value="(.*)">`)
	matches := re.FindStringSubmatch(html)
	if len(matches) != 2 {
		return "", fmt.Errorf("expected 2 matches, got %d", len(matches))
	}
	return matches[1], nil
}
