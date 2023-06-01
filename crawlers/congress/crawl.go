package congress

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetEntriesByYear(year int) ([]DocumentEntry, error) {
	URL := "https://disclosures-clerk.house.gov/public_disc/financial-pdfs/" + strconv.Itoa(year) + "FD.ZIP"
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// The response body is a application/x-zip-compressed
	// We need to unzip it
	body, _ := io.ReadAll(resp.Body)
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, err
	}
	document_txt := zipReader.File[0]
	document_txt_reader, err := document_txt.Open()
	if err != nil {
		return nil, err
	}
	defer document_txt_reader.Close()
	// Read the document.txt file
	document_txt_body, err := io.ReadAll(document_txt_reader)
	if err != nil {
		return nil, err
	}
	// Loop through each line
	lines := strings.Split(string(document_txt_body), "\n")
	entries := make([]DocumentEntry, len(lines)-2)
	counter := 0
	for i, line := range lines {
		if i == 0 {
			continue
		} else if i == len(lines)-1 {
			log.Println("Reached end of file")
			log.Println(line)
			break
		}
		if strings.HasPrefix(line, "	") {
			line = "N/A" + line
		}
		fields := strings.Split(line, "	")
		if len(fields) < 5 {
			log.Println("Skipping line")
			continue
		}
		for j, field := range fields {
			fields[j] = strings.TrimSpace(field)
		}
		if len(fields) != 9 {
			return nil, fmt.Errorf("expected 9 fields, got %d", len(fields))
		}
		if fields[4] != "P" {
			continue
		}
		entry := DocumentEntry{
			Prefix:        fields[0],
			LastName:      fields[1],
			FirstName:     fields[2],
			Suffix:        fields[3],
			FillingType:   fields[4],
			StateDistrict: fields[5],
			Year:          fields[6],
			FillingDate:   fields[7],
			DocID:         fields[8],
		}
		entries[counter] = entry
		counter++
	}
	entries = entries[:counter]
	return entries, nil
}
