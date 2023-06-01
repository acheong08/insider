package main

import (
	"fmt"

	"github.com/acheong08/politics/crawlers/senate"
)

func main() {
	results, err := senate.GetLatestReports(100)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
}
