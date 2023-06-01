package main

import (
	"encoding/json"
	"fmt"

	senate_crawler "github.com/acheong08/politics/crawlers/senate"
	senate_parser "github.com/acheong08/politics/parsers/senate"
	network "github.com/acheong08/politics/utilities/network"
)

func main() {
	client := network.Init()
	results, err := senate_crawler.GetLatestReports(client, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
	transactions, err := senate_parser.GetPTR(client, "ea347ad9-be7e-4ccf-b353-a671bb9fd9f1")
	if err != nil {
		panic(err)
	}
	transactions_json, _ := json.MarshalIndent(transactions, "", "  ")
	fmt.Println(string(transactions_json))
}
