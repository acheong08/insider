package main

import (
	"fmt"
	"log"
	"os"

	"code.sajari.com/docconv"
)

func main() {
	file_path := os.Args[1]
	res, err := docconv.ConvertPath(file_path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
