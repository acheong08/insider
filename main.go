package main

import (
	"fmt"

	"github.com/acheong08/politics/crawlers/senate"
)

func main() {
	results, err := senate.GetLatestReports(10)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
}
