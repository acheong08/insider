package main

import (
	"encoding/json"
	"fmt"

	senate_crawler "github.com/acheong08/politics/crawlers/senate"
	senate_parser "github.com/acheong08/politics/parsers/senate"
)

func main() {
	client := senate_crawler.Init()
	results, err := senate_crawler.GetLatestReports(10)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
	senate_parser.Init(client)
	transactions, err := senate_parser.GetPTR("ea347ad9-be7e-4ccf-b353-a671bb9fd9f1")
	if err != nil {
		panic(err)
	}
	transactions_json, _ := json.MarshalIndent(transactions, "", "  ")
	fmt.Println(string(transactions_json))
}
