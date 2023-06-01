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
		entries[i-1].Prefix = fields[0]
		entries[i-1].LastName = fields[1]
		entries[i-1].FirstName = fields[2]
		entries[i-1].Suffix = fields[3]
		entries[i-1].FillingType = fields[4]
		entries[i-1].StateDistrict = fields[5]
		entries[i-1].Year = fields[6]
		entries[i-1].FillingDate = fields[7]
		entries[i-1].DocID = fields[8]
	}
	return entries, nil
}
