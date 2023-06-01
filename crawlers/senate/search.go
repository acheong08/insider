package senate

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"insider/utilities"
	"insider/utilities/network"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

const SEARCH_ENDPOINT string = network.HOST + "/search/report/data/"

func GetLatestReports(client *tls_client.HttpClient, num int) ([]Report, error) {
	log.Println("Creating query params")
	query := NewQueryParams(0, num)
	log.Println(query)
	log.Println("Sending request")
	req, _ := http.NewRequest("POST", SEARCH_ENDPOINT, strings.NewReader(query))
	utilities.SetHeaders(req, network.HEADERS)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	// Get csrf token
	for _, cookie := range (*client).GetCookieJar().Cookies(&url.URL{Host: network.DOMAIN}) {
		if cookie.Name == "csrftoken" {
			req.Header.Set("X-CSRFToken", cookie.Value)
		}
	}
	resp, err := (*client).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("expected status code 200, got %d", resp.StatusCode)
	}
	log.Println("Parsing response")
	// Map response to struct
	var report_response ReportResponse
	err = json.NewDecoder(resp.Body).Decode(&report_response)
	if err != nil {
		return nil, err
	}
	log.Println("Extracting reports")
	var reports []Report = make([]Report, len(report_response.ReportData))
	for i, report := range report_response.ReportData {
		reports[i].FirstName = report[0]
		reports[i].LastName = report[1]
		reports[i].Office = report[2]
		reports[i].Date = report[4]

		// Extract ptr from href
		ptr, err := utilities.ExtractPtrFromHref(report[3])
		if err != nil {
			continue
		}
		reports[i].Ptr = ptr
	}
	return reports, nil
}
