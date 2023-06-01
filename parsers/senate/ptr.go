package senate

import (
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/acheong08/politics/utilities"
	"github.com/acheong08/soup"
	http "github.com/bogdanfinn/fhttp"
)

var HEADERS map[string]string = map[string]string{
	"user-agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
	"referer":    "https://efdsearch.senate.gov/search/home/",
}

func GetPTR(ptr string) ([]Transaction, error) {
	URL := "https://efdsearch.senate.gov/search/view/ptr/" + ptr + "/"
	req, _ := http.NewRequest("GET", URL, nil)
	utilities.SetHeaders(req, HEADERS)
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
		re := regexp.MustCompile(`\$([0-9]+\.[0-9]+)|([0-9]+\/[0-9]+\/[0-9]+)`)
		matches := re.FindAllStringSubmatch(full_details, -1)
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
