package main

import (
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
	fmt.Println(senate_parser.GetPTR(results[1].Ptr))
}
