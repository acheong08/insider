package senate

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"insider/utilities"
	"insider/utilities/network"

	"github.com/acheong08/soup"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

func GetPTR(client *tls_client.HttpClient, ptr string) ([]Transaction, error) {
	URL := "https://efdsearch.senate.gov/search/view/ptr/" + ptr + "/"
	req, _ := http.NewRequest("GET", URL, nil)
	utilities.SetHeaders(req, network.HEADERS)
	resp, err := (*client).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	doc := soup.HTMLParse(string(body))
	// Find first tbody element
	tbody := doc.Find("tbody")
	// Find all tr elements
	trs := tbody.FindAll("tr")
	transactions := make([]Transaction, len(trs))
	for i, tr := range trs {
		tds := tr.FindAll("td")
		transactions[i] = Transaction{
			Date:    tds[1].Text(),
			Owner:   tds[2].Text(),
			Ticker:  tds[3].Text(),
			Type:    tds[6].Text(),
			Comment: tds[8].Text(),
		}
		asset := Asset{
			Type: tds[5].Text(),
		}
		// Find amount range
		amount_range := tds[7].Text()
		// Remove commas
		amount_range = strings.ReplaceAll(amount_range, ",", "")
		// Regex to find $<int> (2 instances)
		re := regexp.MustCompile(`\$([0-9]+)`)
		matches := re.FindAllStringSubmatch(amount_range, -1)
		if len(matches) != 2 {
			return nil, fmt.Errorf("expected 2 matches, got %d", len(matches))
		}
		// Convert to int
		amount_min, err := strconv.Atoi(matches[0][1])
		if err != nil {
			return nil, err
		}
		amount_max, err := strconv.Atoi(matches[1][1])
		if err != nil {
			return nil, err
		}
		transactions[i].AmountRange = Range{
			Min: amount_min,
			Max: amount_max,
		}
		// Parse asset information
		asset_info := tds[4]
		// Asset name is the first text in asset_info (Not in a div)
		asset.Name = asset_info.Text()
		// Check if div with `text-muted` class exists
		asset_details := asset_info.Find("div", "class", "text-muted")
		if asset_details.Error != nil {
			if asset_details.Error.Type == soup.ErrElementNotFound {
				continue
			} else {
				return nil, asset_details.Error
			}
		}
		asset.OptionType = asset_details.Text()
		full_details := asset_details.FullText()
		// Regex to find $<float> and <int>/<int>/<int>
		re = regexp.MustCompile(`\$([0-9]+\.[0-9]+)|([0-9]+\/[0-9]+\/[0-9]+)`)
		matches = re.FindAllStringSubmatch(full_details, -1)
		if len(matches) == 0 {
			continue
		}
		// If there is a match, then there is a strike price
		asset.StrikePrice, err = strconv.ParseFloat(matches[0][1], 64)
		if err != nil {
			continue
		}
		// If there is a match, then there is an expiration date
		asset.Expires = matches[1][2]
		transactions[i].Asset = asset
	}
	return transactions, nil
}
