package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/acheong08/politics/crawlers/congress"
	senate_crawler "github.com/acheong08/politics/crawlers/senate"
	senate_parser "github.com/acheong08/politics/parsers/senate"
	"github.com/acheong08/politics/utilities/network"
)

// Testing
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
	entries, err := congress.GetEntriesByYear(2023)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, entry := range entries {
		fmt.Println(entry)
	}
}
